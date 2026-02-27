package repository

import (
	"golang-clean-crud/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	Update(user *entity.User) error
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

// FindByID implements [UserRepository].
func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update implements [UserRepository].
func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}
