package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JerryDdie77/Hezze/Server/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	ctx := c.Request.Context()

	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
			return
		case errors.Is(err, domain.ErrInternal):
			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
			return
		case errors.Is(err, domain.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error" : err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, user)
}
