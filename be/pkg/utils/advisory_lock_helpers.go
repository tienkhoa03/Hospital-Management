package utils

import "gorm.io/gorm"

const (
	NamespaceStaffSchedule = 1
)

func AdvisoryLock(tx *gorm.DB, namespace int64, staffID int64) error {
	key := (namespace << 32) | staffID
	return tx.Raw("SELECT pg_advisory_xact_lock(?)", key).Error
}
