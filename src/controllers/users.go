package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
)

type pageInfo struct {
	Page string `json:"page"`
}

type responselist struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo pageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type User struct {
	Id       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type responseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func listAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &responselist{
		Success: true,
		Message: "list all users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: users,
	})
}

func listOneUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.GetOneUsers()
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &responseOnly{
				Success: false,
				Message: "user not found",
			})
		}
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "detail user",
		Results: user,
	})
}
