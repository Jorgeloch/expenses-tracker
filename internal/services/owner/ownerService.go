package ownerService

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	ownerDTO "github.com/jorgeloch/expenses-tracker/internal/dto/owner"
	"github.com/jorgeloch/expenses-tracker/internal/models/owner"
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository *repository.Repository
}

func Init(r *repository.Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) GetAll() ([]ownerModel.Owner, error) {
	return s.Repository.OwnerRepository.GetAll()
}

func (s *Service) GetByID(id string) (ownerModel.Owner, error) {
	return s.Repository.OwnerRepository.GetByID(id)
}

func (s *Service) GetByEmail(email string) (ownerModel.Owner, error) {
	return s.Repository.OwnerRepository.GetByEmail(email)
}

func (s *Service) Create(dto ownerDTO.CreateOwnerDTO) (uuid.UUID, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return uuid.Nil, err
	}

	location := uuid.New()

	err = s.Repository.OwnerRepository.Create(ownerModel.Owner{
		ID:       location,
		Email:    dto.Email,
		Name:     dto.Name,
		Password: string(password),
	})

	if err != nil {
		return uuid.Nil, err
	}

	return location, nil
}

func (s *Service) Update(id string, dto ownerDTO.UpdateOwnerDTO) (ownerModel.Owner, error) {
	owner, err := s.GetByID(id)

	if err != nil {
		return ownerModel.Owner{}, err
	}

	if dto.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return ownerModel.Owner{}, err
		}

		owner.Password = string(password)
	}

	if dto.Name != "" {
		owner.Name = dto.Name
	}

	if dto.Email != "" {
		owner.Email = dto.Email
	}

	return owner, s.Repository.OwnerRepository.Update(owner)
}

func (s *Service) Login(dto ownerDTO.LoginDTO) (string, error) {
	owner, err := s.GetByEmail(dto.Email)
	if owner.ID == uuid.Nil {
		return "", errors.New("User not found")
	}
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(dto.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  owner.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) Delete(id string) error {
	return s.Repository.OwnerRepository.Delete(id)
}
