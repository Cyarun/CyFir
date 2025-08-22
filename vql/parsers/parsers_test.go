package parsers_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Velocidex/ordereddict"
	"github.com/stretchr/testify/suite"
	"github.com/Cyarun/CyFir/json"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vql/parsers"
	"github.com/Cyarun/CyFir/vtesting/goldie"

	_ "github.com/Cyarun/CyFir/accessors/data"
)

type ParserTestSuite struct {
	suite.Suite
}

var yamlTestCases = []string{`
Name: myname
Value:
  - list1
  - list2
`, `
# Test that we maintain the order of keys in yaml documents.
Name: myname1
B: 1
A: 2
c:
  Field1: X
  Field2: Y
  Field0: Z
`, `
# Document with errors - should return null
Field1
  field2
`}

func (self *ParserTestSuite) TestYamlParser() {
	result := ordereddict.NewDict()
	ctx := context.Background()
	scope := vql_subsystem.MakeScope()
	scope.SetLogger(log.New(os.Stderr, "", 0))

	defer scope.Close()

	for idx, item := range yamlTestCases {
		value := parsers.ParseYamlFunction{}.Call(ctx, scope, ordereddict.NewDict().
			Set("filename", item).
			Set("accessor", "data"))
		result.Set(fmt.Sprintf("Case %v", idx), value)
	}
	goldie.Assert(self.T(), "TestYamlParser", json.MustMarshalIndent(result))
}

func TestParsers(t *testing.T) {
	suite.Run(t, &ParserTestSuite{})
}
