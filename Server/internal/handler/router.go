package handler

import "github.com/gin-gonic/gin"

func SetupRouter(h *Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", Ping)
	
	return router
}
