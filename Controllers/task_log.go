package Controllers

import (
	"fmt"
	"gitee.com/svanrj/server/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodayTask(c *gin.Context) {
	var todayTask Models.TodayTask
	todayTasks, err := todayTask.GetTodayTaskModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": todayTasks,
	})
}

func GetTaskLog(c *gin.Context) {
	var todayTask Models.TodayTask
	todayTasks, err := todayTask.GetTaskLogModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": todayTasks,
	})
}
func GetTaskSummary(c *gin.Context) {
	var todayTask Models.TodayTask
	TaskSummarys, err := todayTask.GetTaskSummaryModel()

	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": TaskSummarys,
	})
}
