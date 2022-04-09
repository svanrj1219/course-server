package Models

import "gitee.com/svanrj/server/Databases"

type M struct {
	ID             int `json:"id"`
	Uid            int `json:"uid"`
	RequiredPoints int `json:"required_points"`
}
type MallSummary struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type Unconverted struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	ID   int    `json:"id"`
}

type MallLog struct {
	Detail
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
}

func (mallSummary MallSummary) GetMallSummaryModel() (i []MallSummary, err error) {
	var MallSummarys []MallSummary
	Databases.DB.Raw(`SELECT name,( SELECT count(*) FROM details WHERE merchandise_id = merchandises.id ) AS count FROM merchandises;`).Scan(&MallSummarys)
	return MallSummarys, err
}

func (unconverted Unconverted) GetUnconvertedModel() (i []Unconverted, count int, err error) {
	var detail []Detail
	err = Databases.DB.Where("is_exchange=?", 0).Find(&detail).Count(&count).Error

	var unconverteds = make([]Unconverted, len(detail))
	for index, value := range detail {
		var merchandise Merchandise
		err = Databases.DB.Where("id=?", value.MerchandiseId).First(&merchandise).Error

		unconverteds[index].Name = merchandise.Name
		unconverteds[index].Url = merchandise.PictureUrl
		unconverteds[index].ID = value.ID
	}
	return unconverteds, count, err
}

func (detail Detail) UpdateUnconvertedModel(id string) (err error) {
	var d Detail
	err = Databases.DB.Model(&d).Where("id = ?", id).Update("is_exchange", 1).Error
	return err
}

func (detail Detail) GetDetailsModel(m map[string]int) (detailsList []Detail, err error) {
	err = Databases.DB.Order("id desc").Where("uid=?", m["id"]).Limit(m["size"]).Offset((m["page"] - 1) * m["size"]).Find(&detailsList).Error

	for index, value := range detailsList {
		var mers Merchandise
		Databases.DB.Where("id = ?", value.MerchandiseId).First(&mers)
		detailsList[index].Name = mers.Name
	}
	return detailsList, err
}
func (mallLog MallLog) GetMallLogModel() (r []MallLog, err error) {

	err = Databases.DB.Table("details as d").Select("d.*, u.username,u.avatar,m.name").Joins("left join wx_users as u  on d.uid=u.id").Joins("left join merchandises as m  on d.merchandise_id=m.id").Order("id desc").Find(&r).Error
	return r, err
}
