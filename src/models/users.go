package models

import (
	"database/sql"
	"time"

	"github.com/wisnu31899/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id        int          `db:"id" json:"id"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt,omitempty" validate: "true"`
}

func GetAllUsers() ([]User, error) {
	sql := `SELECT * FROM "users"`
	data := []User{}
	err := db.Select(&data, sql)
	return data, err
}

func GetOneUsers(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE id=$1`
	data := User{}
	err := db.get(&data, sql, id)
	return data, err
}
