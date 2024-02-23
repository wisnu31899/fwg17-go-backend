package admin

import (
	// "fmt"

	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/lib"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func GetAllProducts(c *gin.Context) {

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
	result, err := models.FindAllProducts(keyword, limit, offset, sortField, sortOrder)

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
		Message:  "get list all products successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func GetDetailProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.FindOneProduct(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "product not found",
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
		Message: "get detail product successfully",
		Results: product,
	})
}

func CreateProduct(c *gin.Context) {
	data := models.Product{}

	err := c.ShouldBind(&data)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	// Menangkap kedua nilai yang dikembalikan oleh lib.UploadFile
	image, err := lib.UploadFileProduct(c, "products")
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
				Message: "failed to upload image",
			})
		}
		return
	}

	data.Image = &image

	product, err := models.CreateProduct(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "create product successfully",
		Results: product,
	})
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.Product{}

	c.Bind(&data)
	data.Id = id

	// Mendapatkan nama file foto lama sebelum menggantinya dengan foto baru
	oldProduct, _ := models.FindOneProduct(id)
	var oldImage *string
	if oldProduct.Image != nil {
		oldImage = oldProduct.Image
	}

	// Menangkap kedua nilai yang dikembalikan oleh lib.UploadFile
	image, err := lib.UploadFileProduct(c, "products")
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
				Message: "failed to upload image",
			})
		}
		return
	}

	data.Image = &image

	product, err := models.UpdateProduct(data)

	if err != nil {
		// log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	// Memeriksa apakah foto lama ada
	if oldImage != nil {
		// 	// Menghapus foto lama dari folder penyimpanan
		err := os.Remove(*oldImage)
		if err != nil {
			log.Println("Failed to delete old product:", err)
		}
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "update product successfully",
		Results: product,
	})
}

func DeleteProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.DeleteProduct(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "product not found",
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
		Message: "delete product successfully",
		Results: product,
	})
}
