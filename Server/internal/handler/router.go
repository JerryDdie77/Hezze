package handler

import "github.com/gin-gonic/gin"

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/health", Ping)

	api:= r.Group("/api")
	{
		api.GET("/users/:id")
	}
	return r
}
