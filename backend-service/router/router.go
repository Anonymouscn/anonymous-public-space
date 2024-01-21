package router

import (
	managerSideHandler "anonymous.net.cn/public-system-backend/module/manager/handler"
	"anonymous.net.cn/public-system-backend/module/public/handler"
	"github.com/gin-gonic/gin"
)

// CreateRouter 创建路由
func CreateRouter() *gin.Engine {
	router := gin.Default()
	generateMapping(router)
	return router
}

// generateMapping 生成路由映射
func generateMapping(router *gin.Engine) {
	api := router.Group("api")
	{
		// 第一版 API
		v1 := api.Group("v1")
		{
			// 登录数据接口
			login := v1.Group("login")
			{
				login.POST("/sms", managerSideHandler.AdminSmsLogin)
				login.POST("/password", managerSideHandler.AdminPasswordLogin)
				login.POST("/open-gateway", managerSideHandler.AdminOpenGatewayLogin)
				login.DELETE("/logout", managerSideHandler.AdminLogout)
			}
			// 登出数据接口
			logout := v1.Group("logout")
			{
				logout.DELETE("/", managerSideHandler.AdminLogout)
			}
			// 公开数据接口
			public := v1.Group("public")
			{
				// icp 备案信息查询
				icp := public.Group("icp")
				{
					// 获取 icp 备案展示信息
					icp.GET("/info", handler.GetDisplayIcpLicence)
				}
				// 网页版权信息查询
				copyright := public.Group("copyright")
				{
					// 获取版权展示信息
					copyright.GET("/info", handler.GetDisplayCopyrightLicence)
				}
				// 技术栈信息查询
				ts := public.Group("tech-stack")
				{
					// 获取技术栈展示列表信息
					ts.GET("/list/:limit", handler.GetDisplayTechStackInfo)
				}
				// 联系方式信息查询
				contact := public.Group("contact-me")
				{
					// 获取联系方式展示信息
					contact.GET("/list/:limit", handler.GetDisplayContactInfo)
				}
				// 代码空间信息查询
				cp := public.Group("code-place")
				{
					// 获取代码空间展示信息
					cp.GET("/list/:limit", handler.GetDisplayCodePlaceInfo)
				}
				// 站点入口信息查询
				si := public.Group("site-entrance")
				{
					// 获取站点入口展示信息
					si.GET("/list/:limit", handler.GetSiteEntranceInfo)
				}
				// 用户行为采集
				userAction := public.Group("user-action")
				{
					userAction.POST("/collect", managerSideHandler.CollectUserAction)
				}
			}
			// 管理数据接口
			manager := v1.Group("manager")
			{
				// icp 备案信息管理接口
				icp := manager.Group("icp")
				{
					icp.POST("/edit", managerSideHandler.SetIcpLicence)
					icp.DELETE("/delete", managerSideHandler.DeleteIcpLicence)
				}
				// 版权信息管理接口
				copyright := manager.Group("copyright")
				{
					copyright.POST("/edit", managerSideHandler.SetCopyrightLicence)
					copyright.DELETE("/delete", managerSideHandler.DeleteCopyrightLicence)
				}
				// 技术栈信息管理接口
				techStack := manager.Group("tech-stack")
				{
					techStack.POST("/add", managerSideHandler.AddTechStack)
					techStack.DELETE("/delete/:id-list", managerSideHandler.DeleteTechStack)
					techStack.PUT("/modified", managerSideHandler.ModifiedTechStack)
					techStack.GET("/query/:no/:size", managerSideHandler.QueryTechStack)
				}
				// 联系方式信息管理接口
				contactMe := manager.Group("contact-me")
				{
					contactMe.POST("/add", managerSideHandler.AddContact)
					contactMe.DELETE("/delete/:id-list", managerSideHandler.DeleteContact)
					contactMe.PUT("/modified", managerSideHandler.ModifiedContact)
					contactMe.GET("/query/:no/:size", managerSideHandler.QueryContact)
				}
				// 代码空间信息管理接口
				codePlace := manager.Group("code-place")
				{
					codePlace.POST("/add", managerSideHandler.AddCodePlace)
					codePlace.DELETE("/delete/:id-list", managerSideHandler.DeleteCodePlace)
					codePlace.PUT("/modified", managerSideHandler.ModifiedCodePlace)
					codePlace.GET("/query/:no/:size", managerSideHandler.QueryCodePlace)
				}
				// 站点入口信息管理接口
				siteEntrance := manager.Group("site-entrance")
				{
					siteEntrance.POST("/add", managerSideHandler.AddSiteEntrance)
					siteEntrance.DELETE("/delete/:id-list", managerSideHandler.DeleteSiteEntrance)
					siteEntrance.PUT("/modified", managerSideHandler.ModifiedSiteEntrance)
					siteEntrance.GET("/query/:no/:size", managerSideHandler.QuerySiteEntrance)
				}
				// 管理员信息管理接口
				mine := manager.Group("mine")
				{
					mine.POST("/password", managerSideHandler.AdminResetPassword)
				}
				// 日志信息管理接口
				log := manager.Group("log")
				{
					log.GET("/query", managerSideHandler.QueryLog)
					log.GET("/download", managerSideHandler.DownloadLog)
				}
				// 用户行为信息管理接口
				userAction := manager.Group("user-action")
				{
					userAction.GET("/query", managerSideHandler.QueryUserAction)
					userAction.GET("/summary", managerSideHandler.SummaryUserAction)
				}
			}
		}
	}
}
