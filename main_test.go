package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"fmt"
	"github.com/johnidm/shun"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	router := main.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}


func TestDetailRoute(t *testing.T) {
	router := main.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/detail/834534-dsfsdlk-4343", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAPIEmailTrack(t *testing.T) {
	router := main.SetupRouter()

	body := bytes.NewBuffer([]byte("{\"email\":\"sheldon@cupper.com\"}"))

	w := httptest.NewRecorder()
	req, _  := http.NewRequest("POST", "/api/v1/track/email/S8SND98DS-SDFSKFSD-SDAS00", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

func TestAPIUrlTrack(t *testing.T) {
	router := main.SetupRouter()

	body := bytes.NewBuffer([]byte("{\"title\":\"Home\", \"url\":\"localhost:5000\", \"date\":\"2017-11-19T22:06:13.198Z\"}"))

	w := httptest.NewRecorder()
	req, _  := http.NewRequest("POST", "/api/v1/track/url/S8SND98DS-SDFSKFSD-SDAS00", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)

}



