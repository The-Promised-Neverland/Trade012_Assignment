package handlers

import (
	"net/http"
	"strconv"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
	log     *logger.Logger
}

func NewUserHandler(s *services.UserService, log *logger.Logger) *UserHandler {
	return &UserHandler{
		service: s,
		log:     log,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		h.log.Errorf("Error parsing userId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	u, err := h.service.GetUser(c, userID)
	if err != nil {
		h.log.Errorf("User not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
