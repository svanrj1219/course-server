package Models

import (
	"gitee.com/svanrj/server/Databases"
	"time"
)

type MallClasse struct {
	ID   int           `json:"id"`
	Name string        `json:"name"`
	Mall []Merchandise `json:"mall"`
}
type Merchandise struct {
	ID             int    `json:"id"`
	Cid            int    `json:"cid"`
	Name           string `json:"name"`
	PictureUrl     string `json:"picture_url"`
	RequiredPoints int    `json:"required_points"`
}
type Detail struct {
	ID            int    `json:"id"`
	Uid           int    `json:"uid"`
	MerchandiseId int    `json:"merchandise_id"`
	Date          string `json:"date"`
	CostCredits   int    `json:"cost_credits"`
	IsExchange    int    `json:"is_exchange"`
	Name          string `json:"name" gorm:"-"`
}
type M struct {
	ID             int `json:"id"`
	Uid            int `json:"uid"`
	RequiredPoints int `json:"required_points"`
}

//封装代码
func (mall *MallClasse) GetMallModel() (m []MallClasse, err error) {
	var malls []MallClasse
	err = Databases.DB.Find(&malls).Error
	for index, value := range malls {
		var mers []Merchandise
		Databases.DB.Where("cid = ?", value.ID).Find(&mers)
		malls[index].Mall = mers
	}
	return malls, err
}
func (Merchandise *Merchandise) ExchangeGoodsModel(m M) (message string, err error) {
	var wxUser WxUser
	var d Detail

	err = Databases.DB.Where("id=?", m.Uid).First(&wxUser).Error

	if wxUser.Integration < m.RequiredPoints {
		return "当前积分不足", err
	} else {
		wxUser.ID = uint(m.Uid)

		d.Uid = m.Uid
		d.MerchandiseId = m.ID
		d.CostCredits = m.RequiredPoints
		d.IsExchange = 0
		d.Date = time.Now().Format("2006-01-02")

		Databases.DB.Create(&d)
		err = Databases.DB.Model(&wxUser).Update("integration", wxUser.Integration-m.RequiredPoints).Error

		return "兑换成功", err
	}
}
func (details *Detail) GetDetailsModel(id int) (detailsList []Detail, err error) {
	err = Databases.DB.Order("id desc").Where("uid=?", id).Find(&detailsList).Error

	for index, value := range detailsList {
		var mers Merchandise
		Databases.DB.Where("id = ?", value.MerchandiseId).First(&mers)
		detailsList[index].Name = mers.Name
	}
	return detailsList, err
}
