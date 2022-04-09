package Models

import (
	"fmt"
	"gitee.com/svanrj/server/Databases"
	"gopkg.in/gomail.v2"
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

func (mer *Merchandise) ExchangeGoodsModel(m M) (message string, err error) {
	var wxUser WxUser
	var d Detail
	var mers Merchandise

	err = Databases.DB.Where("id=?", m.Uid).First(&wxUser).Error

	if wxUser.Integration < m.RequiredPoints {
		return "当前积分不足", err
	} else {
		wxUser.ID = uint(m.Uid)

		d.Uid = m.Uid
		d.MerchandiseId = m.ID
		d.CostCredits = m.RequiredPoints
		d.IsExchange = 0
		d.Date = time.Now().Format("2006-01-02 15:04:05")

		Databases.DB.Create(&d)
		err = Databases.DB.Model(&wxUser).Update("integration", wxUser.Integration-m.RequiredPoints).Error
		Databases.DB.Where("id = ?", d.MerchandiseId).First(&mers)

		m := gomail.NewMessage()

		//发送人
		m.SetHeader("From", "2893275727@qq.com")
		//接收人
		m.SetHeader("To", "17633162305@163.com")
		//抄送人
		//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
		//主题
		m.SetHeader("Subject", "积分兑换成功")
		//内容
		content := fmt.Sprintf(` <div style="text-align: center;">
        <table style=" border: 1px solid #ebeef5; border-collapse:collapse; width:500px; margin: 0 auto; color: #606266;">
            <tr style=" border-bottom: 1px solid #ebeef5;">
                <td style="border-bottom: 1px solid #ebeef5; padding: 12px 0;" colspan="2">
                    积分兑换成功
                </td>
            </tr>
            <tr style="border-bottom: 1px solid #ebeef5; color: #909399;">
                <th style="border-bottom: 1px solid #ebeef5; font-weight: 500; padding: 12px 0;">商品名称</th>
                <th style="border-bottom: 1px solid #ebeef5; font-weight: 500; padding: 12px 0;">使用积分</th>
            </tr>
            <tr style="border-bottom: 1px solid #ebeef5;">
                <td style="border-bottom: 1px solid #ebeef5; padding: 12px 0;">%v</td>
                <td style="border-bottom: 1px solid #ebeef5; padding: 12px 0;">%v</td>
            </tr>
            <tr style=" border-bottom: 1px solid #ebeef5;">
                <td style="border-bottom: 1px solid #ebeef5; padding: 12px 0;" colspan="2">
                    快去给宝宝兑换吧
                </td>
            </tr>
        </table>
    </div>`, mers.Name, d.CostCredits)
		m.SetBody("text/html", content)
		//附件
		//m.Attach("./myIpPic.png")

		//拿到token，并进行连接,第4个参数是填授权码
		d := gomail.NewDialer("smtp.qq.com", 587, "2893275727@qq.com", "ynmhdiqcvztnddhe")

		// 发送邮件
		if err := d.DialAndSend(m); err != nil {
			fmt.Printf("DialAndSend err %v:", err)
			panic(err)
		}

		return "兑换成功", err
	}
}
