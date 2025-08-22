package executor

import (
	"context"

	"github.com/Cyarun/CyFir/actions"
	actions_proto "github.com/Cyarun/CyFir/actions/proto"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/constants"
	crypto_proto "github.com/Cyarun/CyFir/crypto/proto"
	"github.com/Cyarun/CyFir/responder"
	"github.com/Cyarun/CyFir/utils"
)

func (self *ClientExecutor) ResumeTransactions(
	ctx context.Context,
	config_obj *config_proto.Config, req *crypto_proto.VeloMessage) {

	if req.ResumeTransactions == nil {
		return
	}

	// Uncancel the flow.
	self.flow_manager.UnCancel(req.SessionId)

	flow_context := self.flow_manager.FlowContext(self.Outbound, req)
	defer flow_context.Close()

	// Responses for transactions go into a special result set called
	// "Resumed Uploads". If the flow was resumed previously the query
	// stats already have such an artifact, otherwise we create a new
	// one.
	var our_responder responder.Responder
	var our_stat *crypto_proto.VeloStatus

	for _, stat := range req.ResumeTransactions.QueryStats {
		_, responder_obj := flow_context.NewResponder(
			&actions_proto.VQLCollectorArgs{})
		defer responder_obj.Close()

		responder_obj.SetStatus(stat)

		if utils.InString(stat.NamesWithResponse,
			constants.UPLOAD_RESUMED_SOURCE) {
			our_responder = responder_obj
			our_stat = stat
		}
	}

	if our_responder == nil {
		_, new_responder := flow_context.NewResponder(
			&actions_proto.VQLCollectorArgs{})
		new_responder.SetStatus(&crypto_proto.VeloStatus{
			Status:            crypto_proto.VeloStatus_PROGRESS,
			NamesWithResponse: []string{constants.UPLOAD_RESUMED_SOURCE},
		})

		our_responder = new_responder
		our_stat = &crypto_proto.VeloStatus{}
	}

	actions.ResumeTransactions(
		ctx, config_obj, our_responder, our_stat, req.ResumeTransactions)
}
