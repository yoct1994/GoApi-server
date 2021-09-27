package utils

import (
	"net/http"

	"github.com/GoAPI-server/dto/response"
)

func BadRequest(rw http.ResponseWriter) {
	res := response.ErrorRes{
		Status:  400,
		Message: "Bad request",
	}
	MarshalAndRW(400, res, rw)
}

func PostNotFound(rw http.ResponseWriter) {
	res := response.ErrorRes{
		Status:  404,
		Message: "Not found",
	}
	MarshalAndRW(404, res, rw)
}
