package repository

import (
	"github.com/jinzhu/gorm"
	user "user-center/proto/user"
)

type AuthRepository interface {
	Create(user *user.User) error
}

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) Create(user *user.User) error {
	return repo.Db.Create(user).Error
}
