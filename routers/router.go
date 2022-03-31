
package routers

import (
	"github.com/gin-gonic/gin"
	"go-example/controllers"
)

func InitRouters(eng *gin.Engine)  {
	v2api := eng.Group("/v2/api")
	{
		controller := new(controllers.Controller)
		v2api.GET("/index", controller.Index)
	}
}