package notebook

import (
	"context"
	"sort"
	"time"

	"github.com/Velocidex/ordereddict"
	api_proto "github.com/Cyarun/CyFir/api/proto"
	"github.com/Cyarun/CyFir/file_store"
	"github.com/Cyarun/CyFir/file_store/api"
	"github.com/Cyarun/CyFir/json"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/paths"
	"github.com/Cyarun/CyFir/result_sets"
	"github.com/Cyarun/CyFir/services"
	"github.com/Cyarun/CyFir/utils"
)

func (self *NotebookManager) CheckNotebookAccess(
	notebook *api_proto.NotebookMetadata, user string) bool {
	return checkNotebookAccess(notebook, user)
}

func checkNotebookAccess(notebook *api_proto.NotebookMetadata, user string) bool {
	if notebook.Public {
		return true
	}

	return notebook.Creator == user || utils.InString(notebook.Collaborators, user)
}

// Returns all the notebooks which are either owned or shared with the
// user. This view is only called from the global notebook view so it
// only needs to return a brief version of the notebooks - it does not
// include uploads and timelines.
func (self *NotebookManager) GetSharedNotebooks(
	ctx context.Context, username string) (api.FSPathSpec, error) {

	notebook_path_manager := paths.NewNotebookPathManager("")
	file_store_factory := file_store.GetFileStore(self.config_obj)
	index_filename := notebook_path_manager.NotebookIndexForUser(username)

	stat, err := file_store_factory.StatFile(index_filename)

	if err == nil && stat.ModTime().Unix() >= self.Store.Version() {
		return index_filename, nil
	}

	logger := logging.GetLogger(self.config_obj, &logging.GUIComponent)
	logger.Debug("Building notebook index for %v\n", username)

	rs_writer, err := result_sets.NewResultSetWriter(
		file_store_factory, index_filename,
		json.DefaultEncOpts(), utils.SyncCompleter,
		result_sets.TruncateMode)
	if err != nil {
		return nil, err
	}
	defer rs_writer.Close()

	all_notebooks, err := self.GetAllNotebooks(ctx,
		services.NotebookSearchOptions{
			Username: username,
		})
	if err != nil {
		return nil, err
	}

	sort.Slice(all_notebooks, func(i, j int) bool {
		return all_notebooks[i].NotebookId > all_notebooks[j].NotebookId
	})

	for _, notebook := range all_notebooks {
		rs_writer.Write(ordereddict.NewDict().
			Set("NotebookId", notebook.NotebookId).
			Set("Name", notebook.Name).
			Set("Description", notebook.Description).
			Set("Creation Time", time.Unix(notebook.CreatedTime, 0)).
			Set("Modified Time", time.Unix(notebook.ModifiedTime, 0)).
			Set("Creator", notebook.Creator).
			Set("Collaborators", notebook.Collaborators))
	}

	return index_filename, nil
}

func (self *NotebookManager) GetAllNotebooks(ctx context.Context,
	opts services.NotebookSearchOptions) ([]*api_proto.NotebookMetadata, error) {
	return self.Store.GetAllNotebooks(ctx, opts)
}
