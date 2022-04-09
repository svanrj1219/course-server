package Models

import (
	"fmt"
	"gitee.com/svanrj/server/Databases"
	"time"
)

type IntegrationLog struct {
	ID          int    `json:"id"`
	Uid         int    `json:"uid"`
	Integration int    `json:"integration"`
	Date        string `json:"date"`
	Remarks     string `json:"remarks"`
}
type Task struct {
	ID        int    `json:"id"`
	Uid       int    `json:"uid"`
	Content   string `json:"content"`
	GetPoints int    `json:"get_points"`
}

//封装代码
func (task Task) GetTaskModel(id int) (t []Task, err error) {
	var tasks []Task
	err = Databases.DB.Where("uid = ?", id).Find(&tasks).Error

	return tasks, err
}

func (task Task) GetLogModel(m map[string]int) (i []IntegrationLog, err error) {
	var Logs []IntegrationLog
	err = Databases.DB.Order("id desc").Where("uid = ?", m["uid"]).Limit(m["size"]).Offset((m["page"] - 1) * m["size"]).Find(&Logs).Error

	return Logs, err
}
func (task Task) DoneTaskModel() (err error) {
	var integrationLog IntegrationLog

	integrationLog.Integration = task.GetPoints
	integrationLog.Uid = task.Uid
	integrationLog.Remarks = task.Content
	integrationLog.Date = time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(integrationLog.Date)

	Databases.DB.Create(&integrationLog)

	var wxUser WxUser
	err = Databases.DB.Where("id=?", task.Uid).First(&wxUser).Error
	wxUser.ID = uint(task.Uid)
	err = Databases.DB.Model(&wxUser).Update("integration", wxUser.Integration+task.GetPoints).Error
	return err
}
