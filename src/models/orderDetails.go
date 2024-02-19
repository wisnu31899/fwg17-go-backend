package models

import (
	"database/sql"
	"fmt"
	"time"
)

type OrderDetails struct {
	Id               int          `db:"id" json:"id"`
	ProductId        *int         `db:"productId" json:"productId" form:"productId"`
	ProductSizeId    *int         `db:"productSizeId" json:"productSizeId" form:"productSizeId"`
	ProductVariantId *int         `db:"productVariantId" json:"productVariantId" form:"productVariantId"`
	Quantity         *int         `db:"quantity" json:"quantity" form:"quantity"`
	OrderId          *int         `db:"orderId" json:"orderId" form:"orderId"`
	SubTotal         *int         `db:"subTotal" json:"subTotal" form:"subTotal"`
	CreatedAt        *time.Time   `db:"createdAt" json:"createdAt"`
	UpdatedAt        sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoOrderDetails struct {
	Data  []OrderDetails
	Count int
}

func FindAllOrderDetails(limit int, offset int) (InfoOrderDetails, error) {
	sql := `SELECT * FROM "orderDetails" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "orderDetails"`
	result := InfoOrderDetails{}
	data := []OrderDetails{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)

	err := row.Scan(&result.Count)
	fmt.Println(err)
	return result, err
}

func FindOneOrderDetail(id int) (OrderDetails, error) {
	sql := `SELECT * FROM "orderDetails" WHERE id=$1`
	data := OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateOrderDetail(data OrderDetails) (OrderDetails, error) {
	sql := `
	INSERT INTO "orderDetails" ("productId","productSizeId","productVariantId","quantity","orderId","subTotal") VALUES
	(:productId, :productSizeId, :productVariantId, :quantity, :orderId, :subTotal)
	RETURNING *`

	result := OrderDetails{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateOrderDetail(data OrderDetails) (OrderDetails, error) {
	sql :=
		`UPDATE "orderDetails" SET 
		"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"productSizeId"=COALESCE(NULLIF(:productSizeId,0),"productSizeId"),
	"productVariantId"=COALESCE(NULLIF(:productVariantId,0),"productVariantId"),
	"quantity"=COALESCE(NULLIF(:quantity,0),"quantity"),
	"orderId"=COALESCE(NULLIF(:orderId,0),"orderId"),
	"subTotal"=COALESCE(NULLIF(:subTotal,0),"subTotal"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := OrderDetails{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteOrderDetail(id int) (OrderDetails, error) {
	sql := `DELETE FROM "orderDetails" WHERE "id" = $1 RETURNING *`
	data := OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}
