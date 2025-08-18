package entity

import "time"

type MedicalRecord struct {
	Id            int64 `gorm:"primaryKey"`
	PatientId     int64
	DoctorId      int64
	AppointmentId *int64

	Symptoms        string
	MedicalServices string
	Diagnosis       string
	Treatment       string
	Note            string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Patient     *Patient     `gorm:"foreignKey:PatientId;references:Id"`
	Doctor      *Doctor      `gorm:"foreignKey:DoctorId;references:Id"`
	Appointment *Appointment `gorm:"foreignKey:AppointmentId;references:Id"`
}
