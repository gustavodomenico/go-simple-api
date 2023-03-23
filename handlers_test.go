package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	// Given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	MockJsonGet(ctx, url.Values{})

	// When
	getAlbums(ctx)

	// Then
	assert.EqualValues(t, http.StatusOK, w.Code)

	var got []Album
	json.NewDecoder(w.Body).Decode(&got)

	assert.Equal(t, albums, got)
}
