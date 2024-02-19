package admin

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

func GetAllProductRatings(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit
	result, err := models.FindAllProductRatings(limit, offset)

	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		LastPage:  int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}

	if err != nil {
		// fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseAll{
		Success:  true,
		Message:  "get list all ProductRatings successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func GetDetailProductRating(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	ProductRating, err := models.FindOneProductRating(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "ProductRating not found",
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
		Message: "get detail ProductRating successfully",
		Results: ProductRating,
	})
}

func CreateProductRating(c *gin.Context) {
	data := models.ProductRatings{}

	err := c.ShouldBind(&data)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	ProductRating, err := models.CreateProductRating(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "create ProductRating successfully",
		Results: ProductRating,
	})
}

func UpdateProductRating(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.ProductRatings{}

	c.Bind(&data)
	data.Id = id

	ProductRating, err := models.UpdateProductRating(data)

	if err != nil {
		// log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "update ProductRating successfully",
		Results: ProductRating,
	})
}

func DeleteProductRating(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	ProductRating, err := models.DeleteProductRating(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "ProductRating not found",
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
		Message: "delete ProductRating successfully",
		Results: ProductRating,
	})
}
