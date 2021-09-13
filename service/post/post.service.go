package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GoAPI-server/dto/request"
	"github.com/GoAPI-server/model/repository"
	"github.com/GoAPI-server/utils"
	"github.com/gorilla/mux"
	"github.com/savsgio/go-logger/v2"
)

func WritePost(rw http.ResponseWriter, r *http.Request) {
	var data request.WritePost
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.BadRequest(rw)
		logger.Error("bad request", r.Body)
		return
	}

	err = repository.CreatePost(data)

	if err != nil {
		logger.Error("error create failed")
		return
	}

	rw.WriteHeader(200)
}

func ReadPost(rw http.ResponseWriter, r *http.Request) {
	pageNum := mux.Vars(r)["pageNum"]
	num, _ := strconv.Atoi(pageNum)
	posts, err := repository.GetAllPosts(num)

	if err != nil {
		logger.Error("get data error")
		return
	}

	err = json.NewEncoder(rw).Encode(&posts)

	if err != nil {
		logger.Error("failed encode")
	}
}
