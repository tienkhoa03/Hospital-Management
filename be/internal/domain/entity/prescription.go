package entity

import "time"

type Prescription struct {
	Id              int64 `gorm:"primaryKey;autoIncrement"`
	MedicalRecordId int64
	MedicineId      int64
	Amount          int

	MedicalRecord *MedicalRecord `gorm:"foreignKey:MedicalRecordId;references:Id"`
	Medicine      *Medicine      `gorm:"foreignKey:MedicineId;references:Id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
