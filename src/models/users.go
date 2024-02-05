package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wisnu31899/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	// Email       string       `db:"email" json:"email" form:"email" binding:"email"` //penggunaan shouldbinding
	Id          int          `db:"id" json:"id"`
	FullName    string       `db:"fullName" json:"fullName" form:"fullName"`
	Email       string       `db:"email" json:"email" form:"email"`
	Password    string       `db:"password" json:"password" form:"password"`
	Address     string       `db:"address" json:"address" form:"address"`
	PhoneNumber string       `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	RoleId      int          `db:"roleId" json:"roleId" form:"roleId"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type Info struct {
	Data  []User
	Count int
}

func FindAllUsers(limit int, offset int) (Info, error) {
	sql := `SELECT * FROM "users" LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "users"`
	result := Info{}
	data := []User{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneUser(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE id=$1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateUser(data User) (User, error) {
	sql := `
	INSERT INTO "users" ("fullName", "email", "password", "address", "phoneNumber", "roleId") VALUES
	(:fullName, :email, :password, :address, :phoneNumber, :roleId)
	RETURNING *`

	// sql := `
	// INSERT INTO "users" ( "email", "password") VALUES
	// (:email, :password)
	// RETURNING *`

	result := User{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateUser(data User) (User, error) {
	// sql := `
	// UPDATE "users" SET
	// fullName=COALESCE(NULLIF(:fullName,''),fullName),
	// email=COALESCE(NULLIF(:email,''),email),
	// password=COALESCE(NULLIF(:password,''),password),
	// address=COALESCE(NULLIF(:address,''),address),
	// phoneNumber=COALESCE(NULLIF(:phoneNumber,''),phoneNumber),
	// roleId=COALESCE(NULLIF(:roleId,''),roleId)
	// WHERE id=:id
	// RETURNING *`

	sql := `
	UPDATE "users" SET
	"email"=COALESCE(NULLIF(:email,''),"email"),
	"password"=COALESCE(NULLIF(:password,''),"password"),
	"address"=COALESCE(NULLIF(:address,''),"address"),
	"updatedAt" = NOW()
	WHERE id=:id
	RETURNING *`

	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteUser(id int) (User, error) {
	sql := `DELETE FROM "users" WHERE id=$1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindUserByEmail(email string) (User, error) {
	sql := `SELECT * FROM "users" WHERE email=$1`
	data := User{}
	err := db.Get(&data, sql, email)
	return data, err
}
