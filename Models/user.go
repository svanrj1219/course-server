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
