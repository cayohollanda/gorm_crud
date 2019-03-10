package controllers

import (
	"github.com/cayohollanda/gorm_crud/cmd/models"
	"github.com/cayohollanda/gorm_crud/cmd/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.AddUser(&user)
	if err != nil {
		utils.ResponseJSON(c, 404, nil)
	} else {
		utils.ResponseJSON(c, 200, user)
	}
}
