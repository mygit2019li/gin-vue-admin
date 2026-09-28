package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type MomentsCategory struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:分类名称"`
	Sort int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序"`
	IsHidden bool   `json:"isHidden" form:"isHidden" gorm:"column:is_hidden;comment:是否隐藏"`
}

func (MomentsCategory) TableName() string {
	return "mp_moments_categories"
}

type MomentsCopywriting struct {
	global.GVA_MODEL
	CategoryID uint   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类ID"`
	Text       string `json:"text" form:"text" gorm:"column:text;type:text;comment:文案内容"`
	Images     string `json:"images" form:"images" gorm:"column:images;type:text;comment:图片列表JSON"`
	Sort       int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序"`
	IsHidden   bool   `json:"isHidden" form:"isHidden" gorm:"column:is_hidden;comment:是否隐藏"`
}

func (MomentsCopywriting) TableName() string {
	return "mp_moments_copywriting"
}
