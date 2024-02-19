package models

import (
	"database/sql"
	"time"
)

type ProductTag struct {
	Id        int          `db:"id" json:"id"`
	ProductId *int         `db:"productId" json:"productId"`
	TagId     *int         `db:"tagId" json:"tagId"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProductTag struct {
	Data  []ProductTag
	Count int
}

func FindAllProductTag(limit int, offset int) (InfoProductTag, error) {
	sql := `SELECT * FROM "productTags" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productTags"`
	result := InfoProductTag{}
	data := []ProductTag{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductTag(id int) (ProductTag, error) {
	sql := `SELECT * FROM "productTags" WHERE id=$1`
	data := ProductTag{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductTag(data ProductTag) (ProductTag, error) {
	sql := `
	INSERT INTO "productTags" ("name","productId","tagId") VALUES
	(:name, :productId, :tagId)
	RETURNING *`

	result := ProductTag{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductTag(data ProductTag) (ProductTag, error) {
	sql :=
		`UPDATE "productTags" SET
		"name"=COALESCE(NULLIF(:name,''),"name"),
		"productId"=COALESCE(NULLIF(:productId,0),"productId"),
		"tagId"=COALESCE(NULLIF(:tagId,0),"tagId"),
		"updatedAt"=NOW()
		WHERE id = :id
		RETURNING *`

	result := ProductTag{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductTag(id int) (ProductTag, error) {
	sql := `DELETE FROM "productTags" WHERE "id" = $1 RETURNING *`
	data := ProductTag{}
	err := db.Get(&data, sql, id)
	return data, err
}
