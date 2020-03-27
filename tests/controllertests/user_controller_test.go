package controllertests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetUserEntryCount(t *testing.T) {
	err := refreshUserPostAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}

	_, _, _, err = seedUsersPostsAndComments()
	if err != nil {
		log.Fatal(err)
	}
	//ToDo: for cycle to check multiple possiblilities based on db seed.
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"user_id": "1", "days": "5"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUserEntryCount)
	handler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, 2, responseMap["result"])
}
