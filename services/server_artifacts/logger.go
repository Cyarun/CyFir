package server_artifacts

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/Velocidex/ordereddict"
	"github.com/Cyarun/CyFir/artifacts"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/constants"
	crypto_proto "github.com/Cyarun/CyFir/crypto/proto"
	"github.com/Cyarun/CyFir/file_store"
	"github.com/Cyarun/CyFir/json"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/paths"
	"github.com/Cyarun/CyFir/result_sets"
	"github.com/Cyarun/CyFir/utils"
)

var (
	defaultLogErrorRegex = regexp.MustCompile(constants.VQL_ERROR_REGEX)
)

// A reference counter around ResultSetWriter to ensure it is only
// closed when no more references are found.
type counterWriter struct {
	result_sets.ResultSetWriter

	mu    sync.Mutex
	count int
}

func (self *counterWriter) Copy() *counterWriter {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.count++
	return self
}

func (self *counterWriter) Close() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.count--

	if self.count == 0 {
		self.ResultSetWriter.Close()
	}

	if self.count < 0 {
		panic("Negative counterWriter!")
	}
}

type serverLogger struct {
	config_obj    *config_proto.Config
	writer        result_sets.ResultSetWriter
	query_context QueryContext
}

func (self *serverLogger) Close() {
	self.writer.Close()
}

func (self *serverLogger) Write(b []byte) (int, error) {
	level, msg := logging.SplitIntoLevelAndLog(b)
	msg = artifacts.DeobfuscateString(self.config_obj, msg)

	self.writer.Write(ordereddict.NewDict().
		Set("Timestamp", utils.GetTime().Now().UTC().UnixNano()/1000).
		Set("Level", level).
		Set("message", msg))

	// Increment the log count.
	if self.query_context != nil {
		self.query_context.UpdateStatus(func(s *crypto_proto.VeloStatus) {
			s.LogRows++
		})

		// If an error occured mark the collection failed.
		if level == "ERROR" || defaultLogErrorRegex.MatchString(msg) {
			self.query_context.UpdateStatus(func(s *crypto_proto.VeloStatus) {
				s.Status = crypto_proto.VeloStatus_GENERIC_ERROR
				s.ErrorMessage = msg
			})
		}
	}

	return len(b), nil
}

func NewServerLogWriter(
	ctx context.Context,
	config_obj *config_proto.Config,
	session_id string) (result_sets.ResultSetWriter, error) {

	path_manager := paths.NewFlowPathManager("server", session_id)
	file_store_factory := file_store.GetFileStore(config_obj)
	log_writer, err := result_sets.NewResultSetWriter(file_store_factory,
		path_manager.Log(), json.DefaultEncOpts(),
		utils.BackgroundWriter, result_sets.AppendMode)
	if err != nil {
		return nil, err
	}

	// Flush the logs every second to make sure the GUI shows
	// progress.
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-time.After(utils.Jitter(time.Second)):
				log_writer.Flush()
			}
		}
	}()

	return log_writer, nil
}
