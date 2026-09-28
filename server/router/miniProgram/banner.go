package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BannerRouter struct{}

func (s *BannerRouter) InitBannerRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	bannerRouter := Router.Group("banner").Use(middleware.OperationRecord())
	bannerRouterWithoutRecord := Router.Group("banner")
	
	publicBannerRouter := PublicRouter.Group("banner")

	bannerApi := v1.ApiGroupApp.MiniProgramApiGroup.BannerApi
	
	{
		bannerRouter.POST("createBanner", bannerApi.CreateBanner)
		bannerRouter.DELETE("deleteBanner", bannerApi.DeleteBanner)
		bannerRouter.PUT("updateBanner", bannerApi.UpdateBanner)
	}
	{
		bannerRouterWithoutRecord.GET("findBanner", bannerApi.GetBanner)
		bannerRouterWithoutRecord.GET("getBannerList", bannerApi.GetBannerList)
	}
	{
		publicBannerRouter.GET("publicList", bannerApi.PublicBannerList)
	}
}
