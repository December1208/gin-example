package server

import (
	"github.com/gin-gonic/gin"
	"go-example/routers"
)

func Init() {
	gin.SetMode(gin.DebugMode)
	r := CreateEngine()
	err := r.Run(":8855")
	if err != nil {
		panic(err.Error())
	}

}

func CreateEngine() *gin.Engine {
	eng := gin.New()
	//eng.Use(sentryGin.New(sentryGin.Options{}))

	//if !util2.IsGoTest() {
	//	eng.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//		// 定制日志格式
	//		return fmt.Sprintf("[%s] - [%s] [%s] [%s] %d %s %s\"\n",
	//			param.TimeStamp.Format(time.RFC3339),
	//			param.ClientIP,
	//			param.Method,
	//			param.Path,
	//			param.StatusCode,
	//			param.Latency,
	//			param.ErrorMessage,
	//		)
	//	}))
	//}
	eng.Use(gin.Recovery())
	//eng.Use(ginTrace.Middleware("go-gin-gorm"))
	// router.Use(middleware.DBMiddleware())
	//middleware.RegMiddleware(eng)
	routers.InitRouters(eng)
	return eng
}