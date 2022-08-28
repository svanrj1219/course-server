package Router

import (
	"gitee.com/svanrj/server/Controllers"
	"gitee.com/svanrj/server/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//初始化项目
	r := gin.Default()
	//使用中间件
	r.Use(Middlewares.Cors())
	//处理请求
	r.POST("/login", Controllers.Login)
	adminGroup := r.Group("/admin")
	adminGroup.Use(Middlewares.JWTAuthMiddleware())
	{
		adminGroup.POST("/setCourse", Controllers.SetCourse)
		adminGroup.GET("/getTodayTask", Controllers.GetTodayTask)
		adminGroup.GET("/getHistoryJf", Controllers.GetHistoryJf)
		adminGroup.GET("/getTaskSummary", Controllers.GetTaskSummary)
		adminGroup.GET("/getMallSummary", Controllers.GetMallSummary)
		adminGroup.GET("/unconverted", Controllers.GetUnconverted)
		adminGroup.GET("/taskLog", Controllers.GetTaskLog)
		adminGroup.GET("/task", Controllers.GetAllTask)
		adminGroup.GET("/mallLog", Controllers.GetMallLog)
		adminGroup.PUT("/unconverted/:id", Controllers.UpdateUnconverted)
		adminGroup.POST("/course", Controllers.AddCourse)
	}
	r.POST("/getCourse", Controllers.GetCourse)
	r.POST("/wxlogin", Controllers.WxLogin)
	r.POST("/setVip", Controllers.SetVip)
	r.POST("/getJf", Controllers.GetJf)
	r.POST("/getTask", Controllers.GetTask)
	r.POST("/getLog", Controllers.GetLog)
	r.POST("/getMallClass", Controllers.GetMallClass)
	r.POST("/exchangeGoods", Controllers.ExchangeGoods)
	r.POST("/getDetails", Controllers.GetDetails)
	r.POST("/doneTask", Controllers.DoneTask)
	err := r.Run()
	if err != nil {
		return
	}
}
