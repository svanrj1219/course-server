package Controllers

import (
	"gitee.com/svanrj/server/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(c *gin.Context) {
	var course Models.Course
	courses, err := course.GetCourse()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "暂无课程",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": courses,
	})
}
