package Controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/svanrj/server/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTask(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]int
	// 反序列化
	_ = json.Unmarshal(b, &m)
	var task Models.Task
	tasks, err := task.GetTaskModel(m["id"])

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": tasks,
	})
}
func GetLog(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]int
	// 反序列化
	_ = json.Unmarshal(b, &m)
	var task Models.Task
	Logs, err := task.GetLogModel(m)

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": Logs,
	})
}
func DoneTask(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var task Models.Task
	_ = json.Unmarshal(b, &task)
	err := task.DoneTaskModel()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "任务完成失败",
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "任务完成成功",
	})
}
