package Router

import (
	"gitee.com/svanrj/server/Controllers"
	"gitee.com/svanrj/server/Middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	//初始化项目
	r := gin.Default()
	//使用中间件
	r.Use(Middlewares.Cors())
	//处理请求
	r.POST("/login",Controllers.Login)
	r.GET("/home",Middlewares.JWTAuthMiddleware() , func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"成功",
		})
	})
	r.Run()
}

