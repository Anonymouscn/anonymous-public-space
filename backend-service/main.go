package main

// @title Public System 后端接口
// @version 1.0
// @description Anonymous Public Space 公共指引服务后端接口
// @termsOfService https://anonymous.net.cn
// @contact.name anonymous
// @contact.url https://anonymous.net.cn
// @contact.email pgl888999@icloud.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 192.168.2.45:8085
// @BasePath /api

import (
	"anonymous.net.cn/public-system-backend/config"
	"anonymous.net.cn/public-system-backend/cors"
	"anonymous.net.cn/public-system-backend/db"
	_ "anonymous.net.cn/public-system-backend/docs"
	"anonymous.net.cn/public-system-backend/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
)

// 服务初始化
func init() {
	config.Init() // 初始化配置
	db.Init()     // 初始化数据库
}

// 服务运行
func main() {
	r := router.CreateRouter()
	r.Use(cors.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":" + strconv.Itoa(config.Config.Server.Port))
	if err != nil {
		return
	}
}
