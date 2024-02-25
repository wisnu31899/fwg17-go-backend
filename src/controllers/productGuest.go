package controllers

import (
	// "fmt"

	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

// func GetDetailProductVariantAndSize(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
// 			Success: false,
// 			Message: "invalid product ID",
// 		})
// 		return
// 	}

// 	product, err := models.FindDetailProductVariantAndSize(id)
// 	fmt.Println(err)
// 	if err != nil {
// 		if strings.HasPrefix(err.Error(), "sql: no rows") {
// 			c.JSON(http.StatusNotFound, &services.ResponseOnly{
// 				Success: false,
// 				Message: "product not found",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
// 			Success: false,
// 			Message: "internal server error",
// 		})
// 		return
// 	}

// 	// Penggabungan produk berdasarkan atribut 'sizes' dan 'variants'
// 	mergedProduct := product

// 	// Penggabungan 'sizes'
// 	for _, size := range product.Sizes {
// 		if !containsSize(mergedProduct.Sizes, size.Id) {
// 			mergedProduct.Sizes = append(mergedProduct.Sizes, size)
// 		}
// 	}

// 	// Penggabungan 'variants'
// 	for _, variant := range product.Variants {
// 		if !containsVariant(mergedProduct.Variants, variant.Id) {
// 			mergedProduct.Variants = append(mergedProduct.Variants, variant)
// 		}
// 	}

// 	c.JSON(http.StatusOK, &services.ResponseDetail{
// 		Success: true,
// 		Message: "get detail product successfully",
// 		Results: mergedProduct,
// 	})
// }

// // Fungsi bantu untuk menentukan apakah ID ukuran sudah ada dalam slice 'Sizes'
// func containsSize(size []models.Size, id int) bool {
// 	for _, size := range size {
// 		if size.Id == id {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Fungsi bantu untuk menentukan apakah ID varian sudah ada dalam slice 'Variants'
// func containsVariant(variant []models.Variant, id int) bool {
// 	for _, variant := range variant {
// 		if variant.Id == id {
// 			return true
// 		}
// 	}
// 	return false
// }
