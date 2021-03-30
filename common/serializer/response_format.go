package serializer

import (
	"net/http"

	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
)

var RESPONSE_OK = ResponseBase{
	Code:    http.StatusOK,
	Message: statuscode.OK.String(),
}

var RESPONSE_FORBIDDEN = ResponseBase{
	Code:    http.StatusForbidden,
	Message: statuscode.NoAccess.String(),
}
