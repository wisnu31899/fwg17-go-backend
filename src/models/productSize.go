package models

import (
	"database/sql"
	"time"
)

type ProductSize struct {
	Id              int          `db:"id" json:"id"`
	Size            *string      `db:"size" json:"size" form:"size"`
	ProductId       *int         `db:"productId" json:"productId" form:"productId"`
	AdditionalPrice int          `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProductSize struct {
	Data  []ProductSize
	Count int
}

func FindAllProductSize(limit int, offset int) (InfoProductSize, error) {
	sql := `SELECT * FROM "productSize" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productSize"`
	result := InfoProductSize{}
	data := []ProductSize{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductSize(id int) (ProductSize, error) {
	sql := `SELECT * FROM "productSize" WHERE id=$1`
	data := ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductSize(data ProductSize) (ProductSize, error) {
	sql := `
	INSERT INTO "productSize" ("size", "productId",  "additionalPrice")
	VALUES (:size, :productId, :additionalPrice)
	RETURNING *`

	result := ProductSize{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductSize(data ProductSize) (ProductSize, error) {
	sql :=
		`UPDATE "productSize" SET 
	"size"=COALESCE(NULLIF(:size, ''),"size"),
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := ProductSize{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductSize(id int) (ProductSize, error) {
	sql := `DELETE FROM "productSize" WHERE "id" = $1 RETURNING *`
	data := ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}
