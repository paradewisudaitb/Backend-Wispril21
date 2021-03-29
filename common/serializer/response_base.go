package serializer

import (
	sc "github.com/paradewisudaitb/Backend/common/constant/statuscode"
)

type ResponseBase struct {
	Code    sc.StatusCode `json:"status"`
	Message string        `json:"status_message"`
}

type ResponseData struct {
	ResponseBase
	Data interface{} `json:"data"`
}
