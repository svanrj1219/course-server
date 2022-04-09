package Models

import (
	"gitee.com/svanrj/server/Databases"
	"time"
)

type TodayTask struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Value int    `json:"value"`
}

type TaskLog struct {
	IntegrationLog
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func (todayTask TodayTask) GetTodayTaskModel() (i []TodayTask, err error) {
	var wxUser []WxUser
	err = Databases.DB.Where("isvip =?", 1).Find(&wxUser).Error

	var TodayTasks = make([]TodayTask, len(wxUser))
	var Date = time.Now().Format("2006-01-02")
	for index, value := range wxUser {
		var count int
		var Logs []IntegrationLog
		err = Databases.DB.Where("uid=? and date like ?", value.ID, Date+"%").Find(&Logs).Count(&count).Error
		TodayTasks[index].Name = value.Username
		TodayTasks[index].Count = count
	}

	return TodayTasks, err
}
func (todayTask TodayTask) GetTaskLogModel() (r []TaskLog, err error) {

	err = Databases.DB.Table("integration_logs as i").Select("i.*, u.username,u.avatar").Joins("left join wx_users as u  on	 i.uid=u.id").Order("id desc").Find(&r).Error
	return r, err
}
func (todayTask TodayTask) GetTaskSummaryModel() (i []TodayTask, err error) {

	var TaskSummarys []TodayTask
	Databases.DB.Raw(`SELECT content as name,( SELECT count(*) FROM integration_logs WHERE content = integration_logs.remarks ) AS value FROM tasks;`).Scan(&TaskSummarys)

	return TaskSummarys, err
}
