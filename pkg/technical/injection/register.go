package injection

import (
	"ddd/pkg/drivingAdapter/api/operation"
	api "ddd/pkg/drivingAdapter/api/shared"
)

func RegisterHTTPHandler(r *api.Router, dHandler *operation.DebugHandler) {
	r.PUT("debug/logLevel", dHandler.UpdateLogLevel)
}
