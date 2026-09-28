package miniProgram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniProgram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type MomentsApi struct{}

// Categories
func (api *MomentsApi) CreateCategory(c *gin.Context) {
	var category miniProgram.MomentsCategory
	_ = c.ShouldBindJSON(&category)
	if err := momentsService.CreateCategory(category); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (api *MomentsApi) DeleteCategory(c *gin.Context) {
	var category miniProgram.MomentsCategory
	_ = c.ShouldBindJSON(&category)
	if err := momentsService.DeleteCategory(category); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (api *MomentsApi) UpdateCategory(c *gin.Context) {
	var category miniProgram.MomentsCategory
	_ = c.ShouldBindJSON(&category)
	if err := momentsService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (api *MomentsApi) GetCategoryList(c *gin.Context) {
	if list, err := momentsService.GetCategoryList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}

// Copywriting
func (api *MomentsApi) CreateCopywriting(c *gin.Context) {
	var cp miniProgram.MomentsCopywriting
	_ = c.ShouldBindJSON(&cp)
	if err := momentsService.CreateCopywriting(cp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (api *MomentsApi) DeleteCopywriting(c *gin.Context) {
	var cp miniProgram.MomentsCopywriting
	_ = c.ShouldBindJSON(&cp)
	if err := momentsService.DeleteCopywriting(cp); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (api *MomentsApi) UpdateCopywriting(c *gin.Context) {
	var cp miniProgram.MomentsCopywriting
	_ = c.ShouldBindJSON(&cp)
	if err := momentsService.UpdateCopywriting(cp); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (api *MomentsApi) GetCopywritingList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	categoryIdStr := c.Query("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)
	if list, total, err := momentsService.GetCopywritingList(pageInfo, uint(categoryId)); err != nil {
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

func (api *MomentsApi) PublicMomentsList(c *gin.Context) {
	if list, err := momentsService.GetPublicMomentsList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
