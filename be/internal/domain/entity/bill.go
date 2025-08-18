package entity

import "time"

type Bill struct {
	Id               int64 `gorm:"primaryKey;autoIncrement"`
	PatientId        int64
	DoctorId         int64
	CashingOfficerId *int64
	MedicalRecordId  int64
	Description      *string
	TotalPrice       int
	Status           string `gorm:"type:bill_status_slug"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	CashingOfficer *User          `gorm:"foreignKey:CashingOfficerId;references:Id"`
	MedicalRecord  *MedicalRecord `gorm:"foreignKey:MedicalRecordId;references:Id"`
	Patient        *Patient       `gorm:"foreignKey:PatientId;references:Id"`
	Doctor         *Doctor        `gorm:"foreignKey:DoctorId;references:Id"`
}
