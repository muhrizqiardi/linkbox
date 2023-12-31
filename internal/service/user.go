package service

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(payload request.CreateUserRequest) (model.UserModel, error)
	GetOneByID(id int) (model.UserModel, error)
	GetOneByUsername(username string) (model.UserModel, error)
	UpdateOneByID(id int, payload request.UpdateUserRequest) (model.UserModel, error)
	DeleteOneByID(id int) (model.UserModel, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (s *userService) Create(payload request.CreateUserRequest) (model.UserModel, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return model.UserModel{}, constant.ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return model.UserModel{}, constant.ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserModel{}, err
	}
	user, err := s.repo.CreateUser(payload.Username, string(hashedPassword))
	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (s *userService) GetOneByID(id int) (model.UserModel, error) {
	user, err := s.repo.GetOneUserByID(id)
	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (s *userService) GetOneByUsername(username string) (model.UserModel, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(username) {
		return model.UserModel{}, constant.ErrInvalidUsername
	}
	user, err := s.repo.GetOneUserByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserModel{}, constant.ErrUserNotFound
		}

		return model.UserModel{}, err
	}

	return user, nil
}

func (s *userService) UpdateOneByID(id int, payload request.UpdateUserRequest) (model.UserModel, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return model.UserModel{}, constant.ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return model.UserModel{}, constant.ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserModel{}, err
	}
	user, err := s.repo.UpdateUserByID(id, payload.Username, string(hashedPassword))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserModel{}, constant.ErrUserNotFound
		}

		return model.UserModel{}, err
	}

	return user, nil
}

func (s *userService) DeleteOneByID(id int) (model.UserModel, error) {
	user, err := s.repo.DeleteUserByID(id)
	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}
