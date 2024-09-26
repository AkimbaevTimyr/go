package helpers

import (
	"encoding/json"
	"net/http"
)

type Envelope map[string]interface{}

func ReadJson(r *http.Request, w http.ResponseWriter, v interface{}) *Error {
	//set request body to 1 mb to prevent any potential dos attacks
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(v)

	if err != nil {
		return &Error{Code: INVALIDPAYLOAD, Message: "invalid payload"}
	}

	return nil

}
