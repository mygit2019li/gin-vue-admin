package miniProgram

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BannerApi
	MomentsApi
}

var bannerService = service.ServiceGroupApp.MiniProgramServiceGroup.BannerService
var momentsService = service.ServiceGroupApp.MiniProgramServiceGroup.MomentsService
