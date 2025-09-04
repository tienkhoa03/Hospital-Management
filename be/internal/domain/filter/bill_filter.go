package filter

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BillFilter struct {
	CreatedAfter  *time.Time `form:"createdAfter"`
	CreatedBefore *time.Time `form:"createdBefore"`
	Status        *string    `form:"status" binding:"omitempty,oneof=paid unpaid"`
	SortBy        string     `form:"sortBy,default=createdAt" binding:"omitempty,oneof=createdAt"`
	Order         string     `form:"order,default=asc" binding:"omitempty,oneof=asc desc"`
	Page          int        `form:"page,default=1"`
	Limit         int        `form:"limit,default=10"`
}

func (f *BillFilter) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.CreatedAfter != nil {
		db = db.Where("created_at >= ?", f.CreatedAfter)
	}
	if f.CreatedBefore != nil {
		db = db.Where("created_at <= ?", f.CreatedBefore)
	}
	if f.Status != nil {
		db = db.Where("status = ?", *f.Status)
	}
	db = db.Order(fmt.Sprintf("%s %s", f.SortBy, f.Order))
	offset := (f.Page - 1) * f.Limit
	db = db.Offset(offset).Limit(f.Limit)
	return db
}
