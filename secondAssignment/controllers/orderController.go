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
	Item         Item
}

type Item struct {
	ItemID      int       `json:"itemId"`
	OrderID     int       `json:"orderId"`
	ItemCode    int       `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderedAt   time.Time `json:"orderedAt"`
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
	insertItem := `INSERT INTO items(order_id, item_code, description, quantity, ordered_at) VALUES ($1, $2, $3, $4, $5) Returning *`

	err := db.QueryRow(insertOrder, newOrder.CustomerName, newOrder.OrderedAt).Scan(&newOrder.OrderID, &newOrder.CustomerName, &newOrder.OrderedAt)
	err2 := db.QueryRow(insertItem, newOrder.OrderID, "123", "tes deskripsi", 98, newOrder.OrderedAt).Scan(&newItem.ItemID, &newItem.OrderID, &newItem.ItemCode, &newItem.Description, &newItem.Quantity, &newItem.OrderedAt)

	if err != nil {
		panic(err)
	}

	if err2 != nil {
		panic(err)
	}
	defer db.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    newOrder,
	})
}
