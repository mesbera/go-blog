package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mesbera/go-blog/api/models"
	"github.com/mesbera/go-blog/api/responses"
)

//GetUserEntryCount counts counts all the posts and comments added by the user in the past X days
func (server *Server) GetUserEntryCount(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var input map[string]interface{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	counted, err := user.CountPostsAndCommentsByUserId(server.DB, uint32(input["user_id"].(float64)), int(input["days"].(float64)))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resultmap := make(map[string]int32)
	resultmap["result"] = counted
	responses.JSON(w, http.StatusOK, resultmap)
}
