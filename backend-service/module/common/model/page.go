package model

// Page 分页信息
type Page[T any] struct {
	Total   uint64 `json:"total"`   // 总记录数
	Pages   uint64 `json:"pages"`   // 总页数
	Size    uint64 `json:"size"`    // 页大小
	No      uint64 `json:"no"`      // 当前页
	Records []T    `json:"records"` // 记录
}

// Display 展示信息
type Display[T any] struct {
	Total   uint64 `json:"total"`   // 总记录数
	Size    uint64 `json:"size"`    // 当前查询范围大小
	Records []T    `json:"records"` // 当前查询记录
}
