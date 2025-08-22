package paths

import "github.com/Cyarun/CyFir/file_store/api"

type OrgPathManager struct {
	org_id string
}

func (self OrgPathManager) Path() api.DSPathSpec {
	return ORGS_ROOT.AddChild(self.org_id)
}

func NewOrgPathManager(org_id string) *OrgPathManager {
	return &OrgPathManager{org_id}
}
