package utils

import (
	"encoding/json"
	"net/http"

	"github.com/savsgio/go-logger/v2"
)

func MarshalAndRW(status int, res interface{}, rw http.ResponseWriter) {
	rw.WriteHeader(status)

	resBytes, err := json.MarshalIndent(res, "", "	")
	HandleErr(err)

	rw.Write(resBytes)
}

func HandleErr(err error) {
	if err != nil {
		logger.Error(err)
	}
}
