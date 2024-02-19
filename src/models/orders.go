package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Orders struct {
	Id              int          `db:"id" json:"id"`
	UserId          *int         `db:"userId" json:"userId" form:"userId"`
	OrderNumber     string       `db:"orderNumber" json:"orderNumber" form:"orderNumber"`
	PromoId         *int         `db:"promoId" json:"promoId" form:"promoId"`
	Total           *int         `db:"total" json:"total" form:"total"`
	TaxAmount       *int         `db:"taxAmount" json:"taxAmount" form:"taxAmount"`
	Status          *string      `db:"status" json:"status" form:"status"`
	DeliveryAddress *string      `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName        string       `db:"fullName" json:"fullName" form:"fullName"`
	Email           string       `db:"email" json:"email" form:"email"`
	CreatedAt       *time.Time   `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoOrders struct {
	Data  []Orders
	Count int
}

func FindAllOrders(limit int, offset int) (InfoOrders, error) {
	sql := `SELECT * FROM "orders" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "orders"`
	result := InfoOrders{}
	data := []Orders{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)

	err := row.Scan(&result.Count)
	fmt.Println(err)
	return result, err
}

func FindOneOrder(id int) (Orders, error) {
	sql := `SELECT * FROM "orders" WHERE id=$1`
	data := Orders{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateOrder(data Orders) (Orders, error) {
	sql := `
	INSERT INTO "orders" ("userId","orderNumber","promoId","total","taxAmount","status","deliveryAddress","fullName","email") VALUES
	(:userId,:orderNumber,:promoId,:total,:taxAmount,:status,:deliveryAddress,:fullName,:email)
	RETURNING *`

	result := Orders{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateOrder(data Orders) (Orders, error) {
	sql :=
		`UPDATE "orders" SET 
		"userId"=COALESCE(NULLIF(:userId,0),"userId"),
	"orderNumber"=COALESCE(NULLIF(:orderNumber,''),"orderNumber"),
	"promoId"=COALESCE(NULLIF(:promoId,0),"promoId"),
	"total"=COALESCE(NULLIF(:total,0),"total"),
	"taxAmount"=COALESCE(NULLIF(:taxAmount,0),"taxAmount"),
	"status"=COALESCE(NULLIF(:status,''),"status"),
	"deliveryAddress"=COALESCE(NULLIF(:deliveryAddress,''),"deliveryAddress"),
	"fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	"email"=COALESCE(NULLIF(:email,''),"email"),
		"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *`

	result := Orders{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteOrder(id int) (Orders, error) {
	sql := `DELETE FROM "orders" WHERE "id" = $1 RETURNING *`
	data := Orders{}
	err := db.Get(&data, sql, id)
	return data, err
}
