package Models

import "gitee.com/svanrj/server/Databases"

type UserInfo struct {
	ID       uint
	Username string
	Password string
}

//封装代码
func (user *UserInfo) UserModel(username string) (u *UserInfo, err error) {
	err = Databases.DB.Where("username=?", username).First(&user).Error
	return user, err
}
