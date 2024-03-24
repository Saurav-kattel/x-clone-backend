package encoder

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter(res http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		res.WriteHeader(500)
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(data)
}
