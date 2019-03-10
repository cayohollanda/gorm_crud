package cmd

import (
	"log"

	"github.com/cayohollanda/gorm-crud/cmd/config"
	"github.com/cayohollanda/gorm-crud/cmd/models"
	"github.com/cayohollanda/gorm-crud/cmd/routes"
	"github.com/cayohollanda/gorm-crud/cmd/utils"
	"github.com/jinzhu/gorm"
)

func Initialize() {
	log.Println("[INFO] Iniciando aplicação")
	var err error
	config.DB, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_crud?charset=utf8&parseTime=True&loc=Local")
	utils.CheckErr(err)
	defer config.DB.Close()
	log.Println("[INFO] Conectado no banco de dados")

	config.DB.AutoMigrate(&models.Person{}, &models.User{})
	log.Println("[INFO] AutoMigrate feito com sucesso")

	r := routes.GetRoutes()

	log.Println("[INFO] Aplicação iniciada")
	r.Run(":8000")
}
