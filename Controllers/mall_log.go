package Controllers

import (
	"fmt"
	"gitee.com/svanrj/server/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMallSummary(c *gin.Context) {
	var mallSummary Models.MallSummary
	mallSummarys, err := mallSummary.GetMallSummaryModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    mallSummarys,
		"message": "商品累积兑换次数获取成功",
	})
}

func GetUnconverted(c *gin.Context) {
	var unconverted Models.Unconverted
	unconverteds, count, err := unconverted.GetUnconvertedModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": gin.H{
			"list":  unconverteds,
			"count": count,
		},
		"message": "未兑换商品获取",
	})
}
func GetMallLog(c *gin.Context) {
	var todayTask Models.MallLog
	mallLogs, err := todayTask.GetMallLogModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": mallLogs,
	})
}

func UpdateUnconverted(c *gin.Context) {

	var detail Models.Detail

	id := c.Param("id")

	err := detail.UpdateUnconvertedModel(id)

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "操作成功",
	})
}
