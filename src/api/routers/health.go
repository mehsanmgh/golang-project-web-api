package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mehsanmgh/golang-project-web-api/api/handlers"
)

func Health(r *gin.RouterGroup){

	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
// 	r.POST("/", handler.HealthPost)
// 	r.POST("/:id", handler.HealthPostById)
}