package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
