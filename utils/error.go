package utils

import (
	"net/http"

	"github.com/GoAPI-server/dto"
)

func BadRequest(rw http.ResponseWriter) {
	res := dto.ErrorRes{
		Status:  400,
		Message: "Bad request",
	}
	MarshalAndRW(400, res, rw)
}
