package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func ListAllTestimoni(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))

	if page <= 0 {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: false,
			Message: "no data",
		})
		return
	}
	offset := (page - 1) * limit

	result, err := models.FindAllTestimoni(limit, offset)

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
		Message:  "get list all testimoni successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailTestimoni(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	testimoni, err := models.FindOneTestimoni(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "testimoni not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "Detail testimoni",
		Results: testimoni,
	})
}

func ListAllTestimoniJoin(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))

	if page <= 0 {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: false,
			Message: "no data",
		})
		return
	}
	offset := (page - 1) * limit

	result, err := models.FindAllTestimoniJoin(keyword, sortBy, order, limit, offset)

	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		LastPage:  int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}
	fmt.Println(err)
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
		Message:  "get list all testimoni successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func CreateTestimoni(c *gin.Context) {
	data := models.TestimoniForm{}
	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	testimoni, err := models.CreateTestimoni(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "testimoni created successfully",
		Results: testimoni,
	})
}

func UpdateTestimoni(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.TestimoniForm{}

	c.ShouldBind(&data)
	data.Id = id

	// isExist, err := models.FindOneTestimoni(id)
	// if err != nil {
	// 	fmt.Println(isExist, err)
	// 	c.JSON(http.StatusNotFound, &services.ResponseOnly{
	// 		Success: false,
	// 		Message: "Testimoni not found",
	// 	})
	// 	return
	// }

	testimoni, err := models.UpdateTestimoni(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "Testimoni not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "Testimoni updated successfully",
		Results: testimoni,
	})
}

func DeleteTestimoni(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	testimoni, err := models.DeleteTestimoni(id)

	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "testimoni not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "Delete testimoni successfully",
		Results: testimoni,
	})
}
