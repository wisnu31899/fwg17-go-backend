package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wisnu31899/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	// Email       string       `db:"email" json:"email" form:"email" binding:"email"` //penggunaan shouldbinding
	Id          int          `db:"id" json:"id"`
	FullName    *string      `db:"fullName" json:"fullName" form:"fullName"`
	Email       string       `db:"email" json:"email" form:"email"`
	Password    string       `db:"password" json:"password" form:"password"`
	Address     *string      `db:"address" json:"address" form:"address"`
	Picture     *string      `db:"picture" json:"picture"`
	PhoneNumber *string      `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	RoleId      *int         `db:"roleId" json:"roleId" form:"roleId"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type Info struct {
	Data  []User
	Count int
}

func FindAllUsers(keyword string, limit int, offset int, sortField string, sortOrder string) (Info, error) {
	var sortDirection string
	if strings.ToLower(sortOrder) == "asc" {
		sortDirection = "ASC"
	} else if strings.ToLower(sortOrder) == "desc" {
		sortDirection = "DESC"
	} else {
		sortDirection = "ASC"
	}

	var sortColumn string
	if strings.ToLower(sortField) == "createdAt" {
		sortColumn = "createdAt"
	} else {
		sortColumn = "id"
	}

	sql := fmt.Sprintf(`SELECT * FROM "users" WHERE "fullName" ILIKE $1 ORDER BY "%s" %s LIMIT $2 OFFSET $3`, sortColumn, sortDirection)
	sqlCount := `SELECT COUNT(*) FROM "users" WHERE "fullName" ILIKE $1`
	result := Info{}
	data := []User{}
	err := db.Select(&data, sql, "%"+keyword+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+keyword+"%")
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
	// Periksa apakah RoleId nil, jika ya, atur nilainya ke 1
	if data.RoleId == nil {
		defaultRoleId := 1
		data.RoleId = &defaultRoleId
	}
	sql := `
	INSERT INTO "users" ("fullName", "email", "password", "address", "picture", "phoneNumber", "roleId") VALUES
	(:fullName, :email, :password, :address, :picture, :phoneNumber, :roleId)
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
	sql := `
	UPDATE "users" SET
	"fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	"email"=COALESCE(NULLIF(:email,''),"email"),
	"password"=COALESCE(NULLIF(:password,''),"password"),
	"address"=COALESCE(NULLIF(:address,''),"address"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber,''),"phoneNumber"),
	"picture"=COALESCE(NULLIF(:picture,''),"picture"),
	"roleId"=COALESCE(NULLIF(CAST(:roleId AS INT),NULL),"roleId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	// sql := `
	// UPDATE "users" SET
	// "email"=COALESCE(NULLIF(:email,''),"email"),
	// "password"=COALESCE(NULLIF(:password,''),"password"),
	// "address"=COALESCE(NULLIF(:address,''),"address"),
	// "updatedAt" = NOW()
	// WHERE id=:id
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
