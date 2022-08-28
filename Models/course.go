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
	Sort       int64  `json:"sort"`
}

func (cou *Course) GetCourse(t string) (c []Course, err error) {
	if t == "全部" {
		var courses []Course
		err = Databases.DB.Order("sort").Find(&courses).Error
		return courses, err
	} else {
		var courses []Course
		err = Databases.DB.Order("sort").Where("belongto=?", t).Find(&courses).Error
		return courses, err
	}
}

func (cou *Course) SetCourse(id interface{}, c map[string]interface{}) (err error) {
	delete(c, "key")
	err = Databases.DB.Model(&cou).Where("id = ?", id).Updates(c).Error
	return err
}
func (cou *Course) AddCourse(course Course) (err error) {
	err = Databases.DB.Create(&course).Error
	return err
}
