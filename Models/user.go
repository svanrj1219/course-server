package Models

import (
	"gitee.com/svanrj/server/Databases"
)

type UserInfo struct {
	ID       uint
	Username string
	Password string
}
type WxUser struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Openid      string `json:"openid"`
	Avatar      string `json:"avatar"`
	Isvip       int64  `json:"isvip"`
	Integration int    `json:"integration"`
}
type HistoryJf struct {
	Name string `json:"name"`
	Jf   int64  `json:"jf"`
}

//封装代码
func (user *UserInfo) UserModel(username string) (u *UserInfo, err error) {
	err = Databases.DB.Where("username=?", username).First(&user).Error
	return user, err
}

func (wxUser *WxUser) WxuserModel(info *WxUser) (wxu *WxUser, err error) {
	err = Databases.DB.Where("openid=?", info.Openid).First(&wxUser).Error
	if err != nil {
		Databases.DB.Create(info)
	}
	return wxUser, err
}
func (wxUser *WxUser) SetVipModel(id int) (err error) {
	wxUser.ID = uint(id)
	err = Databases.DB.Model(&wxUser).Update("isvip", 1).Error
	return err
}
func (wxUser *WxUser) GetJfModel(id int) (jf int, err error) {
	err = Databases.DB.Where("id=?", id).First(&wxUser).Error
	jf = wxUser.Integration
	return jf, err
}
func (historyJf HistoryJf) GetHistoryJfModel() (i []HistoryJf, err error) {
	var wxUser []WxUser
	err = Databases.DB.Where("isvip =?", 1).Find(&wxUser).Error

	var HistoryJfs = make([]HistoryJf, len(wxUser))
	for index, value := range wxUser {
		var sum int64 = 0
		var jfs []int64
		err = Databases.DB.Table("integration_logs").Where("uid=?", value.ID).Pluck("integration", &jfs).Error
		HistoryJfs[index].Name = value.Username
		for _, val := range jfs {
			sum = sum + val
		}
		HistoryJfs[index].Jf = sum

	}

	return HistoryJfs, err
}
