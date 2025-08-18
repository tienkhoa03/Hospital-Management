package entity

import "time"

type Medicine struct {
	Id              int64  `gorm:"primaryKey;autoIncrement"`
	Name            string `gorm:"type:varchar(256)"`
	UsesInstruction string `gorm:"type:varchar(256)"`
	Price           int    `json:"price"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
