package models

import (
	"database/sql"
	"time"
)

type ProductVariant struct {
	Id              int          `db:"id" json:"id"`
	Name            *string      `db:"name" json:"name" form:"name"`
	ProductId       *int         `db:"productId" json:"productId" form:"productId"`
	AdditionalPrice int          `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProductVariant struct {
	Data  []ProductVariant
	Count int
}

func FindAllProductVariant(limit int, offset int) (InfoProductVariant, error) {
	sql := `SELECT * FROM "productVariant" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productVariant"`
	result := InfoProductVariant{}
	data := []ProductVariant{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductVariant(id int) (ProductVariant, error) {
	sql := `SELECT * FROM "productVariant" WHERE id=$1`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductVariant(data ProductVariant) (ProductVariant, error) {
	sql := `
	INSERT INTO "productVariant" ("name", "productId",  "additionalPrice")
	VALUES (:name, :productId, :additionalPrice)
	RETURNING *`

	result := ProductVariant{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductVariant(data ProductVariant) (ProductVariant, error) {
	sql :=
		`UPDATE "productVariant" SET 
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := ProductVariant{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductVariant(id int) (ProductVariant, error) {
	sql := `DELETE FROM "productVariant" WHERE "id" = $1 RETURNING *`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}
