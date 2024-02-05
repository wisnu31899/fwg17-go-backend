package admin

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func GetAllUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit
	result, err := models.FindAllUsers(limit, offset)

	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		LastPage:  int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseAll{
		Success:  true,
		Message:  "get list all users successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func GetDetailUsers(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindOneUser(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "user not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "get detail user successfully",
		Results: user,
	})
}

func CreateUsers(c *gin.Context) {
	data := models.User{}

	// c.Bind(&data)

	// belum digunakan
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed hash password",
		})
		return
	}

	data.Password = hash.String()

	user, err := models.CreateUser(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "create user successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.User{}

	c.Bind(&data)
	data.Id = id

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed hash password",
		})
		return
	}

	data.Password = hash.String()

	user, err := models.UpdateUser(data)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "update user successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.DeleteUser(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "user not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "delete user successfully",
		Results: user,
	})
}
