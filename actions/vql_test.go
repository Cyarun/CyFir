package actions_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/Cyarun/CyFir/actions"
	actions_proto "github.com/Cyarun/CyFir/actions/proto"
	artifacts_proto "github.com/Cyarun/CyFir/artifacts/proto"
	crypto_proto "github.com/Cyarun/CyFir/crypto/proto"
	"github.com/Cyarun/CyFir/file_store/test_utils"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/responder"
	"github.com/Cyarun/CyFir/vtesting"
	"github.com/Cyarun/CyFir/vtesting/assert"

	// For execve and query
	_ "github.com/Cyarun/CyFir/vql/common"
	_ "github.com/Cyarun/CyFir/vql/tools"
)

type ClientVQLTestSuite struct {
	test_utils.TestSuite
}

func (self *ClientVQLTestSuite) SetupTest() {
	self.ConfigObj = self.LoadConfig()
	self.ConfigObj.Client.PreventExecve = true
	self.TestSuite.SetupTest()
}

func (self *ClientVQLTestSuite) TestCPUThrottler() {
	request := &actions_proto.VQLCollectorArgs{
		Query: []*actions_proto.VQLRequest{
			{
				Name: "Query",
				VQL:  "SELECT 'Boo' FROM scope()",
			},
		},
	}

	// Query is not limited
	resp := responder.TestResponderWithFlowId(
		self.ConfigObj, "TestCPUThrottler")
	actions.VQLClientAction{}.StartQuery(
		self.ConfigObj, self.Sm.Ctx, resp, request)
	resp.Close()

	assert.NotContains(self.T(), getLogs(resp.Drain.Messages()),
		"Will throttle query")

	// Query will now be limited
	resp = responder.TestResponderWithFlowId(
		self.ConfigObj, "TestCPUThrottler2")
	defer resp.Close()

	request.CpuLimit = 20
	actions.VQLClientAction{}.StartQuery(
		self.ConfigObj, self.Sm.Ctx, resp, request)

	var responses []*crypto_proto.VeloMessage
	vtesting.WaitUntil(5*time.Second, self.T(), func() bool {
		responses = resp.Drain.Messages()
		return strings.Contains(getLogs(responses), "Will throttle query")
	})

	assert.Contains(self.T(), getLogs(responses), "Will throttle query")
}

// Make sure that dependent artifacts are properly used
func (self *ClientVQLTestSuite) TestDependentArtifacts() {
	resp := responder.TestResponderWithFlowId(
		self.ConfigObj, "TestDependentArtifacts")

	actions.VQLClientAction{}.StartQuery(self.ConfigObj, self.Sm.Ctx, resp,
		&actions_proto.VQLCollectorArgs{
			Query: []*actions_proto.VQLRequest{
				{
					Name: "Query",
					VQL:  "SELECT * FROM Artifact.Custom.Foo.Bar.Baz.A()",
				},
			},
			Artifacts: []*artifacts_proto.Artifact{
				{
					Name: "Custom.Foo.Bar.Baz.A",
					Sources: []*artifacts_proto.ArtifactSource{
						{
							Query: "SELECT * FROM Artifact.Custom.Foo.Bar.Baz.B()",
						},
					},
				},
				{
					Name: "Custom.Foo.Bar.Baz.B",
					Sources: []*artifacts_proto.ArtifactSource{
						{
							Query: "SELECT * FROM Artifact.Custom.Foo.Bar.Baz.C()",
						},
					},
				},
				{
					Name: "Custom.Foo.Bar.Baz.C",
					Sources: []*artifacts_proto.ArtifactSource{
						{
							Query: "SELECT 1 AS X FROM scope()",
						},
					},
				},
			},
		})

	var responses []*crypto_proto.VeloMessage
	vtesting.WaitUntil(5*time.Second, self.T(), func() bool {
		responses = resp.Drain.Messages()
		return "Target: Query, JSONL: {\"X\":1,\"_Source\":\"Custom.Foo.Bar.Baz.A\"}\n\n" ==
			getVQLResponse(responses)
	})
}

