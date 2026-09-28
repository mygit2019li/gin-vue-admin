import service from '@/utils/request'

// @Tags Banner
// @Summary 创建Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body miniProgram.Banner true "创建Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /miniProgram/banner/createBanner [post]
export const createBanner = (data) => {
    return service({
        url: '/banner/createBanner',
        method: 'post',
        data
    })
}

// @Tags Banner
// @Summary 删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body miniProgram.Banner true "删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniProgram/banner/deleteBanner [delete]
export const deleteBanner = (data) => {
    return service({
        url: '/banner/deleteBanner',
        method: 'delete',
        data
    })
}

// @Tags Banner
// @Summary 更新Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body miniProgram.Banner true "更新Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniProgram/banner/updateBanner [put]
export const updateBanner = (data) => {
    return service({
        url: '/banner/updateBanner',
        method: 'put',
        data
    })
}

// @Tags Banner
// @Summary 分页获取Banner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Banner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniProgram/banner/getBannerList [get]
export const getBannerList = (params) => {
    return service({
        url: '/banner/getBannerList',
        method: 'get',
        params
    })
}

// Moments Categories
export const createCategory = (data) => {
    return service({
        url: '/moments/createCategory',
        method: 'post',
        data
    })
}

export const deleteCategory = (data) => {
    return service({
        url: '/moments/deleteCategory',
        method: 'delete',
        data
    })
}

export const updateCategory = (data) => {
    return service({
        url: '/moments/updateCategory',
        method: 'put',
        data
    })
}

export const getCategoryList = () => {
    return service({
        url: '/moments/getCategoryList',
        method: 'get'
    })
}

// Moments Copywriting
export const createCopywriting = (data) => {
    return service({
        url: '/moments/createCopywriting',
        method: 'post',
        data
    })
}

export const deleteCopywriting = (data) => {
    return service({
        url: '/moments/deleteCopywriting',
        method: 'delete',
        data
    })
}

export const updateCopywriting = (data) => {
    return service({
        url: '/moments/updateCopywriting',
        method: 'put',
        data
    })
}

export const getCopywritingList = (params) => {
    return service({
        url: '/moments/getCopywritingList',
        method: 'get',
        params
    })
}
