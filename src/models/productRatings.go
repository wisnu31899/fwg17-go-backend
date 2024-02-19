package models

import (
	"database/sql"
	"fmt"
	"time"
)

type ProductRatings struct {
	Id            int          `db:"id" json:"id"`
	ProductId     *int         `db:"productId" json:"productId" form:"productId"`
	Rate          int          `db:"rate" json:"rate" form:"rate"`
	ReviewMessage *string      `db:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	UserId        *int         `db:"userId" json:"userId" form:"userId"`
	CreatedAt     *time.Time   `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProductRatings struct {
	Data  []ProductRatings
	Count int
}

func FindAllProductRatings(limit int, offset int) (InfoProductRatings, error) {
	sql := `SELECT * FROM "productRatings" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productRatings"`
	result := InfoProductRatings{}
	data := []ProductRatings{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)

	err := row.Scan(&result.Count)
	fmt.Println(err)
	return result, err
}

func FindOneProductRating(id int) (ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE id=$1`
	data := ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductRating(data ProductRatings) (ProductRatings, error) {
	sql := `
	INSERT INTO "productRatings" ("productId","rate","reviewMessage","userId") VALUES
	(:productId, :rate, :reviewMessage, :userId)
	RETURNING *`

	result := ProductRatings{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductRating(data ProductRatings) (ProductRatings, error) {
	sql :=
		`UPDATE "productRatings" SET 
		"productId"=COALESCE(NULLIF(:productId,0),"productId"),
		"rate"=COALESCE(NULLIF(:rate,0),"rate"),
		"reviewMessage"=COALESCE(NULLIF(:reviewMessage,''),"reviewMessage"),
		"userId"=COALESCE(NULLIF(:userId,0),"userId"),
		"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := ProductRatings{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductRating(id int) (ProductRatings, error) {
	sql := `DELETE FROM "productRatings" WHERE "id" = $1 RETURNING *`
	data := ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}
