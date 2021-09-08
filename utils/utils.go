package utils

import (
	"encoding/json"
	"net/http"
)

func MarshalAndRW(status int, res interface{}, rw http.ResponseWriter) {
	rw.WriteHeader(status)

	resBytes, _ := json.MarshalIndent(res, "", " ")

	rw.Write(resBytes)
}
