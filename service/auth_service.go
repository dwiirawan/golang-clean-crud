package service

import (
	"errors"
	"golang-clean-crud/dto"
	"golang-clean-crud/entity"
	"golang-clean-crud/repository"
	"golang-clean-crud/utils"

	"golang.org/x/crypto/bcrypt"
)

// INTERFACE
type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
	RefreshToken(refreshToken string) (string, error)
	Logout(userID uint) error
}

// IMPLEMENTATION STRUCT
type authService struct {
	repo repository.UserRepository
}

// CONSTRUCTOR (DEPENDENCY INJECTION)
func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{r}
}

// Register implements [AuthService].
func (s *authService) Register(req dto.RegisterRequest) error {
	// email exsisting check
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil && existing.ID != 0 {
		return errors.New("Email already registered")
	}

	// hash password
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
		Role:     "user",
	}

	return s.repo.Create(&user)
}

// Login implements [AuthService].
func (s *authService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return dto.LoginResponse{}, errors.New("Invalid password")
	}

	// Generate token
	access, _ := utils.GenerateAccessToken(user.ID, user.Role)
	refresh, _ := utils.GenerateRefreshToken(user.ID)

	// Save refresh token in DB
	user.RefreshToken = refresh
	err = s.repo.Update(user)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// RefreshToken implements [AuthService].
func (s *authService) RefreshToken(token string) (string, error) {
	claims, err := utils.VerifyRefreshToken(token)
	if err != nil {
		return "", errors.New("Invalid refresh token")
	}

	userID := uint(claims["user_id"].(float64))

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return "", errors.New("User not found")
	}

	// Check token
	if user.RefreshToken != token {
		return "", errors.New("Refresh token revoked")
	}

	access, err := utils.GenerateAccessToken(userID, user.Role)
	if err != nil {
		return "", err
	}

	return access, nil
}

// Logout implements [AuthService].
func (s *authService) Logout(userID uint) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	user.RefreshToken = ""

	return s.repo.Update(user)
}
