package config

import (
	"BE_Friends_Management/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var admins = []entity.Admin{
	{
		Name:     "admin1",
		Email:    "admin@gmail.com",
		Password: "$2a$10$uD2Sp/ceVMQs.Fxa9883Lejcy4QSiEsWFIihuosOkCqwQaCrs011.",
	},
}

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DB_DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error:", err)
	}

	createRoleEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_slug') THEN
			CREATE TYPE role_slug AS ENUM (
				'manager', 
				'doctor',
				'patient',
				'cashing_officer'
			);
		END IF;
	END
	$$;
	`
	db.Exec(createRoleEnumSQL)

	createGenderEnumSQL := `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_slug') THEN
			CREATE TYPE gender_slug AS ENUM (
				'male', 
				'female'
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

	err = db.AutoMigrate(&entity.Admin{}, &entity.Appointment{}, &entity.Bill{}, &entity.BillItem{}, &entity.Doctor{}, &entity.Manager{}, &entity.MedicalRecord{}, &entity.Medicine{}, &entity.Patient{}, &entity.Prescription{}, &entity.Role{}, &entity.Task{}, &entity.UserToken{}, &entity.User{})
	if err != nil {
		log.Fatal("Error migrate to database. Error:", err)
	}
	for _, admin := range admins {
		var existing entity.Admin
		db.Model(&entity.Admin{}).Where("email = ?", admin.Email).FirstOrCreate(&existing, admin)
	}
	return db
}
