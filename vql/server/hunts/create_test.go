package hunts

import (
	"testing"
	"time"

	"github.com/Velocidex/ordereddict"
	"github.com/stretchr/testify/suite"
	"github.com/Cyarun/CyFir/file_store/test_utils"
	"github.com/Cyarun/CyFir/json"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/services"
	"github.com/Cyarun/CyFir/services/hunt_dispatcher"
	"github.com/Cyarun/CyFir/utils"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vql/acl_managers"
	"github.com/Cyarun/CyFir/vtesting/assert"
	"github.com/Cyarun/CyFir/vtesting/goldie"
)

var testArtifacts = []string{`
name: Test.Artifact
`, `
name: Server.Audit.Logs
type: SERVER_EVENT
`}

type TestSuite struct {
	test_utils.TestSuite
	client_id, flow_id string

	acl_manager vql_subsystem.ACLManager
}

func (self *TestSuite) SetupTest() {
	self.ConfigObj = self.LoadConfig()
	self.ConfigObj.Services.HuntDispatcher = true

	self.TestSuite.SetupTest()

	// Create an administrator user
	err := services.GrantRoles(self.ConfigObj, "admin", []string{"administrator"})
	assert.NoError(self.T(), err)

	self.acl_manager = acl_managers.NewServerACLManager(
		self.ConfigObj, "admin")
}

var testCases = []struct {
	description string
	args        *ordereddict.Dict
}{
	{
		description: "simple hunt",
		args: ordereddict.NewDict().
			Set("description", "foo").
			Set("include_labels", "Label1").
			Set("artifacts", "Test.Artifact"),
	},
	{
		description: "exclude label hunt",
		args: ordereddict.NewDict().
			Set("description", "foo").
			Set("exclude_labels", "ExcludeLabel").
			Set("artifacts", "Test.Artifact"),
	},
	{
		description: "include label hunt",
		args: ordereddict.NewDict().
			Set("description", "foo").
			Set("include_labels", "IncludeLabel").
			Set("artifacts", "Test.Artifact"),
	},
	{
		description: "include and exclude label hunt",
		args: ordereddict.NewDict().
			Set("description", "foo").
			Set("exclude_labels", "ExcludeLabel").
			Set("include_labels", "IncludeLabel").
			Set("artifacts", "Test.Artifact"),
	},
	{
		description: "os hunt",
		args: ordereddict.NewDict().
			Set("description", "foo").
			Set("os", "windows").
			Set("artifacts", "Test.Artifact"),
	},
	{
		description: "os and label hunt",
		args: ordereddict.NewDict().
			Set("description", "foo").
			Set("os", "windows").
			Set("include_labels", "IncludeLabel").
			Set("artifacts", "Test.Artifact"),
	},
}

func (self *TestSuite) TestCreateHunt() {
	result := ordereddict.NewDict()
	hunt_dispatcher.SetHuntIdForTests("H.1234")

	closer := utils.MockTime(utils.NewMockClock(time.Unix(100, 10)))
	defer closer()

	repository := self.LoadArtifacts(testArtifacts...)
	builder := services.ScopeBuilder{
		Config:     self.ConfigObj,
		ACLManager: self.acl_manager,
		Repository: repository,
		Logger: logging.NewPlainLogger(
			self.ConfigObj, &logging.FrontendComponent),
		Env: ordereddict.NewDict(),
	}

	manager, err := services.GetRepositoryManager(self.ConfigObj)
	assert.NoError(self.T(), err)
	scope := manager.BuildScope(builder)
	defer scope.Close()

	plugin := &ScheduleHuntFunction{}
	for _, test_case := range testCases {
		hunt_dispatcher.SetHuntIdForTests("H.1234")

		result.Set(test_case.description, plugin.Call(
			self.Ctx, scope, test_case.args))
	}
	goldie.Assert(self.T(), "TestCreateHunt",
		json.MustMarshalIndent(result))
}

func TestHuntPlugin(t *testing.T) {
	suite.Run(t, &TestSuite{
		client_id: "C.123",
		flow_id:   "F.123",
	})
}
