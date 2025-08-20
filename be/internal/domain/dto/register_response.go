package dto

import "time"

type RegisterResponse struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CitizenId   string    `json:"citizen_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender"`
	RoleId      int64     `json:"role_id"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	PatientInfo *PatientInfoResponse `json:"patient_info,omitempty"`
	StaffInfo   *StaffInfoResponse   `json:"staff_info,omitempty"`
	ManagerInfo *ManagerInfoResponse `json:"manager_info,omitempty"`
}

type PatientInfoResponse struct {
	Id                int64   `json:"id"`
	InsuranceNumber   *string `json:"insurance_number"`
	BloodType         *string `json:"blood_type"`
	Allergies         *string `json:"allergies"`
	ChronicConditions *string `json:"chronic_conditions"`
	Status            string  `json:"status"`
}
type StaffInfoResponse struct {
	Id         int64               `json:"id"`
	Department string              `json:"department"`
	Status     string              `json:"status"`
	RoleId     int64               `json:"role_id"`
	Role       string              `json:"role"`
	DoctorInfo *DoctorInfoResponse `json:"doctor_info,omitempty"`
	NurseInfo  *NurseInfoResponse  `json:"nurse_info,omitempty"`
}

type ManagerInfoResponse struct {
	Id         int64  `json:"id"`
	Department string `json:"department"`
	Status     string `json:"status"`
}

type NurseInfoResponse struct {
	Id                   int64  `json:"id"`
	NursingLicenseNumber string `json:"nursing_license_number"`
}

type DoctorInfoResponse struct {
	Id                   int64  `json:"id"`
	Specialization       string `json:"specialization"`
	MedicalLicenseNumber string `json:"medical_license_number"`
}
