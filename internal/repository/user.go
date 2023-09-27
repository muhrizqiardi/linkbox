package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/query"
)

type UserRepository interface {
	CreateUser(username string, password string) (model.UserModel, error)
	GetOneUserByID(id int) (model.UserModel, error)
	GetOneUserByUsername(username string) (model.UserModel, error)
	UpdateUserByID(id int, username string, password string) (model.UserModel, error)
	DeleteUserByID(id int) (model.UserModel, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(username string, password string) (model.UserModel, error) {
	stmt, err := r.db.Preparex(query.QueryCreateUser)
	if err != nil {
		return model.UserModel{}, err
	}

	var user model.UserModel
	if err := stmt.Get(&user, username, password); err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (r *userRepository) GetOneUserByID(id int) (model.UserModel, error) {
	stmt, err := r.db.Preparex(query.QueryGetOneUserByID)
	if err != nil {
		return model.UserModel{}, err
	}

	var user model.UserModel
	if err := stmt.Get(&user, id); err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (r *userRepository) GetOneUserByUsername(username string) (model.UserModel, error) {
	stmt, err := r.db.Preparex(query.QueryGetOneUserByUsername)
	if err != nil {
		return model.UserModel{}, err
	}

	var user model.UserModel
	if err := stmt.Get(&user, username); err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (r *userRepository) UpdateUserByID(id int, username string, password string) (model.UserModel, error) {
	stmt, err := r.db.Preparex(query.QueryUpdateOneUserByID)
	if err != nil {
		return model.UserModel{}, err
	}

	var user model.UserModel
	if err := stmt.Get(&user, id, username, password); err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(id int) (model.UserModel, error) {
	stmt, err := r.db.Preparex(query.QueryDeleteOneUserByID)
	if err != nil {
		return model.UserModel{}, err
	}

	var user model.UserModel
	if err := stmt.Get(&user, id); err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}
