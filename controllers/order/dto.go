package order

import (
	"assignment2/models"
)

//CREATE ORDER

type RequestCreateOrder struct {
	CustomerName string        `json:"customer_name" example:"nathan fernando"`
	Items        []models.Item `json:"items"`
}

type ResponseCreateOrder struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

// GET ALL ORDERS

type ResponseGetOrders struct {
	Orders []models.Order `json:"orders"`
}

// GET ORDER BY ID

type ResponseGetOrder struct {
	Order models.Order `json:"order"`
}

// UPDATE ORDER

type RequestUpdateOrder struct {
	CustomerName string `json:"customer_name" example:"nathan fernando"`
	//Items        []models.Item `json:"items"`
}

type ResponseUpdateOrder struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

// DELETE ORDER

type RequestDeleteOrder struct {
	ID string `json:"id" example:"51asd-asd-as14"`
}

type ResponseDeleteOrder struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}
