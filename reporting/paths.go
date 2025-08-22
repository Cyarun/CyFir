package reporting

import (
	"github.com/Cyarun/CyFir/accessors"
)

type NotebookExportPathManager struct {
	notebook_id string
	root        *accessors.OSPath
}

func (self *NotebookExportPathManager) AttachmentRoot() *accessors.OSPath {
	return self.root.Append(self.notebook_id, "attach")
}

func (self *NotebookExportPathManager) CellDirectory(
	cell_id string) *accessors.OSPath {
	return self.root.Append(self.notebook_id, cell_id)
}

func NewNotebookExportPathManager(notebook_id string) *NotebookExportPathManager {
	root, _ := accessors.NewZipFilePath("/")

	return &NotebookExportPathManager{
		notebook_id: notebook_id,
		root:        root,
	}
}
