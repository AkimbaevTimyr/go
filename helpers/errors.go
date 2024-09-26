package helpers

import (
	"net/http"
	"strings"
)

type errorCode string

const (
	ENOTFOUND      errorCode = "not_found"
	EINTERNAL      errorCode = "internal server error"
	UNAUTHORIZED   errorCode = "unauthorized"
	INVALIDPAYLOAD errorCode = "Invalid request payload"
)

var codeToHTTPStatusMap = map[errorCode]int{
	ENOTFOUND:      http.StatusNotFound,
	EINTERNAL:      http.StatusInternalServerError,
	UNAUTHORIZED:   http.StatusUnauthorized,
	INVALIDPAYLOAD: http.StatusBadRequest,
}

type Error struct {
	Err     error `json:"err"`
	Fields  map[string]interface{}
	Code    errorCode `json:"code"`
	Message string    `json:"message"`
	Op      string    `json:"op"`
}

func (e *Error) Error() string {
	var buf strings.Builder

	// if e.Op != "" {
	// 	buf.WriteString(e.Op)
	// 	buf.WriteString(": ")
	// }

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			buf.WriteRune('<')
			buf.WriteString(string(e.Code))
			buf.WriteRune('>')
		}
		if e.Code != "" && e.Message != "" {
			buf.WriteRune(' ')
		}
		buf.WriteString(e.Message)
	}

	return buf.String()
}

func (e *Error) HTTPStatus() int {
	if status, ok := codeToHTTPStatusMap[e.Code]; ok {
		return status
	}
	return http.StatusInternalServerError
}

func (e *Error) Details() map[string]any {
	return map[string]any{
		"message": e.Message,
	}
}
