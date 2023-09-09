package user

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidUsername error = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
var ErrConfirmPasswordNotMatched error = errors.New("Confirm Password field should be equal to Password")

type Service interface {
	Create(payload CreateUserDTO) (UserEntity, error)
	GetOneByID(id int) (UserEntity, error)
	GetOneByUsername(username string) (UserEntity, error)
	UpdateOneByID(id int, payload UpdateUserDTO) (UserEntity, error)
	DeleteOneByID(id int) (UserEntity, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Create(payload CreateUserDTO) (UserEntity, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return UserEntity{}, ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return UserEntity{}, ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserEntity{}, err
	}
	user, err := s.repo.CreateUser(payload.Username, string(hashedPassword))
	if err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (s *service) GetOneByID(id int) (UserEntity, error) {
	user, err := s.repo.GetOneUserByID(id)
	if err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (s *service) GetOneByUsername(username string) (UserEntity, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(username) {
		return UserEntity{}, ErrInvalidUsername
	}
	user, err := s.repo.GetOneUserByUsername(username)
	if err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (s *service) UpdateOneByID(id int, payload UpdateUserDTO) (UserEntity, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return UserEntity{}, ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return UserEntity{}, ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserEntity{}, err
	}
	user, err := s.repo.UpdateUserByID(id, payload.Username, string(hashedPassword))
	if err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (s *service) DeleteOneByID(id int) (UserEntity, error) {
	user, err := s.repo.DeleteUserByID(id)
	if err != nil {
		return UserEntity{}, err
	}

	return user, nil
}
