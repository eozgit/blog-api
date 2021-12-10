package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	_, w := setupTest(t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":200}`, removeSpaces(w.Body.String()))
}

func TestCreatePost(t *testing.T) {
	app, w := setupTest(t)

	payload := `{"Title":"Certainty Of Death? Small Chance Of Success?","Content":"What Are We Waitin'' For?","Categories":["Literature", "Cinema"]}`
	req, _ := http.NewRequest("POST", "/post", strings.NewReader(payload))
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("gimli:noonetossesadwarf")))
	app.r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":200}`, removeSpaces(w.Body.String()))
}

func removeSpaces(str string) string {
	return strings.Join(strings.Fields(str), "")
}

func setupTest(t *testing.T) (*App, *httptest.ResponseRecorder) {
	app = setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"Username":"gimli","Password":"nobodytossesadwarf"}`))
	app.r.ServeHTTP(w, req)

	return app, w
}
