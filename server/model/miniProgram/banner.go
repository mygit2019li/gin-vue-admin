package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Banner struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;comment:图片名称"`
	Url   string `json:"url" form:"url" gorm:"column:url;comment:图片地址"`
	Link  string `json:"link" form:"link" gorm:"column:link;comment:跳转链接"`
	Sort  int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序"`
	IsShow *bool  `json:"isShow" form:"isShow" gorm:"column:is_show;comment:是否显示"`
}

func (Banner) TableName() string {
	return "mp_banners"
}
