package Models

import (
	"gitee.com/svanrj/server/Databases"
)

type Course struct {
	ID         int64  `json:"key"`
	Coursename string `json:"courseName"`
	Time       string `json:"time"`
	Teacher    string `json:"teacher"`
	Address    string `json:"address"`
	Per        string `json:"per"`
	Week       string `json:"week"`
	Belongto   string `json:"belongto"`
}

func (cou *Course) GetCourse() (c []Course, err error) {

	var courses []Course
	err = Databases.DB.Find(&courses).Error

	return courses, err
}
