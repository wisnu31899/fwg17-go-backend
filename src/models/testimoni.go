package models

import (
	"database/sql"
	"time"
)

type TestimoniJoin struct {
	Id        int          `db:"id" json:"id"`
	Rating    int          `db:"rating" json:"rating"`
	Review    string       `db:"review" json:"review"`
	FullName  string       `db:"fullName" json:"fullName"`
	Picture   *string      `db:"picture" json:"picture"`
	RoleId    int          `db:"roleId" json:"roleId"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type infoTestimoniJoin struct {
	Data  []TestimoniJoin `json:"data"`
	Count int             `json:"count"`
}

type Testimoni struct {
	Id        int          `db:"id" json:"id"`
	UserId    int          `db:"userId" json:"userId" form:"userId"`
	Rating    int          `db:"rating" json:"rating" form:"rating"`
	Review    string       `db:"review" json:"review" form:"review"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type TestimoniForm struct {
	Id        int          `db:"id" json:"id"`
	UserId    *int         `db:"userId" json:"userId" form:"userId"`
	Rating    *int         `db:"rating" json:"rating" form:"rating"`
	Review    *string      `db:"review" json:"review" form:"review"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type infoTestimoni struct {
	Data  []Testimoni `json:"data"`
	Count int         `json:"count"`
}

func FindAllTestimoni(limit int, offset int) (infoTestimoni, error) {
	sql := `
	SELECT * FROM "testimoni" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "testimoni"
	`

	result := infoTestimoni{}
	data := []Testimoni{}
	err := db.Select(&data, sql, limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneTestimoni(id int) (Testimoni, error) {
	sql := `SELECT * FROM "testimoni" WHERE id = $1`
	data := Testimoni{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateTestimoni(data TestimoniForm) (TestimoniForm, error) {
	sql := `INSERT INTO "testimoni" ("userId", "rating", "review")
	VALUES
	(:userId, :rating, :review)
	RETURNING *
	`
	result := TestimoniForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateTestimoni(data TestimoniForm) (TestimoniForm, error) {
	sql := `UPDATE "testimoni" SET
	"userId"=COALESCE(NULLIF(:userId, 0),"userId"),
	"rating"=COALESCE(NULLIF(:rating, 0),"rating"),
	"review"=COALESCE(NULLIF(:review, ''),"review"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := TestimoniForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteTestimoni(id int) (TestimoniForm, error) {
	sql := `DELETE FROM "testimoni" WHERE id = $1 RETURNING *`
	data := TestimoniForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindAllTestimoniJoin(keyword string, sortBy string, order string, limit int, offset int) (infoTestimoniJoin, error) {

	sql := `
    SELECT ts.id, ts.rating, ts.review, ts."createdAt", ts."updatedAt", u."fullName", u."picture", u."roleId"
    FROM "testimoni" ts
    JOIN "users" u ON ts."userId" = u."id"
    WHERE u."fullName" ILIKE $1
    ORDER BY ts.` + sortBy + ` ` + order + `
    LIMIT $2 OFFSET $3
`
	sqlCount := `
    SELECT COUNT(*) FROM "testimoni" ts
    JOIN "users" u ON ts."userId" = u."id"
    WHERE u."fullName" ILIKE $1
`

	result := infoTestimoniJoin{}
	data := []TestimoniJoin{}
	err := db.Select(&data, sql, "%"+keyword+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+keyword+"%")
	err = row.Scan(&result.Count)

	return result, err
}
