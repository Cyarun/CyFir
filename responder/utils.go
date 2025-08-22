package responder

import (
	constants "github.com/Cyarun/CyFir/constants"
	crypto_proto "github.com/Cyarun/CyFir/crypto/proto"
	"github.com/Cyarun/CyFir/json"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/utils"
)

func MakeErrorResponse(
	output chan *crypto_proto.VeloMessage, session_id string, message string) {
	output <- &crypto_proto.VeloMessage{
		SessionId: session_id,
		RequestId: constants.LOG_SINK,
		LogMessage: &crypto_proto.LogMessage{
			NumberOfRows: 1,
			Jsonl: json.Format(
				"{\"client_time\":%d,\"level\":%q,\"message\":%q}\n",
				int(utils.GetTime().Now().Unix()), logging.ERROR, message),
			ErrorMessage: message,
		},
	}
}
