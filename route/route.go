package route

import (
	"application/controller"
	"application/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoRoute(app *gin.Engine) {
	app.GET("/todos", func(ctx *gin.Context) {
		result, err := controller.TodoFindAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": "server internal error",
			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]any{
			"message": "get all todos",
			"data":    result,
		})
	})
	app.GET("/todos/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		result, err := controller.TodoFindById(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]any{
			"message": "Get By ID",
			"data":    result,
		})

	})
	app.POST("/todos", func(ctx *gin.Context) {
		type request struct {
			Title string `json:"title" binding:"required"`
		}

		var body request
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed",
			})
			return
		}

		result, err := controller.TodoCreate(ctx, body.Title)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to create",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "todo created",
			"data":    result,
		})
	})
	app.PUT("/todos/:id", func(ctx *gin.Context) {

		// id := ctx.Params.byName("id")
		var body model.Todo
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed Update",
			})
			return
		}

		result, err := controller.TodoUpdate(ctx, body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to Update",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "todo Update",
			"data":    result,
		})
	})
	app.DELETE("/todos/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		err := controller.TodoDelete(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]string{
			"message": "deleted",
		})
	})

}
