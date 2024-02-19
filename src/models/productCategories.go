package models

import (
	"database/sql"
	"time"
)

type ProductCategories struct {
	Id         int          `db:"id" json:"id"`
	ProductId  *int         `db:"productId" json:"productId" form:"productId"`
	CategoryId *int         `db:"categoryId" json:"categoryId" form:"categoryId"`
	CreatedAt  time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt  sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProductCategories struct {
	Data  []ProductCategories
	Count int
}

func FindAllProductCategories(limit int, offset int) (InfoProductCategories, error) {
	sql := `SELECT * FROM "productCategories" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productCategories"`
	result := InfoProductCategories{}
	data := []ProductCategories{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductCategories(id int) (ProductCategories, error) {
	sql := `SELECT * FROM "productCategories" WHERE id=$1`
	data := ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql := `
	INSERT INTO "productCategories" ("productId","categoryId") VALUES
	(:productId,:categoryId)
	RETURNING *`

	result := ProductCategories{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql :=
		`UPDATE "productCategories" SET 
		"productId"=COALESCE(NULLIF(:productId,''),"productId"),
		"categoryId"=COALESCE(NULLIF(:categoryId,''),"categoryId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := ProductCategories{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductCategories(id int) (ProductCategories, error) {
	sql := `DELETE FROM "ProductCategories" WHERE "id" = $1 RETURNING *`
	data := ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}
