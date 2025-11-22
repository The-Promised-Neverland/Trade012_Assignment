package middleware

import (
	"net/http"
	"strconv"

	"github.com/ThePromisedNeverland/021trade/internal/services"
	"github.com/gin-gonic/gin"
)

func UserExistsMiddleware(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdStr := c.Param("userId")

		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid user_id",
			})
			return
		}

		_, err = userService.GetUser(c.Request.Context(), userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
