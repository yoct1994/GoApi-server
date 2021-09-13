package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoAPI-server/db"
	"github.com/GoAPI-server/service/middleware"
	"github.com/GoAPI-server/service/post"
	"github.com/gorilla/mux"
	"github.com/savsgio/go-logger/v2"
)

var port string

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)

	err := db.InitDB()
	if err != nil {
		db.CloseDB()
		logger.Error("Failed to initialize")
		return
	}

	db.AutoCreateDB()

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.SetJsonContentType)

	router.HandleFunc("/post", post.WritePost).Methods("POST")
	router.HandleFunc("/post/{pageNum}", post.ReadPost).Methods("GET")

	logger.Infof("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
