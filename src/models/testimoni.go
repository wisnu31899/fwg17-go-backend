package models

import (
	"database/sql"
	"time"
)

type Testimoni struct {
	Id        int          `db:"id" json:"id"`
	FullName  string       `db:"fullName" json:"fullName" form:"fullName"`
	RoleId    int          `db:"roleId" json:"roleId" form:"roleId"`
	Rating    int          `db:"rating" json:"rating" form:"rating"`
	Review    string       `db:"review" json:"review" form:"review"`
	Picture   string       `db:"picture" json:"picture" form:"picture"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type TestimoniForm struct {
	Id        int          `db:"id" json:"id"`
	FullName  *string      `db:"fullName" json:"fullName" form:"fullName"`
	RoleId    *int         `db:"roleId" json:"roleId" form:"roleId"`
	Rating    *int         `db:"rating" json:"rating" form:"rating" binding:"required,eq=5|eq=4|eq=3|eq=2|eq=1"`
	Review    *string      `db:"review" json:"review" form:"review"`
	Picture   string       `db:"picture" json:"picture"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type infoTestimoni struct {
	Data  []Testimoni
	Count int
}

func FindAllTestimoni(keyword string, sortBy string, order string, limit int, offset int) (InfoTs, error) {
	sql := `
	SELECT * FROM "testimoni" 
	WHERE "fullName" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "testimoni"
	WHERE "fullName" ILIKE $1
	`

	result := infoTestimoni{}
	data := []Testimoni{}
	err := db.Select(&data, sql, "%"+keyword+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+keyword+"%")
	err = row.Scan(&result.Count)

	return result, err
}
