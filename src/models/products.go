package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Product struct {
	Id            int          `db:"id" json:"id"`
	Name          string       `db:"name" json:"name" form:"name"`
	Description   *string      `db:"description" json:"description" form:"description"`
	BasePrice     int          `db:"basePrice" json:"basePrice" form:"basePrice"`
	Image         *string      `db:"image" json:"image"`
	Discount      *int         `db:"discount" json:"discount" form:"discount"`
	IsRecommended *bool        `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	Stock         *int         `db:"stock" json:"stock" form:"stock"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProduct struct {
	Data  []Product
	Count int
}

func FindAllProducts(keyword string, limit int, offset int, sortField string, sortOrder string) (InfoProduct, error) {
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
	} else if strings.ToLower(sortField) == "name" {
		sortColumn = "name"
	} else if strings.ToLower(sortField) == "basePrice" {
		sortColumn = "basePrice"
	} else if strings.ToLower(sortField) == "stock" {
		sortColumn = "stock"
	} else {
		sortColumn = "id"
	}

	sql := fmt.Sprintf(`SELECT * FROM "products" WHERE "name" ILIKE $1 ORDER BY "%s" %s LIMIT $2 OFFSET $3`, sortColumn, sortDirection)
	sqlCount := `SELECT COUNT(*) FROM "products" WHERE "name" ILIKE $1`
	result := InfoProduct{}
	data := []Product{}
	err := db.Select(&data, sql, "%"+keyword+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+keyword+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProduct(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE id=$1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProduct(data Product) (Product, error) {
	sql := `
	INSERT INTO "products" ("name", "description",  "basePrice", "image", "discount", "isRecommended", "stock")
	VALUES (:name, :description, :basePrice, :image, :discount, :isRecommended, :stock)
	RETURNING *`

	result := Product{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProduct(data Product) (Product, error) {
	sql :=
		`UPDATE "products" SET 
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"basePrice"=COALESCE(NULLIF(:basePrice, 0),"basePrice"),
	"image"=COALESCE(NULLIF(:image, ''),"image"),
	"discount"=COALESCE(NULLIF(:discount, 0),"discount"),
	"isRecommended"=COALESCE(NULLIF(:isRecommended, false),"isRecommended"),
	"stock"=COALESCE(NULLIF(:stock, 0),"stock"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := Product{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProduct(id int) (Product, error) {
	sql := `DELETE FROM "products" WHERE "id" = $1 RETURNING *`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductByName(name string) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "name" = $1`
	data := Product{}
	err := db.Get(&data, sql, name)
	return data, err
}
