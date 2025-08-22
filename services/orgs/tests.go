package orgs

import (
	"context"
	"sync"

	api_proto "github.com/Cyarun/CyFir/api/proto"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/services"
)

type TestOrgManager struct {
	*OrgManager
	services *ServiceContainer
}

func (self *TestOrgManager) SetOrgIdForTesting(a string) {
	self.OrgManager.SetOrgIdForTesting(a)
}

func (self *TestOrgManager) Start(
	ctx context.Context,
	org_config *config_proto.Config,
	wg *sync.WaitGroup) error {
	logger := logging.GetLogger(org_config, &logging.FrontendComponent)
	logger.Info("<green>Starting</> Test Org Manager service.")

	self.mu.Lock()
	service_container := &ServiceContainer{}
	if self.services != nil {
		service_container = self.services
	}

	org_context := &OrgContext{
		record:     &api_proto.OrgRecord{},
		config_obj: org_config,
		service:    service_container,
		sm:         services.NewServiceManager(ctx, org_config),
	}

	self.orgs[services.ROOT_ORG_ID] = org_context
	self.mu.Unlock()

	return self.startOrgFromContext(org_context)
}

func StartTestOrgManager(
	ctx context.Context,
	wg *sync.WaitGroup,
	config_obj *config_proto.Config,
	service_container *ServiceContainer) error {

	service := &TestOrgManager{
		services: service_container,
		OrgManager: &OrgManager{
			config_obj: config_obj,
			ctx:        ctx,
			parent_wg:  wg,

			orgs:            make(map[string]*OrgContext),
			org_id_by_nonce: make(map[string]string),
		}}
	services.RegisterOrgManager(service)

	return service.Start(ctx, config_obj, wg)
}
