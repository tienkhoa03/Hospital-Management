package config

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var users = []entity.User{
	{
		Name:     "admin1",
		Email:    "admin@gmail.com",
		Password: "$2a$10$uD2Sp/ceVMQs.Fxa9883Lejcy4QSiEsWFIihuosOkCqwQaCrs011.",
		RoleId:   int64(1),
	},
}

var userRoles = []entity.UserRole{
	{
		RoleSlug: "admin",
	},
	{
		RoleSlug: "manager",
	},
	{
		RoleSlug: "staff",
	},
	{
		RoleSlug: "patient",
	},
}

var staffRoles = []entity.StaffRole{
	{
		RoleSlug: "doctor",
	},
	{
		RoleSlug: "nurse",
	},
	{
		RoleSlug: "cashing_officer",
	},
}

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DB_DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error:", err)
	}

	createUserRoleEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role_slug') THEN
			CREATE TYPE user_role_slug AS ENUM (
				'manager', 
				'staff',
				'patient'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createUserRoleEnumSQL)

	createStaffRoleEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'staff_role_slug') THEN
			CREATE TYPE staff_role_slug AS ENUM (
				'doctor', 
				'nurse',
				'cashing_officer'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createStaffRoleEnumSQL)

	createGenderEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_slug') THEN
			CREATE TYPE gender_slug AS ENUM (
				'male', 
				'female',
				'unknown'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createGenderEnumSQL)

	createAppointmentEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'appointment_status_slug') THEN
			CREATE TYPE appointment_status_slug AS ENUM (
				'scheduled', 
				'completed',
				'canceled'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createAppointmentEnumSQL)
	createTaskEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status_slug') THEN
			CREATE TYPE task_status_slug AS ENUM (
				'scheduled', 
				'completed',
				'canceled'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createTaskEnumSQL)
	createBillEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bill_status_slug') THEN
			CREATE TYPE bill_status_slug AS ENUM (
				'paid', 
				'unpaid'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createBillEnumSQL)

	createPatientEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'patient_status_slug') THEN
			CREATE TYPE patient_status_slug AS ENUM (
				'in_treatment', 
				'never_treated',
				'treated_before',
				'inactive'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createPatientEnumSQL)

	createStaffEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'staff_status_slug') THEN
			CREATE TYPE staff_status_slug AS ENUM (
				'working', 
				'on_leave',
				'inactive'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createStaffEnumSQL)

	createBloodEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'blood_type_slug') THEN
			CREATE TYPE blood_type_slug AS ENUM (
				'A+',  'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createBloodEnumSQL)

	err = db.AutoMigrate(&entity.Appointment{}, &entity.Bill{}, &entity.BillItem{}, &entity.Doctor{}, &entity.Manager{}, &entity.MedicalRecord{}, &entity.Medicine{}, &entity.Nurse{}, &entity.Patient{}, &entity.Prescription{}, &entity.Staff{}, &entity.StaffRole{}, &entity.Task{}, &entity.UserToken{}, &entity.User{}, &entity.UserRole{})
	if err != nil {
		log.Fatal("Error migrate to database. Error:", err)
	}
	for _, user := range users {
		var existing entity.User
		db.Model(&entity.User{}).Where("email = ?", user.Email).FirstOrCreate(&existing, user)
	}
	for _, userRole := range userRoles {
		var existing entity.UserRole
		db.Model(&entity.UserRole{}).Where("role_slug = ?", userRole.RoleSlug).FirstOrCreate(&existing, userRole)
	}
	for _, staffRole := range staffRoles {
		var existing entity.StaffRole
		db.Model(&entity.StaffRole{}).Where("role_slug = ?", staffRole.RoleSlug).FirstOrCreate(&existing, staffRole)
	}
	return db
}
