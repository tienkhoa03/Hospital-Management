package dto

type UserInfoRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CitizenId   string `json:"citizen_id" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Role        string `json:"role" binding:"required"`

	PatientInfo *PatientInfoRequest `json:"patient_info,omitempty"`
	StaffInfo   *StaffInfoRequest   `json:"staff_info,omitempty"`
	ManagerInfo *ManagerInfoRequest `json:"manager_info,omitempty"`
}

type PatientInfoRequest struct {
	InsuranceNumber   *string `json:"insurance_number"`
	BloodType         *string `json:"blood_type"`
	Allergies         *string `json:"allergies"`
	ChronicConditions *string `json:"chronic_conditions"`
	Status            string  `json:"status"`
}
type StaffInfoRequest struct {
	Department string             `json:"department" binding:"required"`
	Status     string             `json:"status"`
	Role       string             `json:"role" binding:"required"`
	DoctorInfo *DoctorInfoRequest `json:"doctor_info,omitempty"`
	NurseInfo  *NurseInfoRequest  `json:"nurse_info,omitempty"`
}

type ManagerInfoRequest struct {
	Department string `json:"department" binding:"required"`
	Status     string `json:"status"`
}

type NurseInfoRequest struct {
	NursingLicenseNumber string `json:"nursing_license_number" binding:"required"`
}

type DoctorInfoRequest struct {
	Specialization       string `json:"specialization" binding:"required"`
	MedicalLicenseNumber string `json:"medical_license_number" binding:"required"`
}
