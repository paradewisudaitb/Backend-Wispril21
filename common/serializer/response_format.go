package serializer

import (
	"net/http"

	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
)

var RESPONSE_OK = ResponseBase{
	Code:    http.StatusOK,
	Message: statuscode.OK.String(),
}

var RESPONSE_NOT_FOUND = ResponseBase{
	Code:    http.StatusNotFound,
	Message: statuscode.NotFound.String(),
}
