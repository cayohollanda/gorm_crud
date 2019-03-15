package main

import (
	"log"

	"github.com/cayohollanda/gorm_crud/cmd/config"
	"github.com/cayohollanda/gorm_crud/cmd/models"
	"github.com/cayohollanda/gorm_crud/cmd/routes"
	"github.com/cayohollanda/gorm_crud/cmd/utils"
	"github.com/jinzhu/gorm"
)

func main() {
	log.Println("[INFO] Iniciando aplicação")

	initializeDB()

	config.DB.AutoMigrate(&models.Person{}, &models.User{})
	log.Println("[INFO] AutoMigrate feito com sucesso")

	r := routes.GetRoutes()

	log.Println("[INFO] Aplicação iniciada")
	r.Run(":8000")
}

func initializeDB() {
	var err error
	config.DB, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_crud?charset=utf8&parseTime=True&loc=Local")
	utils.CheckErr(err)
	log.Println("[INFO] Conectado no banco de dados")

	var checkIfAdminExists models.User
	checkIfAdminExists.Username = "admin"
	if err := models.FindUserByUsername(&checkIfAdminExists); err != nil {
		createAdmin()
	}
}

func createAdmin() {
	adminUser := models.User{
		Username: "admin",
		Password: "admin",
		Role:     "ADMIN",
	}
	if err := models.AddUser(&adminUser); err != nil {
		log.Println("[INFO] Erro ao criar usuário administrador")
	} else {
		log.Println("[INFO] Usuário administrador criado com sucesso")
	}
}
