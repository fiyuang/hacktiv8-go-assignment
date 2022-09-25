package controllers

import (
	"net/http"
	"secondAssignment/db"
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderID      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Item    `json:"items"`
}

type Item struct {
	ItemID      int    `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func CreateOrder(ctx *gin.Context) {
	db := db.GetDB()
	var newOrder Order
	var newItem Item

	if err := ctx.BindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertOrder := `INSERT INTO orders(customer_name, ordered_at) VALUES ($1, $2) Returning *`
	err := db.QueryRow(insertOrder, newOrder.CustomerName, newOrder.OrderedAt).Scan(&newOrder.OrderID, &newOrder.CustomerName, &newOrder.OrderedAt)
	if err != nil {
		panic(err)
	}

	for _, item := range newOrder.Item {
		insertItem := `INSERT INTO items(order_id, item_code, description, quantity) VALUES ($1, $2, $3, $4) Returning *`
		errItem := db.QueryRow(insertItem, newOrder.OrderID, item.ItemCode, item.Description, item.Quantity).Scan(&newItem.ItemID, &newOrder.OrderID, &newItem.ItemCode, &newItem.Description, &newItem.Quantity)
		if errItem != nil {
			panic(errItem)
		}
	}

	defer db.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    newOrder,
	})
}
