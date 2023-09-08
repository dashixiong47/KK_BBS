package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// 分页
type Paging struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

// GetPaging 接收gin的上下文然后获取分页信息
func (p *Paging) GetPaging(ctx *gin.Context) *Paging {
	p.Page = 1
	p.PageSize = 10
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pageSize")
	if pageStr != "" {
		// 转换成int
		p.Page, _ = strconv.Atoi(pageStr)
	}
	if pageSizeStr != "" {
		p.PageSize, _ = strconv.Atoi(pageSizeStr)
	}
	return p
}

// SetPaging 设置分页信息
func (p *Paging) SetPaging(db *gorm.DB) *gorm.DB {
	return db.Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize)
}
