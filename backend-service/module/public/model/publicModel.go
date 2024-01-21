package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// IcpLicence 备案信息
type IcpLicence struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`          // 备案信息 id
	Province    string             `json:"province" bson:"province"`         // 备案省份
	Number      string             `json:"number" bson:"number"`             // 备案号
	No          string             `json:"no" bson:"no"`                     // 附加号码
	GmtRegistry string             `json:"gmt-registry" bson:"gmt_registry"` // 登记时间
	GmtCreate   time.Time          `json:"-" bson:"gmt_create"`              // 记录创建时间
	GmtModified time.Time          `json:"-" bson:"gmt_update"`              // 记录最新修改时间
}

// CopyrightLicence 版权信息
type CopyrightLicence struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` // 版权信息 id
	From        string             `json:"from" bson:"from"`        // 生效开始标记
	To          string             `json:"to" bson:"to"`            // 生效结束标记
	Author      string             `json:"author" bson:"author"`    // 版权作者
	GmtCreate   time.Time          `json:"-" bson:"gmt_create"`     // 记录创建时间
	GmtModified time.Time          `json:"-" bson:"gmt_modified"`   // 记录最新修改时间
}

// TechStack 技术栈信息
type TechStack struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` // 技术栈信息 id
	Icon        string             `json:"icon" bson:"icon"`        // 图标 svg 名称
	Name        string             `json:"name" bson:"name"`        // 技术栈名称
	Weight      int                `json:"weight" bson:"weight"`    // 权重
	GmtCreate   time.Time          `json:"-" bson:"gmt_create"`     // 记录创建时间
	GmtModified time.Time          `json:"-" bson:"gmt_modified"`   // 记录最新修改时间
}

// Contact 联系方式信息
type Contact struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` // 联系方式信息 id
	Icon        string             `json:"icon" bson:"icon"`        // 联系方式图标名称
	Name        string             `json:"name" bson:"name"`        // 联系方式名称
	Url         string             `json:"url" bson:"url"`          // 跳转 url (可选)
	Image       string             `json:"image" bson:"image"`      // 二维码图片 url (可选)
	Type        int                `json:"type" bson:"type"`        // 展示类型 [0: url, 1: 二维码]
	Weight      int                `json:"weight" bson:"weight"`    // 权重
	GmtCreate   time.Time          `json:"-" bson:"gmt_create"`     // 记录创建时间
	GmtModified time.Time          `json:"-" bson:"gmt_modified"`   // 记录最新修改时间
}

// SiteEntrance 站点入口信息
type SiteEntrance struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` // 站点信息 id
	Name        string             `json:"name" bson:"name"`        // 站点名称
	Url         string             `json:"url" bson:"url"`          // 站点 url
	Weight      int                `json:"weight" bson:"weight"`    // 权重
	GmtCreate   string             `json:"-" bson:"gmt_create"`     // 记录创建时间
	GmtModified string             `json:"-" bson:"gmt_modified"`   // 记录最新修改时间
}
