package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mehsanmgh/golang-project-web-api/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working", true, 0))
}

// func (h *HealthHandler) HealthPost(c *gin.Context) {
// 	c.JSON(http.StatusOK, "Working post!")
// 	return
// }

// func (h *HealthHandler) HealthPostById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	c.JSON(http.StatusOK, fmt.Sprintf("Working post by id: %s!", id))
// 	return
// }
