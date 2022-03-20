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
	r.POST("/wxlogin", Controllers.WxLogin)
	r.POST("/setVip", Controllers.SetVip)
	r.POST("/getJf", Controllers.GetJf)
	r.POST("/getTask", Controllers.GetTask)
	r.POST("/getLog", Controllers.GetLog)
	r.POST("/getMallClass", Controllers.GetMallClass)
	r.POST("/exchangeGoods", Controllers.ExchangeGoods)
	r.POST("/getDetails", Controllers.GetDetails)
	r.POST("/doneTask", Controllers.DoneTask)
	r.POST("/getCourse", Controllers.GetCourse)
	r.POST("/setCourse", Controllers.SetCourse)
	r.Run()
}
