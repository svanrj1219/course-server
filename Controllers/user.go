package Controllers

import (
	"gitee.com/svanrj/server/Databases"
	"gitee.com/svanrj/server/Models"
	"gitee.com/svanrj/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){
	var user Models.UserInfo
	if err:=c.ShouldBind(&user);err==nil{
		password:=user.Password
		username:=user.Username
		password=utils.Md5Test(password)
		if err:=Databases.DB.Where("username=?",username).First(&user).Error;err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"code":1,
				"message":"用户名不存在",
			})
			return
		}
		if password==user.Password{
			c.JSON(http.StatusOK,gin.H{
				"code":0,
				"message":"登录成功",
				"data":user.Username,
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"code":1,
				"message":"密码错误",
			})
		}
	}

}