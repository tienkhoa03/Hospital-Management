package filter

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AppointmentFilter struct {
	TimeFrom *time.Time `form:"timeFrom"`
	TimeTo   *time.Time `form:"timeTo"`
	Status   *string    `form:"status" binding:"omitempty,oneof=scheduled completed canceled"`
	SortBy   string     `form:"sortBy,default=beginTime" binding:"omitempty,oneof=beginTime createdAt"`
	Order    string     `form:"order,default=asc" binding:"omitempty,oneof=asc desc"`
	Page     int        `form:"page,default=1"`
	Limit    int        `form:"limit,default=10"`
}

func (f *AppointmentFilter) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.TimeFrom != nil {
		db = db.Where("begin_time >= ?", f.TimeFrom)
	}
	if f.TimeTo != nil {
		db = db.Where("finish_time <= ?", f.TimeTo)
	}
	if f.Status != nil {
		db = db.Where("status = ?", *f.Status)
	}
	db = db.Order(fmt.Sprintf("%s %s", f.SortBy, f.Order))
	offset := (f.Page - 1) * f.Limit
	db = db.Offset(offset).Limit(f.Limit)
	return db
}
