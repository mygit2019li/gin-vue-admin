package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniProgram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MomentsService struct{}

// Categories
func (s *MomentsService) CreateCategory(category miniProgram.MomentsCategory) (err error) {
	err = global.GVA_DB.Create(&category).Error
	return err
}

func (s *MomentsService) DeleteCategory(category miniProgram.MomentsCategory) (err error) {
	err = global.GVA_DB.Delete(&category).Error
	return err
}

func (s *MomentsService) UpdateCategory(category miniProgram.MomentsCategory) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

func (s *MomentsService) GetCategoryList() (list []miniProgram.MomentsCategory, err error) {
	err = global.GVA_DB.Order("sort desc").Find(&list).Error
	return
}

// Copywriting
func (s *MomentsService) CreateCopywriting(cp miniProgram.MomentsCopywriting) (err error) {
	err = global.GVA_DB.Create(&cp).Error
	return err
}

func (s *MomentsService) DeleteCopywriting(cp miniProgram.MomentsCopywriting) (err error) {
	err = global.GVA_DB.Delete(&cp).Error
	return err
}

func (s *MomentsService) UpdateCopywriting(cp miniProgram.MomentsCopywriting) (err error) {
	err = global.GVA_DB.Save(&cp).Error
	return err
}

func (s *MomentsService) GetCopywritingList(info request.PageInfo, categoryId uint) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&miniProgram.MomentsCopywriting{})
	if categoryId > 0 {
		db = db.Where("category_id = ?", categoryId)
	}
	var cps []miniProgram.MomentsCopywriting
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort desc").Find(&cps).Error
	return cps, total, err
}

// Public API for Mini Program
type MomentsPublicResponse struct {
	Category miniProgram.MomentsCategory   `json:"category"`
	Items    []miniProgram.MomentsCopywriting `json:"items"`
}

func (s *MomentsService) GetPublicMomentsList() (res []MomentsPublicResponse, err error) {
	// 初始化res为空切片，确保即使没有数据也返回空数组而不是nil
	res = []MomentsPublicResponse{}
	
	var categories []miniProgram.MomentsCategory
	err = global.GVA_DB.Where("is_hidden = ?", false).Order("sort desc").Find(&categories).Error
	if err != nil {
		return res, err
	}

	for _, cat := range categories {
		var items []miniProgram.MomentsCopywriting
		err = global.GVA_DB.Where("category_id = ? AND is_hidden = ?", cat.ID, false).Order("sort desc").Find(&items).Error
		if err != nil {
			return res, err
		}
		res = append(res, MomentsPublicResponse{
			Category: cat,
			Items:    items,
		})
	}
	return res, nil
}
