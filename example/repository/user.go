package repository

import (
	"backend_example/model"

	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (r *User) FindByID(id int64) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Get(user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	return user, nil
}
