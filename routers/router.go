
package routers

import (
	"github.com/gin-gonic/gin"
	"go-example/controllers"
)

func InitRouters(eng *gin.Engine)  {
	api := eng.Group("/api")
	{
		controller := new(controllers.Controller)
		api.GET("/index", controller.Index)
	}
}