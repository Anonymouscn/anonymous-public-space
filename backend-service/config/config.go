package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config 服务配置
var Config Service

// Service 后端服务配置
type Service struct {
	Server Server `yaml:"server"` // 服务器配置
	DB     DB     `yaml:"db"`     // 数据库配置
	Cache  Cache  `yaml:"cache"`  // 缓存配置
}

// Server 服务器配置
type Server struct {
	Port int `yaml:"port"` // 服务器端口
}

// DB 数据库配置
type DB struct {
	Mongo Mongo `yaml:"mongo"` // mongo 数据库配置
}

// Mongo 数据库配置
type Mongo struct {
	DB       string `yaml:"db"`       // 数据库名
	Host     string `yaml:"host"`     // 主机
	Port     int    `yaml:"port"`     // 端口
	Username string `yaml:"username"` // 用户名
	Password string `yaml:"password"` // 密码
	Pool     Pool   `yaml:"pool"`     // 客户端连接池
}

// Pool 客户端链接池
type Pool struct {
	Min int `yaml:"min"` // 最小可用客户端
	Max int `yaml:"max"` // 最大可用客户端
}

// Cache 缓存配置
type Cache struct {
}

func Init() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln("配置文件读取失败")
	}
	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		log.Fatalln("配置文件初始化失败")
	}
	if Config.Server.Port == 0 {
		Config.Server.Port = 8080
	}
}
