package main

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	var router *gin.Engine = gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", getAlbums)

	return router
}
