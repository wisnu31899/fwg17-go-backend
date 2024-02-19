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
