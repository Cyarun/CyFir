package server_artifacts

import (
	"context"
	"time"

	"github.com/Cyarun/CyFir/result_sets"
	"github.com/Cyarun/CyFir/utils"
)

func ResultSetFlusher(ctx context.Context, rs_writer result_sets.ResultSetWriter) func() {
	sub_ctx, cancel := context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-sub_ctx.Done():
				return

			case <-time.After(utils.Jitter(time.Duration(10) * time.Second)):
				rs_writer.Flush()
			}
		}
	}()

	return cancel
}
