package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoAPI-server/service/middleware"
	"github.com/gorilla/mux"
)

var port string

func Start(aPort int) {
	port = fmt.Sprintf("%d", aPort)

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.SetJsonContentType)

	log.Fatal(http.ListenAndServe(port, router))
}
