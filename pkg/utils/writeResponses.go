package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func createdResponse(c *gin.Context, data interface{}) {
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Created",
// 		"data": data,
// 	})
// }

func CreatedResponse(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created successfully",
	})
}

func OkResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": data,
	})
}

// func noContentResponse(c *gin.Context) {
// 	c.JSON(http.StatusNoContent, gin.H{
// 		"message": "No content",
// 	})
// }

// func badRequestResponse(c *gin.Context) {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message": "Invalid request",
// 	})
// }

// func internalServerErrorResponse(c *gin.Context) {
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"message": "Internal server error",
// 	})
// }

