package admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/lib"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func GetAllUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	keyword := c.DefaultQuery("keyword", "")
	sortField := c.DefaultQuery("sortField", "id")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	if page <= 0 {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: false,
			Message: "no data",
		})
		return
	}

	offset := (page - 1) * limit
	result, err := models.FindAllUsers(keyword, limit, offset, sortField, sortOrder)

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

	if page > pageInfo.LastPage {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: false,
			Message: "no data",
		})
		return
	}

	nextPage := page + 1
	previousPage := page - 1

	if nextPage > pageInfo.LastPage {
		nextPage = 0
	}

	if previousPage <= 0 {
		previousPage = 0
	}

	pageInfo.NextPage = nextPage
	pageInfo.PreviousPage = previousPage

	c.JSON(http.StatusOK, &services.ResponseAll{
		Success:  true,
		Message:  "get list all users successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func GetDetailUser(c *gin.Context) {

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

func CreateUser(c *gin.Context) {
	data := models.User{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	// Menangkap kedua nilai yang dikembalikan oleh lib.UploadFile
	picture, err := lib.UploadFile(c, "profile")
	fmt.Println(err)
	if err != nil {
		if err.Error() == "file size exceeds the limit of 5MB" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "file size exceeds the limit of 5MB",
			})
		} else if err.Error() == "file format is not supported" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "file format is not supported. Only JPG or PNG files are allowed",
			})
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "failed to upload picture",
			})
		}
		return
	}

	// Menetapkan nilai Picture ke dalam data.User
	data.Picture = &picture

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

	// Mendapatkan nama file foto lama sebelum menggantinya dengan foto baru
	oldUser, _ := models.FindOneUser(id)
	var oldPicture *string
	if oldUser.Picture != nil {
		oldPicture = oldUser.Picture
	}

	// Menangkap kedua nilai yang dikembalikan oleh lib.UploadFile
	picture, err := lib.UploadFile(c, "profile")
	fmt.Println(picture)
	if err != nil {
		if err.Error() == "file size exceeds the limit of 5MB" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "file size exceeds the limit of 5MB",
			})
		} else if err.Error() == "file format is not supported" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "file format is not supported. Only JPG or PNG files are allowed",
			})
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "failed to upload picture",
			})
		}
		return
	}

	// Menetapkan nilai Picture ke dalam data.User
	data.Picture = &picture

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

	// Memeriksa apakah foto lama ada
	if oldPicture != nil {
		// Menghapus foto lama dari folder penyimpanan
		err := os.Remove(*oldPicture)
		if err != nil {
			log.Println("Failed to delete old picture:", err)
		}
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
