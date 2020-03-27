package modeltests

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/assert.v1"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
)

func TestCountPostsAndCommentsByUserId(t *testing.T) {
	err := refreshUserPostAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	user, _, _, err := seedUsersPostsAndComments()
	if err != nil {
		log.Fatal(err)
	}
	counted, err := userInstance.CountPostsAndCommentsByUserId(server.DB, user[0].UserId, 5)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, counted, int32(2))
}
