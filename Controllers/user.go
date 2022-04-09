package Controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/svanrj/server/Models"
	"gitee.com/svanrj/server/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Vip struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
}

func Login(c *gin.Context) {
	var user Models.UserInfo
	if err := c.ShouldBind(&user); err == nil {
		password := user.Password
		username := user.Username
		password = utils.Md5Test(password)
		userInfo, err := user.UserModel(username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "用户名不存在",
			})
			return
		}
		if password == userInfo.Password {
			tokenString, _ := utils.GenToken(username)
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "登录成功",
				"token":   tokenString,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "密码错误",
			})
		}
	}

}
func WxLogin(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]string
	// 反序列化
	_ = json.Unmarshal(b, &m)
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wxa36eef4162bac240&secret=b8ea8be6519384bca46f60fee46531d4&js_code=%v&grant_type=authorization_code", m["code"])

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get err:%v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	var user = new(Models.WxUser)

	_ = json.Unmarshal(body, &m)

	user.Username = m["username"]
	user.Avatar = m["avatarUrl"]
	user.Openid = m["openid"]

	wxu, err := user.WxuserModel(user)

	c.JSON(http.StatusOK, gin.H{
		"code":        1,
		"isVip":       wxu.Isvip,
		"id":          wxu.ID,
		"integration": wxu.Integration,
	})
}

func SetVip(c *gin.Context) {

	var v Vip
	c.ShouldBind(&v)

	var user Models.WxUser
	if v.Code == "112604" {
		err := user.SetVipModel(v.Id)

		if err != nil {
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "欢迎使用",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "邀请码错误",
		})
	}
}
func GetJf(c *gin.Context) {

	var v Vip
	c.ShouldBind(&v)

	var user Models.WxUser

	jf, err := user.GetJfModel(v.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"jf":   jf,
	})
}
func GetHistoryJf(c *gin.Context) {
	var historyJf Models.HistoryJf
	historyJfs, err := historyJf.GetHistoryJfModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": historyJfs,
	})
}
