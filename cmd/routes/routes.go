package routes

import (
	"github.com/cayohollanda/gorm-crud/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	{
		v1.GET("persons", controllers.ListPersons)
		v1.GET("persons/:id", controllers.GetPerson)
		v1.POST("persons", controllers.CreatePerson)
		v1.DELETE("persons/:id", controllers.DeletePerson)
		v1.PUT("persons", controllers.UpdatePerson)
	}

	return r
}
