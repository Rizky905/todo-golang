package main

import (
	"application/route"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	//router
	route.TodoRoute(app)
	app.Run("0.0.0.0:8000")
}
