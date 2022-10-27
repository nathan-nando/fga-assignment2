package order

import (
	"assignment2/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// CreateOrder godoc
// @Router /order [post]
// @Summary Create new Order
// @Accept json
// @Produce json
// @Param payload body RequestCreateOrder true " "
// @Success 201 {object} ResponseCreateOrder "Success"
func CreateOrder(c *gin.Context) {
	db := models.GetDB()

	var req RequestCreateOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	orderId := uuid.NewString()
	order := models.Order{
		ID:           orderId,
		CustomerName: req.CustomerName,
		Items:        req.Items,
		OrderedAt:    time.Now(),
	}
	db.Create(&order)
	c.JSON(http.StatusOK, ResponseCreateOrder{
		ID:      orderId,
		Message: "successfully created order",
	})
}

// GetOrders godoc
// @Router /orders [get]
// @Summary Get all orders
// @Accept json
// @Produce json
// @Success 200 {object} ResponseGetOrders "Success"
func GetOrders(c *gin.Context) {
	db := models.GetDB()

	var orders []models.Order
	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, ResponseGetOrders{
		Orders: orders,
	})
}

// GetOrder godoc
// @Router /order/{id} [get]
// @Summary Get order by ID
// @Accept json
// @Param id path string true "id"
// @Produce json
// @Success 200 {object} ResponseGetOrder "Success"
func GetOrder(c *gin.Context) {
	db := models.GetDB()

	var order models.Order
	if err := db.Where("id=?", c.Param("id")).Preload("Items").First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found with id : " + c.Param("id"),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseGetOrder{
		Order: order,
	})
}

// UpdateOrder godoc
// @Router /order/{id} [put]
// @Summary Update Order
// @Accept json
// @Param id path string true "id"
// @Produce json
// @Param payload body RequestUpdateOrder true " "
// @Success 200 {object} ResponseUpdateOrder "Success"
func UpdateOrder(c *gin.Context) {
	db := models.GetDB()

	var order models.Order

	if err := db.Where("id=?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found with id : " + c.Param("id"),
		})
		return
	}

	var req RequestUpdateOrder

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RequestUpdateOrder{
			CustomerName: req.CustomerName,
			//Items:        req.Items,
		})
		return
	}

	if err := db.Model(&order).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ResponseDeleteOrder{
			Message: "failed updated",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseUpdateOrder{
		Message: "successfully updated",
		ID:      c.Param("id"),
	})
}

// DeleteOrder godoc
// @Router /order/{id} [delete]
// @Summary Delete Order
// @Accept json
// @Param id path string true "id"
// @Produce json
// @Success 200 {object} ResponseDeleteOrder "Success"
func DeleteOrder(c *gin.Context) {
	db := models.GetDB()

	var order models.Order
	if err := db.Where("id=?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found with id : " + c.Param("id"),
		})
		return
	}

	if err := db.Select("Items").Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ResponseDeleteOrder{
			Message: "failed deleted data",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseDeleteOrder{
		Message: "successfully deleted data",
		ID:      c.Param("id"),
	})
}
