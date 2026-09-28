package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniProgram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BannerService struct{}

func (bannerService *BannerService) CreateBanner(banner miniProgram.Banner) (err error) {
	err = global.GVA_DB.Create(&banner).Error
	return err
}

func (bannerService *BannerService) DeleteBanner(banner miniProgram.Banner) (err error) {
	err = global.GVA_DB.Delete(&banner).Error
	return err
}

func (bannerService *BannerService) UpdateBanner(banner miniProgram.Banner) (err error) {
	err = global.GVA_DB.Save(&banner).Error
	return err
}

func (bannerService *BannerService) GetBanner(id uint) (banner miniProgram.Banner, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&banner).Error
	return
}

func (bannerService *BannerService) GetBannerInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&miniProgram.Banner{})
	var banners []miniProgram.Banner
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort desc").Find(&banners).Error
	return banners, total, err
}

func (bannerService *BannerService) GetPublicBannerList() (list []miniProgram.Banner, err error) {
	err = global.GVA_DB.Model(&miniProgram.Banner{}).Where("is_show = ?", true).Order("sort desc").Find(&list).Error
	return
}
