package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniProgram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BannerApi struct{}

func (bannerApi *BannerApi) CreateBanner(c *gin.Context) {
	var banner miniProgram.Banner
	_ = c.ShouldBindJSON(&banner)
	if err := bannerService.CreateBanner(banner); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (bannerApi *BannerApi) DeleteBanner(c *gin.Context) {
	var banner miniProgram.Banner
	_ = c.ShouldBindJSON(&banner)
	if err := bannerService.DeleteBanner(banner); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (bannerApi *BannerApi) UpdateBanner(c *gin.Context) {
	var banner miniProgram.Banner
	_ = c.ShouldBindJSON(&banner)
	if err := bannerService.UpdateBanner(banner); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (bannerApi *BannerApi) GetBanner(c *gin.Context) {
	var banner miniProgram.Banner
	_ = c.ShouldBindQuery(&banner)
	if resBanner, err := bannerService.GetBanner(banner.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"banner": resBanner}, c)
	}
}

func (bannerApi *BannerApi) GetBannerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bannerService.GetBannerInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (bannerApi *BannerApi) PublicBannerList(c *gin.Context) {
	if list, err := bannerService.GetPublicBannerList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
