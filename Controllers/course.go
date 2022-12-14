package Controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/svanrj/server/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]string
	// 反序列化
	_ = json.Unmarshal(b, &m)
	var course Models.Course
	courses, err := course.GetCourse(m["type"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "暂无课程",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    courses,
		"message": "获取成功",
	})
}

func SetCourse(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)
	var course Models.Course

	err := course.SetCourse(m["key"], m)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "修改成功",
		})
	}
}
func AddCourse(c *gin.Context) {
	var course Models.Course

	if err := c.ShouldBind(&course); err == nil {
		if err := course.AddCourse(course); err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "message": "添加成功"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "添加失败"})
		}
	} else {
		fmt.Println(err)
	}
}
