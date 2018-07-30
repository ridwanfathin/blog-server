package user

import (
	"api/postgres"

	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB, user *User) (uint, error) {
	err := db.Create(user).Error
	if err != nil {
		if postgres.IsUniqueConstraintError(err, UniqueConstraintUsername) {
			return 0, &UsernameDuplicateError{Username: user.Username}
		}
		if postgres.IsUniqueConstraintError(err, UniqueConstraintEmail) {
			return 0, &EmailDuplicateError{Email: user.Email}
		}
		return 0, err
	}
	return user.ID, nil
}
