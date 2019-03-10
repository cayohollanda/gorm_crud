package routes

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/cayohollanda/gorm_crud/cmd/controllers"
	"github.com/cayohollanda/gorm_crud/cmd/models"
	"github.com/cayohollanda/gorm_crud/cmd/utils"
	"github.com/gin-gonic/gin"
)

type ErrorMap struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "user",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"user": v.Username,
					"role": v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims["user"].(string),
			}
		},
		Authenticator: Authenticator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/createUser", controllers.CreateUser)

	// Auth Group
	{
		authGroup := r.Group("/auth")
		authGroup.POST("/login", authMiddleware.LoginHandler)
		authGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	// v1 Group
	v1 := r.Group("/v1")
	v1.Use(authMiddleware.MiddlewareFunc())

	// Persons Group
	{
		personsGroup := v1.Group("/persons")
		personsGroup.Use(roleCheckerMiddleware("ADMIN"))

		personsGroup.GET("/", controllers.ListPersons)
		personsGroup.GET("/:id", controllers.GetPerson)
		personsGroup.POST("/", controllers.CreatePerson)
		personsGroup.DELETE("/:id", controllers.DeletePerson)
		personsGroup.PUT("/", controllers.UpdatePerson)
	}

	return r
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var userLogin models.User
	if err := c.ShouldBind(&userLogin); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	password := userLogin.Password
	if err := models.FindUserByUsername(&userLogin); err != nil {
		return nil, jwt.ErrFailedAuthentication
	} else {
		var hash utils.Hash
		if err := hash.Compare(userLogin.Password, password); err != nil {
			return nil, jwt.ErrFailedAuthentication
		}
		return &models.User{
			Username: userLogin.Username,
			Role:     userLogin.Role,
		}, nil
	}

}

func getRoleFromContext(c *gin.Context) (role string) {
	claims := jwt.ExtractClaims(c)

	role, ok := claims["role"].(string)
	if !ok {
		c.AbortWithStatusJSON(403, &ErrorMap{
			Code:    403,
			Message: "You may not have permissions to access this resource",
		})
	}

	return role
}

func roleCheckerMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		contextRole := getRoleFromContext(c)
		if contextRole == role {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(403, &ErrorMap{
			Code:    403,
			Message: "You may not have permissions to access this resource",
		})
	}
}
