package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(username string, password string) (common.UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return common.UserEntity{}, err
	}

	var user common.UserEntity
	if err := stmt.Get(&user, username, password); err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (r *repository) GetOneUserByID(id int) (common.UserEntity, error) {
	stmt, err := r.db.Preparex(QueryGetOneUserByID)
	if err != nil {
		return common.UserEntity{}, err
	}

	var user common.UserEntity
	if err := stmt.Get(&user, id); err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (r *repository) GetOneUserByUsername(username string) (common.UserEntity, error) {
	stmt, err := r.db.Preparex(QueryGetOneUserByUsername)
	if err != nil {
		return common.UserEntity{}, err
	}

	var user common.UserEntity
	if err := stmt.Get(&user, username); err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (r *repository) UpdateUserByID(id int, username string, password string) (common.UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return common.UserEntity{}, err
	}

	var user common.UserEntity
	if err := stmt.Get(&user, id, username, password); err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}

func (r *repository) DeleteUserByID(id int) (common.UserEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateUser)
	if err != nil {
		return common.UserEntity{}, err
	}

	var user common.UserEntity
	if err := stmt.Get(&user, id); err != nil {
		return common.UserEntity{}, err
	}

	return user, nil
}
