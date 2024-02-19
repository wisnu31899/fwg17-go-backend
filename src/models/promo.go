package models

import (
	"database/sql"
	"time"
)

type Promo struct {
	Id            int          `db:"id" json:"id"`
	Name          string       `db:"name" json:"name" form:"name"`
	Code          string       `db:"code" json:"code" form:"code"`
	Description   *string      `db:"description" json:"description" form:"description"`
	Percentage    float64      `db:"percentage" json:"percentage" form:"percentage"`
	IsExpired     *bool        `db:"isExpired" json:"isExpired" form:"isExpired"`
	MaximumPromo  *int         `db:"maximumPromo" json:"maximumPromo" form:"maximumPromo"`
	MinimumAmount *int         `db:"minimumAmount" json:"minimumAmount" form:"minimumAmount"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoPromo struct {
	Data  []Promo
	Count int
}

func FindAllPromo(limit int, offset int) (InfoPromo, error) {
	sql := `SELECT * FROM "promo" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "promo"`
	result := InfoPromo{}
	data := []Promo{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOnePromo(id int) (Promo, error) {
	sql := `SELECT * FROM "promo" WHERE id=$1`
	data := Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreatePromo(data Promo) (Promo, error) {
	sql := `
	INSERT INTO "promo" ("name","code","description","percentage","isExpired","maximumPromo","minimumAmount") 
	VALUES
	(:name, :code, :description, :percentage, :isExpired, :maximumPromo, :minimumAmount)
	RETURNING *`

	result := Promo{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdatePromo(data Promo) (Promo, error) {
	sql :=
		`UPDATE "promo" SET
		"name"=COALESCE(NULLIF(:name,''),"name"),
		"code"=COALESCE(NULLIF(:code,''),"code"),
		"description"=COALESCE(NULLIF(:description,''),"description"),
		"percentage"=COALESCE(NULLIF(:percentage,0.0),"percentage"),
		"isExpired"=COALESCE(NULLIF(:isExpired,false),"isExpired"),
		"maximumPromo"=COALESCE(NULLIF(:maximumPromo,0),"maximumPromo"),
		"minimumAmount"=COALESCE(NULLIF(:minimumAmount,0),"minimumAmount"),
		"updatedAt"=NOW()
		WHERE id = :id
		RETURNING *`

	result := Promo{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeletePromo(id int) (Promo, error) {
	sql := `DELETE FROM "Promo" WHERE "id" = $1 RETURNING *`
	data := Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}
