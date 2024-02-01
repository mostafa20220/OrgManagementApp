package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/pkg/api/models"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/pkg/utils"
)


func Signup(c *gin.Context) {
	
	var newUser models.User
	err :=  c.BindJSON(&newUser)

	if err != nil {
		utils.BadRequestError(c)
		return
	}

	

	utils.CreatedResponse(c)
}