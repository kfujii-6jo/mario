package routers

import (
	"github.com/gin-gonic/gin"
	"mario/internal/app/controllers"
)


func SetupRouter() *gin.Engine {
    router := gin.Default()

	internalApi := router.Group("/internal/api")
	{
		v1 := internalApi.Group("/v1")
		{
			v1.GET("/user", controllers.GetUsers)
			v1.POST("/user", controllers.CreateUser)
		}
	}
	externalApi := router.Group("/api")
	{
		v1 := externalApi.Group("/v1")
		{
			v1.POST("/login", controllers.Login)
		}
	}

    return router
}
