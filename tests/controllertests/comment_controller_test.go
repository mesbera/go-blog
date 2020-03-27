package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateComment(t *testing.T) {

	err := refreshUserPostAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	users, _, _, err := seedUsersPostsAndComments()
	if err != nil {
		log.Fatal(err)
	}

	token, err := server.SignIn(users[0].Email, "password")
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenString := fmt.Sprintf("Bearer %v", token)

	samples := []struct {
		inputJSON    string
		statusCode   int
		author_id    uint32
		post_id      uint64
		comment      string
		tokenGiven   string
		errorMessage string
	}{
		{
			inputJSON:    `{"comment": "random comment1", "post_id": 1, "author_id": 1 }`,
			statusCode:   201,
			author_id:    1,
			post_id:      1,
			comment:      "random comment1",
			tokenGiven:   tokenString,
			errorMessage: "",
		},
		{
			inputJSON:    `{"comment": "random comment2", "post_id": 2, "author_id": 1 }`,
			statusCode:   401,
			author_id:    1,
			post_id:      2,
			comment:      "random comment2",
			tokenGiven:   tokenString,
			errorMessage: "You are not authorized to write a comment to this post",
		},
		{
			inputJSON:    `{"comment": "random comment3", "post_id": 1, "author_id": 2 }`,
			statusCode:   401,
			author_id:    2,
			post_id:      1,
			comment:      "random comment3",
			tokenGiven:   tokenString,
			errorMessage: "Unauthorized",
		},
		//ToDo: check for missing post_id and author_id too
		{
			inputJSON:    `{"comment": "", "post_id": 1, "author_id": 1 }`,
			statusCode:   422,
			author_id:    1,
			post_id:      1,
			comment:      "",
			tokenGiven:   tokenString,
			errorMessage: "Required comment",
		},
	}
	for _, v := range samples {

		req, err := http.NewRequest("POST", "/comments", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateComment)

		req.Header.Set("Authorization", v.tokenGiven)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["comment"], v.comment)
			assert.Equal(t, responseMap["author_id"], float64(v.author_id)) //just for both ids to have the same type
			assert.Equal(t, responseMap["post_id"], float64(v.post_id))
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
