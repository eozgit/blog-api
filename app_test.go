package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	_, w := setupTest(t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{ \"status\": \"User created\" }", unindent(w.Body.String()))
}

func TestCreatePost(t *testing.T) {
	app, _ := setupTest(t)

	payload := `{"Title":"Certainty Of Death? Small Chance Of Success?","Content":"What Are We Waitin'' For?","Categories":["Literature", "Cinema"]}`
	req, _ := http.NewRequest("POST", "/post", strings.NewReader(payload))
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("gimli:noonetossesadwarf")))
	rec := httptest.NewRecorder()
	app.r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{ \"status\": \"Post created\" }", unindent(rec.Body.String()))
}

func TestCreateComment(t *testing.T) {
	app, _ := setupTest(t)

	payload := `{"Title":"I would have gone with you to the end","Content":"Into the very fires of Mordor"}`
	req, _ := http.NewRequest("POST", "/post/1/comment", strings.NewReader(payload))
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("gimli:noonetossesadwarf")))
	rec := httptest.NewRecorder()
	app.r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{ \"status\": \"Comment created\" }", unindent(rec.Body.String()))
}

func unindent(str string) string {
	var re = regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(str, " ")
}

func setupTest(t *testing.T) (*App, *httptest.ResponseRecorder) {
	app = setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"Username":"gimli","Password":"noonetossesadwarf"}`))
	app.r.ServeHTTP(w, req)

	return app, w
}
