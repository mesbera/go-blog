package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mesbera/go-blog/api/auth"
	"github.com/mesbera/go-blog/api/models"
	"github.com/mesbera/go-blog/api/responses"
	"github.com/mesbera/go-blog/api/utils/formaterror"
)

//CreateComment: Creates a comment to the related post by the author of the post
func (server *Server) CreateComment(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	comment := models.Comment{}
	err = json.Unmarshal(body, &comment)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	comment.Prepare()
	err = comment.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//ToDo: What's the difference here between the two error responses below?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if uid != comment.AuthorId {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	isAuthor, err := comment.IsPostAuthor(server.DB)
	if err != nil || isAuthor == false {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("You are not authorized to write a comment to this post"))
		return
	}

	commentCreated, err := comment.SaveComment(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, commentCreated.CommentId))
	responses.JSON(w, http.StatusCreated, commentCreated)
}
