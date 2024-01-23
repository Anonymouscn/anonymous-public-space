package handler

import (
	commonModel "anonymous.net.cn/public-system-backend/module/common/model"
	"anonymous.net.cn/public-system-backend/module/manager/repo"
	publicModel "anonymous.net.cn/public-system-backend/module/public/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
)

// ========== 系统信息管理 ========== //

// ========== 用户行为管理 ========== //

// CollectUserAction 采集用户行为
func CollectUserAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// QueryUserAction 查询用户行为
func QueryUserAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// SummaryUserAction 分析汇总用户行为
func SummaryUserAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ========== 日志信息管理 ========== //

// QueryLog 日志查询接口
func QueryLog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// DownloadLog 下载日志接口
func DownloadLog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ========== 管理员信息管理 ========== //

// AdminSmsLogin 管理员短信登录
func AdminSmsLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// AdminPasswordLogin 管理员密码登录
func AdminPasswordLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// AdminOpenGatewayLogin 管理员开放平台登录 [待开发]
// @Tags  管理员信息管理接口
// @Summary 管理员登出
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/logout [delete]
func AdminOpenGatewayLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// AdminLogout 管理员登出
// @Tags  管理员信息管理接口
// @Summary 管理员登出
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/logout [delete]
func AdminLogout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// AdminResetPassword
// @Tags  管理员信息管理接口
// @Summary 管理员重置密码
// @Produce json
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/mine/password [post]
func AdminResetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "备案信息设置失败"))
		return
	}
	id, err := repo.SetIcpLicence(*icp)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "备案信息设置失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(id))
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
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessDeleteError, "备案信息删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "版权信息设置失败"))
		return
	}
	id, err := repo.SetCopyrightLicence(*copyright)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "版权信息设置失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(id))
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
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessDeleteError, "版权信息删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "技术栈创建失败"))
		return
	}
	id, err := repo.AddTechStack(*techStack)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "技术栈创建失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(id))
}

// DeleteTechStack +
// @Tags 技术栈信息管理接口
// @Summary 删除技术栈信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/tech-stack/delete/{id-list} [delete]
func DeleteTechStack(ctx *gin.Context) {
	idStrList := strings.Split(ctx.Param("id-list"), ",")
	idList := make([]primitive.ObjectID, 0, len(idStrList))
	for _, v := range idStrList {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "技术栈删除失败"))
			return
		}
		idList = append(idList, id)
	}
	res, err := repo.DeleteTechStack(idList)
	if err != nil || !res {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "技术栈删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ModifiedTechStack +
// @Tags 技术栈信息管理接口
// @Summary 修改技术栈信息
// @Produce json
// @Param copyright body publicModel.TechStack false "技术栈对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/tech-stack/modified [put]
func ModifiedTechStack(ctx *gin.Context) {
	var ts publicModel.TechStack
	err := ctx.BindJSON(&ts)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "技术栈修改失败"))
		return
	}
	_, err = repo.ModifiedTechStack(ts)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "技术栈修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ========== 联系方式信息管理 ========== //

// AddContact +
// @Tags 联系方式信息管理接口
// @Summary 添加联系方式信息
// @Produce json
// @Param copyright body publicModel.Contact false "联系方式对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/add [post]
func AddContact(ctx *gin.Context) {
	contact := &publicModel.Contact{}
	err := ctx.BindJSON(contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "联系方式创建失败"))
		return
	}
	id, err := repo.AddContactMe(*contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "联系方式创建失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(id))
}

// DeleteContact +
// @Tags 联系方式信息管理接口
// @Summary 删除联系方式信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/delete/{id-list} [delete]
func DeleteContact(ctx *gin.Context) {
	idStrList := strings.Split(ctx.Param("id-list"), ",")
	idList := make([]primitive.ObjectID, 0, len(idStrList))
	for _, v := range idStrList {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "联系方式删除失败"))
			return
		}
		idList = append(idList, id)
	}
	res, err := repo.DeleteContactMe(idList)
	if err != nil || !res {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "联系方式删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ModifiedContact +
// @Tags 联系方式信息管理接口
// @Summary 修改联系方式信息
// @Produce json
// @Param copyright body publicModel.Contact false "联系方式对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/contact-me/modified [put]
func ModifiedContact(ctx *gin.Context) {
	var cm publicModel.Contact
	err := ctx.BindJSON(&cm)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "联系方式修改失败"))
		return
	}
	_, err = repo.ModifiedContactMe(cm)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "联系方式修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ========== 代码空间信息管理 ========== //

