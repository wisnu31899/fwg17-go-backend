package models

import (
	"database/sql"
	"time"
)

type Categories struct {
	Id        int          `db:"id" json:"id"`
	Name      string       `db:"name" json:"name" form:"name"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoCategories struct {
	Data  []Categories
	Count int
}

func FindAllCategories(limit int, offset int) (InfoCategories, error) {
	sql := `SELECT * FROM "categories" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "categories"`
	result := InfoCategories{}
	data := []Categories{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneCategories(id int) (Categories, error) {
	sql := `SELECT * FROM "categories" WHERE id=$1`
	data := Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateCategories(data Categories) (Categories, error) {
	sql := `
	INSERT INTO "categories" ("name")
	VALUES (:name)
	RETURNING *`

	result := Categories{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateCategories(data Categories) (Categories, error) {
	sql :=
		`UPDATE "categories" SET 
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := Categories{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteCategories(id int) (Categories, error) {
	sql := `DELETE FROM "categories" WHERE "id" = $1 RETURNING *`
	data := Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}
