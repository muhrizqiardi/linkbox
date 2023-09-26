package service

import (
	"errors"
	"regexp"

	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidUsername error = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
var ErrConfirmPasswordNotMatched error = errors.New("Confirm Password field should be equal to Password")

type UserService interface {
	Create(payload request.CreateUserRequest) (model.UserModel, error)
	GetOneByID(id int) (model.UserModel, error)
	GetOneByUsername(username string) (model.UserModel, error)
	UpdateOneByID(id int, payload request.UpdateUserRequest) (model.UserModel, error)
	DeleteOneByID(id int) (model.UserModel, error)
}

type userService struct {
	repo repository.UserRepository
	fs   FolderService
}

func NewUserService(repo repository.UserRepository, fs FolderService) *userService {
	return &userService{repo, fs}
}

func (s *userService) Create(payload request.CreateUserRequest) (model.UserModel, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return model.UserModel{}, ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return model.UserModel{}, ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserModel{}, err
	}
	user, err := s.repo.CreateUser(payload.Username, string(hashedPassword))
	if err != nil {
		return model.UserModel{}, err
	}

	newDefaultFolderPayload := request.CreateFolderRequest{
		UniqueName: "default",
		UserID:     user.ID,
	}
	// TODO: make the query create default folder by default
	if _, err := s.fs.Create(newDefaultFolderPayload); err != nil {
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
		return model.UserModel{}, ErrInvalidUsername
	}
	user, err := s.repo.GetOneUserByUsername(username)
	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (s *userService) UpdateOneByID(id int, payload request.UpdateUserRequest) (model.UserModel, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return model.UserModel{}, ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return model.UserModel{}, ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserModel{}, err
	}
	user, err := s.repo.UpdateUserByID(id, payload.Username, string(hashedPassword))
	if err != nil {
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
