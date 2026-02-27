package service

import (
	"errors"
	"golang-clean-crud/dto"
	"golang-clean-crud/entity"
	"golang-clean-crud/repository"
	"golang-clean-crud/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{r}
}

// Login implements [AuthService].
func (s *authService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("Invalid password")
	}

	token, err := utils.GenerateJWT(user.ID)

	return token, err
}

// Register implements [AuthService].
func (s *authService) Register(req dto.RegisterRequest) error {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repo.Create(&user)
}
