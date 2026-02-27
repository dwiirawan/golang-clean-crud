package repository

import (
	"golang-clean-crud/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	Create(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create implements [UserRepository].
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// FindByEmail implements [UserRepository].
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
