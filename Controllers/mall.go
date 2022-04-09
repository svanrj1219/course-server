package Controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/svanrj/server/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMallClass(c *gin.Context) {
	var mall Models.MallClasse
	malls, err := mall.GetMallModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    malls,
		"message": "所有商品获取成功",
	})
}
func ExchangeGoods(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var mall Models.Merchandise
	var m Models.M
	_ = json.Unmarshal(b, &m)
	message, err := mall.ExchangeGoodsModel(m)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": message,
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": message,
	})
}
func GetDetails(c *gin.Context) {
	var detail Models.Detail
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]int
	_ = json.Unmarshal(b, &m)

	details, err := detail.GetDetailsModel(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    details,
		"message": "兑换详情获取成功",
	})
}
