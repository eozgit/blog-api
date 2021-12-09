package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	app = setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader("{\"Username\":\"gimli\",\"Password\":\"nobodytossesadwarf\"}"))
	app.r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"status\":200}", w.Body.String())
}
