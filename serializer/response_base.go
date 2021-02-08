package serializer

import (
	"github.com/paradewisudaitb/Backend/common/constant/stcode"
)

type ResponseBase struct {
	Code    stcode.StatusCode `json:"status"`
	Message string            `json:"status_message"`
}
