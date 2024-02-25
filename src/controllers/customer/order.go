package customer

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func GetDetailOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	Order, err := models.FindOneOrder(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Order not found",
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
		Message: "get detail Order successfully",
		Results: Order,
	})
}

func CreateOrder(c *gin.Context) {
	data := models.Orders{}

	err := c.ShouldBind(&data)
	fmt.Println("2")
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	userId, err := models.FindOneUser(data.UserId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			fmt.Println("3")
			fmt.Println(err)
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "user not found",
			})
			return
		}
		fmt.Println("4")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	// Mengisi data orders jika kosong
	if data.DeliveryAddress == "" && userId.Address != nil {
		data.DeliveryAddress = *userId.Address
	}

	if data.FullName == "" && userId.FullName != nil {
		data.FullName = *userId.FullName
	}

	if data.Email == "" {
		data.Email = userId.Email
	}

	// Menghitung TaxAmount
	if data.Total != 0 {
		tax := int(float64(data.Total) * 0.1)
		data.TaxAmount = tax
	} else {
		fmt.Println("5")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	data.Status = "ON-PROCCESS"

	lastOrderById, err := models.GetLastOrderId()
	if err != nil {
		fmt.Println("6")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}
	// Membuat nomor pesanan baru
	currentTime := time.Now()
	orderNumber := fmt.Sprintf("#%d%d%d-%d", currentTime.Day(), currentTime.Month(), currentTime.Year(), lastOrderById+1)

	data.OrderNumber = orderNumber

	createorder, err := models.CreateOrder(data)

	if err != nil {
		fmt.Println("7")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "create user successfully",
		Results: createorder,
	})
}
