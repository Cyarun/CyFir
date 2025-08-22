package acl_managers_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/Cyarun/CyFir/acls"
	api_proto "github.com/Cyarun/CyFir/api/proto"
	"github.com/Cyarun/CyFir/file_store/test_utils"
	"github.com/Cyarun/CyFir/services"
	"github.com/Cyarun/CyFir/services/sanity"
	"github.com/Cyarun/CyFir/vql/acl_managers"
	"github.com/Cyarun/CyFir/vtesting/assert"
)

var (
	mock_definitions = []string{`
name: Server.Internal.UserManager
type: INTERNAL
`}
)

type TestSuite struct {
	test_utils.TestSuite
}

func (self *TestSuite) SetupTest() {
	self.ConfigObj = self.TestSuite.LoadConfig()
	self.LoadArtifactsIntoConfig(mock_definitions)
	self.TestSuite.SetupTest()
}

func (self *TestSuite) TestLockdown() {
	users_manager := services.GetUserManager()

	// Create an admin user with administrator role.
	err := users_manager.SetUser(self.Ctx, &api_proto.VelociraptorUser{
		Name: "admin",
	})
	assert.NoError(self.T(), err)

	err = services.GrantRoles(self.ConfigObj, "admin", []string{"administrator"})
	assert.NoError(self.T(), err)

	acl_manager := acl_managers.NewServerACLManager(self.ConfigObj, "admin")

	// Check the user has COLLECT_CLIENT
	ok, err := acl_manager.CheckAccess(acls.COLLECT_CLIENT)
	assert.NoError(self.T(), err)
	assert.True(self.T(), ok)

	// Now simulate lockdown - first set the config file.
	self.ConfigObj.Lockdown = true

	// Now start the sanity checker because it will configure lock down.
	err = sanity.NewSanityCheckService(self.Ctx, nil, self.ConfigObj)
	assert.NoError(self.T(), err)

	// Checking again should reject it due to lockdown.
	ok, err = acl_manager.CheckAccess(acls.COLLECT_CLIENT)
	assert.ErrorContains(self.T(), err, "Server locked down")
	assert.False(self.T(), ok)
}

func TestServerACLManager(t *testing.T) {
	suite.Run(t, &TestSuite{})
}
