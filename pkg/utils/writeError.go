package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestError(c *gin.Context) {
	
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "Invalid request",
	})
}

func InternalServerError(c *gin.Context) {
	
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Internal server error",
	})
}

func UnauthorizedError(c *gin.Context) {
	
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}

func ForbiddenError(c *gin.Context) {
	
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": "Forbidden",
	})
}

func NotFoundError(c *gin.Context) {
	
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "Not found",
	})
}

