package controllers

import (
	"github.com/cayohollanda/gorm-crud/cmd/models"
	"github.com/cayohollanda/gorm-crud/cmd/utils"
	"github.com/gin-gonic/gin"
)

func ListPersons(c *gin.Context) {
	var persons []models.Person
	err := models.GetAllPersons(&persons)
	if err != nil {
		utils.ResponseJSON(c, 404, nil)
	} else {
		utils.ResponseJSON(c, 200, persons)
	}
}

func GetPerson(c *gin.Context) {
	var person models.Person
	id := c.Param("id")

	err := models.FindPerson(&person, id)
	if err != nil {
		utils.ResponseJSON(c, 404, nil)
	} else {
		utils.ResponseJSON(c, 200, person)
	}
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)
	err := models.AddPerson(&person)
	if err != nil {
		utils.ResponseJSON(c, 404, nil)
	} else {
		utils.ResponseJSON(c, 200, person)
	}
}

func DeletePerson(c *gin.Context) {
	var person models.Person
	id := c.Param("id")
	err := models.DeletePerson(&person, id)
	if err != nil {
		utils.ResponseJSON(c, 404, nil)
	} else {
		utils.ResponseJSON(c, 200, person)
	}
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)
	err := models.UpdatePerson(&person)
	if err != nil {
		utils.ResponseJSON(c, 404, nil)
	} else {
		utils.ResponseJSON(c, 200, person)
	}
}
