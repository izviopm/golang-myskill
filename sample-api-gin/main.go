package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type People struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	//ping
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// coba grouping
	v1 := r.Group("v1")
	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": param,
		})
	})

	v1.POST("/user", func(ctx *gin.Context) {
		var data People
		ctx.BindJSON(&data)

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": data,
		})
	})
	
	v1.GET("/user", func(ctx *gin.Context) {
		query := ctx.Query("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": query,
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}