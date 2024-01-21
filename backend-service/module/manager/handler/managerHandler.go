package handler

import (
	"anonymous.net.cn/public-system-backend/module/common/model"
	"anonymous.net.cn/public-system-backend/module/manager/repo"
	publicModel "anonymous.net.cn/public-system-backend/module/public/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ========== 系统信息管理 ========== //

// ========== 用户行为管理 ========== //

// CollectUserAction 采集用户行为
func CollectUserAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// QueryUserAction 查询用户行为
func QueryUserAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// SummaryUserAction 分析汇总用户行为
func SummaryUserAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 日志信息管理 ========== //

// QueryLog 日志查询接口
func QueryLog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// DownloadLog 下载日志接口
func DownloadLog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 管理员信息管理 ========== //

// AdminSmsLogin 管理员短信登录
func AdminSmsLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// AdminPasswordLogin 管理员密码登录
func AdminPasswordLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// AdminOpenGatewayLogin 管理员开放平台登录 [待开发]
// @Tags  管理员信息管理接口
// @Summary 管理员登出
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/logout [delete]
func AdminOpenGatewayLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// AdminLogout 管理员登出
// @Tags  管理员信息管理接口
// @Summary 管理员登出
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/logout [delete]
func AdminLogout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// AdminResetPassword
// @Tags  管理员信息管理接口
// @Summary 管理员重置密码
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/mine/password [post]
func AdminResetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 备案信息管理 ========== //

// SetIcpLicence +
// @Tags 备案信息管理接口
// @Summary 设置 icp 备案信息
// @Produce json
// @Param copyright body publicModel.IcpLicence false "新增备案信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/icp/edit [post]
func SetIcpLicence(ctx *gin.Context) {
	icp := &publicModel.IcpLicence{}
	err := ctx.BindJSON(icp)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessInsertError, "备案信息设置失败"))
		return
	}
	id, err := repo.SetIcpLicence(*icp)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessInsertError, "备案信息设置失败"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessData(id))
}

// DeleteIcpLicence +
// @Tags 备案信息管理接口
// @Summary 删除 icp 备案信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/icp/delete [delete]
func DeleteIcpLicence(ctx *gin.Context) {
	_, err := repo.DeleteIcpLicence()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessDeleteError, "备案信息删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 版权信息管理 ========== //

// SetCopyrightLicence +
// @Tags 版权信息管理接口
// @Summary 设置版权信息
// @Produce json
// @Param copyright body publicModel.CopyrightLicence false "新增版权信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/copyright/edit [post]
func SetCopyrightLicence(ctx *gin.Context) {
	copyright := &publicModel.CopyrightLicence{}
	err := ctx.BindJSON(copyright)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessInsertError, "版权信息设置失败"))
		return
	}
	id, err := repo.SetCopyrightLicence(*copyright)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessInsertError, "版权信息设置失败"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessData(id))
}

// DeleteCopyrightLicence +
// @Tags 版权信息管理接口
// @Summary 删除版权信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/copyright/delete [delete]
func DeleteCopyrightLicence(ctx *gin.Context) {
	_, err := repo.DeleteCopyrightLicence()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessDeleteError, "版权信息删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 技术栈信息管理 ========== //

// AddTechStack +
// @Tags 技术栈信息管理接口
// @Summary 添加技术栈信息
// @Produce json
// @Param copyright body publicModel.TechStack false "技术栈对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/tech-stack/add [post]
func AddTechStack(ctx *gin.Context) {
	techStack := &publicModel.TechStack{}
	err := ctx.BindJSON(techStack)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessInsertError, "技术栈创建失败"))
		return
	}
	id, err := repo.AddTechStack(*techStack)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.RespText(model.BusinessInsertError, "技术栈创建失败"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessData(id))
}

// DeleteTechStack
// @Tags 技术栈信息管理接口
// @Summary 移除技术栈信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/tech-stack/delete/{id-list} [delete]
func DeleteTechStack(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ModifiedTechStack
// @Tags 技术栈信息管理接口
// @Summary 修改技术栈信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/tech-stack/modified [put]
func ModifiedTechStack(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// QueryTechStack
// @Tags 技术栈信息管理接口
// @Summary 分页查询技术栈信息
// @Produce json
// @Param no path int true "分页页码"
// @Param size path int true "分页大小"
// @Param keyword query string false "搜索关键字"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/tech-stack/query/{no}/{size} [get]
func QueryTechStack(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 联系方式信息管理 ========== //

// AddContact
// @Tags 联系方式信息管理接口
// @Summary 添加联系方式信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/add [post]
func AddContact(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// DeleteContact
// @Tags 联系方式信息管理接口
// @Summary 删除联系方式信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/delete/{id-list} [delete]
func DeleteContact(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ModifiedContact
// @Tags 联系方式信息管理接口
// @Summary 修改联系方式信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/modified [put]
func ModifiedContact(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// QueryContact
// @Tags 联系方式信息管理接口
// @Summary 分页查询联系方式信息
// @Produce json
// @Param no path int true "分页页码"
// @Param size path int true "分页大小"
// @Param keyword query string false "搜索关键字"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/query/{no}/{size} [get]
func QueryContact(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 代码空间信息管理 ========== //

// AddCodePlace
// @Tags 代码空间信息管理接口
// @Summary 添加代码空间信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/add [post]
func AddCodePlace(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// DeleteCodePlace
// @Tags 代码空间信息管理接口
// @Summary 删除代码空间信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/delete/{id-list} [delete]
func DeleteCodePlace(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ModifiedCodePlace
// @Tags 代码空间信息管理接口
// @Summary 修改代码空间信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/modified [put]
func ModifiedCodePlace(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// QueryCodePlace
// @Tags 代码空间信息管理接口
// @Summary 查询代码空间信息
// @Produce json
// @Param no path int true "分页页码"
// @Param size path int true "分页大小"
// @Param keyword query string false "搜索关键字"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/query/{no}/{size} [get]
func QueryCodePlace(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ========== 站点入口信息管理 ========== //

// AddSiteEntrance
// @Tags 站点入口信息管理接口
// @Summary 添加站点入口信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/add [post]
func AddSiteEntrance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// DeleteSiteEntrance
// @Tags 站点入口信息管理接口
// @Summary 删除站点信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/delete/{id-list} [delete]
func DeleteSiteEntrance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// ModifiedSiteEntrance
// @Tags 站点入口信息管理接口
// @Summary 修改站点信息
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/modified [put]
func ModifiedSiteEntrance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}

// QuerySiteEntrance
// @Tags 站点入口信息管理接口
// @Summary 查询站点信息
// @Produce json
// @Param no path int true "分页页码"
// @Param size path int true "分页大小"
// @Param keyword query string false "搜索关键字"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/query/{no}/{size} [get]
func QuerySiteEntrance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessText())
}
