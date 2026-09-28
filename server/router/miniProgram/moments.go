package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MomentsRouter struct{}

func (s *MomentsRouter) InitMomentsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	momentsRouter := Router.Group("moments").Use(middleware.OperationRecord())
	momentsRouterWithoutRecord := Router.Group("moments")
	
	publicMomentsRouter := PublicRouter.Group("moments")

	momentsApi := v1.ApiGroupApp.MiniProgramApiGroup.MomentsApi
	
	{
		momentsRouter.POST("createCategory", momentsApi.CreateCategory)
		momentsRouter.DELETE("deleteCategory", momentsApi.DeleteCategory)
		momentsRouter.PUT("updateCategory", momentsApi.UpdateCategory)
		
		momentsRouter.POST("createCopywriting", momentsApi.CreateCopywriting)
		momentsRouter.DELETE("deleteCopywriting", momentsApi.DeleteCopywriting)
		momentsRouter.PUT("updateCopywriting", momentsApi.UpdateCopywriting)
	}
	{
		momentsRouterWithoutRecord.GET("getCategoryList", momentsApi.GetCategoryList)
		momentsRouterWithoutRecord.GET("getCopywritingList", momentsApi.GetCopywritingList)
	}
	{
		publicMomentsRouter.GET("publicList", momentsApi.PublicMomentsList)
	}
}
