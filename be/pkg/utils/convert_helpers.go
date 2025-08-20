package utils

import (
	"BE_Hospital_Management/internal/domain/entity"
)

func ConvertUsersToEmails(users []*entity.User) []string {
	var emails []string
	for _, user := range users {
		if user != nil {
			emails = append(emails, user.Email)
		}
	}
	return emails
}
