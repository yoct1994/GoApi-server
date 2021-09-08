package move

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GoAPI-server/dto"
	"github.com/GoAPI-server/utils"
	"github.com/savsgio/go-logger/v2"
)

var posts map[int64]interface{} = make(map[int64]interface{})
var key int64 = 1

func WritePost(rw http.ResponseWriter, r *http.Request) {
	var data dto.WritePost
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.BadRequest(rw)
	}
	post := dto.Post{Title: data.Title, Content: data.Content, Name: data.Name, WriteAt: time.Now()}

	posts[key] = post

	key++

	logger.Info(post.Title)
}
