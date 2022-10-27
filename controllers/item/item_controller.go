package item

import (
	"assignment2/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateItem(c *gin.Context) {
	db := models.GetDB()

	var req RequestCreateItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	itemId := uuid.NewString()
	order := models.Item{
		ID:          itemId,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
	}
	db.Create(&order)
	c.JSON(http.StatusOK, ResponseCreateItem{
		ID:      itemId,
		message: "successfully created item",
	})
}
