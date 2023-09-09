package user

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUser(username string, password string) (UserEntity, error)
	GetOneUserByID(id int) (UserEntity, error)
	GetOneUserByUsername(username string) (UserEntity, error)
	UpdateUserByID(id int, username string, password string) (UserEntity, error)
	DeleteUserByID(id int) (UserEntity, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(username string, password string) (UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return UserEntity{}, err
	}

	var user UserEntity
	if err := stmt.Get(&user, username, password); err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (r *repository) GetOneUserByID(id int) (UserEntity, error) {
	stmt, err := r.db.Preparex(QueryGetOneUserByID)
	if err != nil {
		return UserEntity{}, err
	}

	var user UserEntity
	if err := stmt.Get(&user, id); err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (r *repository) GetOneUserByUsername(username string) (UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return UserEntity{}, err
	}

	var user UserEntity
	if err := stmt.Get(&user, username); err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (r *repository) UpdateUserByID(id int, username string, password string) (UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return UserEntity{}, err
	}

	var user UserEntity
	if err := stmt.Get(&user, id, username, password); err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (r *repository) DeleteUserByID(id int) (UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return UserEntity{}, err
	}

	var user UserEntity
	if err := stmt.Get(&user, id); err != nil {
		return UserEntity{}, err
	}

	return user, nil
}
