package repository

import (
	"clean_architecture/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Find(Email string) (model.User, error)
	GetAll() ([]model.User, error)
	Create(data model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Find(Email string) (model.User, error) {
	var user model.User

	if err := r.db.Model(&user).Where("Email = ?", Email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Create(data model.User) error {
	return r.db.Create(&data).Error
}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
