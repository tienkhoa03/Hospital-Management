package dto

import "time"

type UpdateUserRequest struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender"`
}

type UpdateManagerRequest struct {
	Department string `json:"department"`
	Status     string `json:"status"`
}
type UpdateDoctorRequest struct {
	Department           string `json:"department"`
	Status               string `json:"status"`
	Specialization       string `json:"specialization"`
	MedicalLicenseNumber string `json:"medical_license_number"`
}

type UpdateNurseRequest struct {
	Department           string `json:"department"`
	Status               string `json:"status"`
	NursingLicenseNumber string `json:"nursing_license_number"`
}

type UpdateCashingOfficerRequest struct {
	Department string `json:"department"`
	Status     string `json:"status"`
}

type UpdatePatientRequest struct {
	InsuranceNumber   string `json:"insurance_number"`
	BloodType         string `json:"blood_type"`
	Allergies         string `json:"allergies"`
	ChronicConditions string `json:"chronic_conditions"`
	Status            string `json:"status"`
}
