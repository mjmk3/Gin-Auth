package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	route := app
	isValidated := true

	route.GET("/", func(ctx *gin.Context) {

		if !isValidated {
			ctx.AbortWithStatusJSON(400, gin.H{
				"message": "Bad Request",
			})

			return
		}

		ctx.JSON(200, gin.H{
			"hello": "Golang",
		})
	})

	app.Run(":8999")
}
