package models

import "github.com/beego/beego/v2/client/orm"

type HotdogType struct {
	Id          int    `orm:"auto;pk"`
	Name        string `orm:"unique"`
	Description string `orm:"type(text)"`
	Price       float64
	Inventory   []*HotdogStock `orm:"reverse(many)"`
	Sample      string         `orm:"type(text)"`
	Human       string
	RIDESH      int
}

type Stand struct {
	Id          int            `orm:"auto;pk"`
	Name        string         `orm:"unique;index"`
	Address     string         `orm:"type(text)"`
	Description string         `orm:"type(text)"`
	Inventory   []*HotdogStock `orm:"reverse(many)"`
}

type HotdogStock struct {
	Id       int         `orm:"auto;pk"`
	Quantity int         `orm:"default(0)"`
	Hotdog   *HotdogType `orm:"rel(fk);on_delete(cascade);index"`
	Stand    *Stand      `orm:"rel(fk);on_delete(cascade);index"`
}

type IndianCricketCouncil struct {
	IccId        int `orm:"auto;pk"`
	PlayerName   string
	PlayerStatus string
	PlayerRole   string
	PlayerAge    int `orm:"default(0)"`
}

func init() {
	orm.RegisterModel(new(HotdogType), new(Stand), new(HotdogStock), new(IndianCricketCouncil))
}
