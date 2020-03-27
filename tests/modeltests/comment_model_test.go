package modeltests

import (
	"testing"

	"github.com/mesbera/go-blog/api/models"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/assert.v1"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
)

func TestIsPostAuthor(t *testing.T) {

}

func TestSaveComment(t *testing.T) {
	err := refreshUserPostAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	_, _, _, err = seedUsersPostsAndComments()
	if err != nil {
		log.Fatal(err)
	}
	newComment := models.Comment{
		CommentId: 3,
		AuthorId:  1,
		PostId:    1,
		Comment:   "random comment test",
	}
	savedComment, err := newComment.SaveComment(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newComment.AuthorId, savedComment.AuthorId)
	assert.Equal(t, newComment.PostId, savedComment.PostId)
	assert.Equal(t, newComment.Comment, savedComment.Comment)
}
