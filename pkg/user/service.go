package user

import (
	"errors"
	"regexp"

	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidUsername error = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
var ErrConfirmPasswordNotMatched error = errors.New("Confirm Password field should be equal to Password")

type folderService interface {
}

type service struct {
	repo common.UserRepository
	fs   common.FolderService
}

func NewService(repo common.UserRepository, fs common.FolderService) *service {
	return &service{repo, fs}
}

func (s *service) Create(payload common.CreateUserDTO) (common.UserEntity, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return common.UserEntity{}, ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return common.UserEntity{}, ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.UserEntity{}, err
	}
	user, err := s.repo.CreateUser(payload.Username, string(hashedPassword))
	if err != nil {
		return common.UserEntity{}, err
	}

	newDefaultFolderPayload := common.CreateFolderDTO{
		UniqueName: "default",
		UserID:     user.ID,
	}
	if _, err := s.fs.Create(newDefaultFolderPayload); err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (s *service) GetOneByID(id int) (common.UserEntity, error) {
	user, err := s.repo.GetOneUserByID(id)
	if err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (s *service) GetOneByUsername(username string) (common.UserEntity, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(username) {
		return common.UserEntity{}, ErrInvalidUsername
	}
	user, err := s.repo.GetOneUserByUsername(username)
	if err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (s *service) UpdateOneByID(id int, payload common.UpdateUserDTO) (common.UserEntity, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return common.UserEntity{}, ErrInvalidUsername
	}
	if payload.Password != payload.ConfirmPassword {
		return common.UserEntity{}, ErrConfirmPasswordNotMatched
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.UserEntity{}, err
	}
	user, err := s.repo.UpdateUserByID(id, payload.Username, string(hashedPassword))
	if err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (s *service) DeleteOneByID(id int) (common.UserEntity, error) {
	user, err := s.repo.DeleteUserByID(id)
	if err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}
