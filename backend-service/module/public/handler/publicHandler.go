package handler

import (
	commonModel "anonymous.net.cn/public-system-backend/module/common/model"
	publicRepo "anonymous.net.cn/public-system-backend/module/public/repo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetDisplayIcpLicence
// @Tags 公共信息获取接口
// @Summary 获取 icp 备案展示信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/public/icp/info [get]
func GetDisplayIcpLicence(ctx *gin.Context) {
	licence, err := publicRepo.GetDisplayIcpLicence()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(500, "查询失败"))
	} else {
		ctx.JSON(http.StatusOK, commonModel.SuccessData(licence))
	}
}

// GetDisplayCopyrightLicence
// @Tags 公共信息获取接口
// @Summary 获取版权展示信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/public/copyright/info [get]
func GetDisplayCopyrightLicence(ctx *gin.Context) {
	licence, err := publicRepo.GetDisplayCopyrightLicence()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(500, "查询失败"))
	} else {
		ctx.JSON(http.StatusOK, commonModel.SuccessData(licence))
	}
}

// GetDisplayTechStackInfo
// @Tags 公共信息获取接口
// @Summary 获取技术栈展示信息
// @Produce json
// @Param limit path int true "Top [limit] 条记录"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/public/tech-stack/list/{limit} [get]
func GetDisplayTechStackInfo(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.WrongParam, "参数错误"))
		return
	}
	info, err := publicRepo.GetDisplayTechStack(int64(limit))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.BusinessQueryError, "查询失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(info))
}

// GetDisplayContactInfo
// @Tags 公共信息获取接口
// @Summary 获取联系方式展示信息
// @Produce json
// @Param limit path int true "Top [limit] 条记录"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/public/contact-me/list/{limit} [get]
func GetDisplayContactInfo(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.WrongParam, "参数错误"))
		return
	}
	info, err := publicRepo.GetDisplayContactMe(int64(limit))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.BusinessQueryError, "查询失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(info))
}

// GetDisplayCodePlaceInfo
// @Tags 公共信息获取接口
// @Summary 获取代码空间展示信息
// @Produce json
// @Param limit path int true "Top [limit] 条记录"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/public/code-place/list/{limit} [get]
func GetDisplayCodePlaceInfo(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.WrongParam, "参数错误"))
		return
	}
	info, err := publicRepo.GetDisplayCodePlace(int64(limit))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.BusinessQueryError, "查询失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(info))
}

// GetSiteEntranceInfo
// @Tags 公共信息获取接口
// @Summary 获取站点入口展示信息
// @Produce json
// @Param limit path int true "Top [limit] 条记录"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/public/site-entrance/list/{limit} [get]
func GetSiteEntranceInfo(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.WrongParam, "参数错误"))
		return
	}
	info, err := publicRepo.GetDisplaySiteEntrance(int64(limit))
	if err != nil {
		ctx.JSON(commonModel.BusinessError, commonModel.RespText(commonModel.BusinessQueryError, "查询失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(info))
}
