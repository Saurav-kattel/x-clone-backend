package encoder

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWriter(res http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(data)
}
