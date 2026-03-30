package handler

import (
	"github.com/JerryDdie77/Hezze/Server/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *Handler {
	return &Handler{userService: userService}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"status" : "ok"})
}