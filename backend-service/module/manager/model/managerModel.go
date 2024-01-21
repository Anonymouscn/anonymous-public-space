package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// System 系统信息
type System struct {
}

// Admin 管理员信息
type Admin struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`  // id
	Name     string             `json:"name" bson:"name"`         // 姓名
	Password string             `json:"password" bson:"password"` // 密码
}

// Log 日志信息
type Log struct {
}