func (self *ClientVQLTestSuite) TestMaxRows() {
	resp := responder.TestResponderWithFlowId(self.ConfigObj, "TestMaxRows")

	actions.VQLClientAction{}.StartQuery(self.ConfigObj, self.Sm.Ctx, resp,
		&actions_proto.VQLCollectorArgs{
			MaxRow: 10,
			Query: []*actions_proto.VQLRequest{
				{
					Name: "Query",
					VQL:  "SELECT * FROM range(end=20)",
				},
			},
		})

	var responses []*crypto_proto.VeloMessage
	vtesting.WaitUntil(time.Second, self.T(), func() bool {
		responses = resp.Drain.Messages()
		payloads := getResponsePacketCounts(responses)
		return len(payloads) == 2 && payloads[0] == 10 && payloads[1] == 10
	})
}

func (self *ClientVQLTestSuite) TestExecve() {
	resp := responder.TestResponderWithFlowId(self.ConfigObj, "TestMaxRows")

	logging.ClearMemoryLogs()

	actions.VQLClientAction{}.StartQuery(self.ConfigObj, self.Sm.Ctx, resp,
		&actions_proto.VQLCollectorArgs{
			MaxRow: 10,
			Query: []*actions_proto.VQLRequest{
				{
					Name: "Query",
					VQL:  "SELECT * FROM execve(argv='ls')",
					//					VQL:  "SELECT * FROM query(query={ SELECT * FROM execve(argv='ls') })",
				},
			},
		})

	vtesting.WaitUntil(time.Second, self.T(), func() bool {
		return vtesting.MemoryLogsContainRegex(
			"execve: Not allowed to execve by configuration.")
	})

	logging.ClearMemoryLogs()

	// Make sure the query() plugin propagates the execve flag
	actions.VQLClientAction{}.StartQuery(self.ConfigObj, self.Sm.Ctx, resp,
		&actions_proto.VQLCollectorArgs{
			MaxRow: 10,
			Query: []*actions_proto.VQLRequest{
				{
					Name: "Query",
					VQL:  "SELECT * FROM query(query={ SELECT * FROM execve(argv='ls') })",
				},
			},
		})

	vtesting.WaitUntil(time.Second, self.T(), func() bool {
		return vtesting.MemoryLogsContainRegex(
			"execve: Not allowed to execve by configuration.")
	})

}

func (self *ClientVQLTestSuite) TestMaxWait() {
	assert.True(self.T(), test_utils.Retry(self.T(), 5, time.Millisecond*500,
		func(r *test_utils.R) {
			resp := responder.TestResponderWithFlowId(self.ConfigObj, "TestMaxWait")

			actions.VQLClientAction{}.StartQuery(self.ConfigObj, self.Sm.Ctx, resp,
				&actions_proto.VQLCollectorArgs{
					MaxRow:  1000,
					MaxWait: 1,
					Query: []*actions_proto.VQLRequest{
						{
							Name: "Query",
							VQL:  "SELECT sleep(ms=400) FROM range(end=4)",
						},
					},
				})

			var responses []*crypto_proto.VeloMessage

			vtesting.WaitUntil(5*time.Second, r, func() bool {
				responses = resp.Drain.Messages()
				payloads := getResponsePacketCounts(responses)
				// Message will be split into 2 packets 2 messages in each
				return len(payloads) == 2 && payloads[0] == 2 && payloads[1] == 2
			})
		}))
}

func TestClientVQL(t *testing.T) {
	suite.Run(t, &ClientVQLTestSuite{})
}

func getResponsePacketCounts(responses []*crypto_proto.VeloMessage) []uint64 {
	result := []uint64{}
	for _, item := range responses {
		if item.VQLResponse != nil {
			result = append(result, item.VQLResponse.TotalRows)
		}
	}

	return result
}
