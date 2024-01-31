package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pageInfo struct {
	Page string `json:"page"`
}

type response struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo pageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.GET("/:users", func(c *gin.Context) {
		page := c.Query("page")
		data := []User{
			{
				Id:       1,
				Email:    "ADMIN@gmail.com",
				Password: "1234",
			},
			{
				Id:       2,
				Email:    "GUEST@gmail.com",
				Password: "1234",
			},
		}
		c.JSON(http.StatusOK, &response{
			Success: true,
			Message: "OK",
			PageInfo: pageInfo{
				Page: page,
			},
			Results: data,
		})
	})
	r.Run()
}
