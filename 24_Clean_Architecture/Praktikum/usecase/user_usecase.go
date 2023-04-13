package usecase

import (
	"clean_architecture/dto"
	"clean_architecture/model"
	"clean_architecture/repository"
)

type UserUsecase interface {
	Create(payloads []dto.CreateUserRequest) error
	GetAll() ([]model.User, error)

}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) Create(payloads []dto.CreateUserRequest) error {
	var (
		userData model.User
	)
	if err := s.userRepository.Create(userData); err != nil {
		return err
	}
	return nil

}

func (s *userUsecase) GetAll() ([]model.User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

