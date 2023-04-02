package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAlbums(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums", getAlbums)

	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got []Album
	json.NewDecoder(w.Body).Decode(&got)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, albums, got)
}

func TestPostAlbums(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albums", getAlbums)

	var newAlbum Album = Album{
		ID:     "1",
		Title:  "New Album",
		Artist: "Gustavo",
		Price:  100,
	}

	jsonValue, _ := json.Marshal(newAlbum)

	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got []Album
	json.NewDecoder(w.Body).Decode(&got)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, albums, got)
}