// AddCodePlace +
// @Tags 代码空间信息管理接口
// @Summary 添加代码空间信息
// @Produce json
// @Param copyright body publicModel.Contact false "代码空间对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/add [post]
func AddCodePlace(ctx *gin.Context) {
	codePlace := &publicModel.Contact{}
	err := ctx.BindJSON(codePlace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "代码空间创建失败"))
		return
	}
	id, err := repo.AddCodePlace(*codePlace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "代码空间创建失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(id))
}

// DeleteCodePlace +
// @Tags 代码空间信息管理接口
// @Summary 删除代码空间信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/delete/{id-list} [delete]
func DeleteCodePlace(ctx *gin.Context) {
	idStrList := strings.Split(ctx.Param("id-list"), ",")
	idList := make([]primitive.ObjectID, 0, len(idStrList))
	for _, v := range idStrList {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "代码空间删除失败"))
			return
		}
		idList = append(idList, id)
	}
	res, err := repo.DeleteCodePlace(idList)
	if err != nil || !res {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "代码空间删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ModifiedCodePlace +
// @Tags 代码空间信息管理接口
// @Summary 修改代码空间信息
// @Produce json
// @Param copyright body publicModel.Contact false "代码空间对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/code-place/modified [put]
func ModifiedCodePlace(ctx *gin.Context) {
	var cp publicModel.Contact
	err := ctx.BindJSON(&cp)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "代码空间修改失败"))
		return
	}
	_, err = repo.ModifiedCodePlace(cp)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "代码空间修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ========== 站点入口信息管理 ========== //

// AddSiteEntrance +
// @Tags 站点入口信息管理接口
// @Summary 添加站点入口信息
// @Produce json
// @Param copyright body publicModel.SiteEntrance false "站点入口对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/add [post]
func AddSiteEntrance(ctx *gin.Context) {
	siteEntrance := &publicModel.SiteEntrance{}
	err := ctx.BindJSON(siteEntrance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "站点入口创建失败"))
		return
	}
	id, err := repo.AddSiteEntrance(*siteEntrance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "站点入口创建失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessData(id))
}

// DeleteSiteEntrance +
// @Tags 站点入口信息管理接口
// @Summary 删除站点信息
// @Produce json
// @Param id-list path []string true "待删除项id列表"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/delete/{id-list} [delete]
func DeleteSiteEntrance(ctx *gin.Context) {
	idStrList := strings.Split(ctx.Param("id-list"), ",")
	idList := make([]primitive.ObjectID, 0, len(idStrList))
	for _, v := range idStrList {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "站点入口删除失败"))
			return
		}
		idList = append(idList, id)
	}
	res, err := repo.DeleteSiteEntrance(idList)
	if err != nil || !res {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "站点入口删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ModifiedSiteEntrance +
// @Tags 站点入口信息管理接口
// @Summary 修改站点信息
// @Produce json
// @Param copyright body publicModel.SiteEntrance false "站点入口对象信息"
// @Success 200 {object} model.Result "请求成功"
// @Failure 500 {object} model.Result "请求失败"
// @Router /v1/manager/site-entrance/modified [put]
func ModifiedSiteEntrance(ctx *gin.Context) {
	var se publicModel.SiteEntrance
	err := ctx.BindJSON(&se)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "站点入口修改失败"))
		return
	}
	_, err = repo.ModifiedSiteEntrance(se)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, commonModel.RespText(commonModel.BusinessInsertError, "站点入口修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
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
	ctx.JSON(http.StatusOK, commonModel.SuccessText())
}

// ========== 网站图标信息管理 ==========//
