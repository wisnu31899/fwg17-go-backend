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

func GetAllProductTag(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit
	result, err := models.FindAllProductTag(limit, offset)

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
		Message:  "get list all ProductTags successfully",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func GetDetailProductTag(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	ProductTag, err := models.FindOneProductTag(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "ProductTag not found",
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
		Message: "get detail ProductTag successfully",
		Results: ProductTag,
	})
}

func CreateProductTag(c *gin.Context) {
	data := models.ProductTag{}

	err := c.ShouldBind(&data)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	ProductTag, err := models.CreateProductTag(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "create ProductTag successfully",
		Results: ProductTag,
	})
}

func UpdateProductTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.ProductTag{}

	c.Bind(&data)
	data.Id = id

	ProductTag, err := models.UpdateProductTag(data)

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
		Message: "update ProductTag successfully",
		Results: ProductTag,
	})
}

func DeleteProductTag(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	ProductTag, err := models.DeleteProductTag(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "ProductTag not found",
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
		Message: "delete ProductTag successfully",
		Results: ProductTag,
	})
}
