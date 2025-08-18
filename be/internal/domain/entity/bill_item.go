package entity

import "time"

type BillItem struct {
	Id         int64 `gorm:"primaryKey"`
	BillId     int64
	MedicineId int64
	Amount     int
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Bill     *Bill     `gorm:"foreignKey:BillId;references:Id"`
	Medicine *Medicine `gorm:"foreignKey:MedicineId;references:Id"`
}
