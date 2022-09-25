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

type OrderUpdate struct {
	CustomerName string       `json:"customerName"`
	OrderedAt    time.Time    `json:"orderedAt"`
	Item         []ItemUpdate `json:"items"`
}

type ItemUpdate struct {
	LineItemID  int    `json:"lineItemId"`
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
		"success": true,
		"message": "success",
		"data":    newOrder,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := db.GetDB()
	id := ctx.Param("orderId")
	var newOrder OrderUpdate

	if err := ctx.BindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateOrder := `UPDATE orders SET customer_name = $2, ordered_at = $3 WHERE order_id = $1;`
	_, err := db.Exec(updateOrder, id, newOrder.CustomerName, newOrder.OrderedAt)
	if err != nil {
		panic(err)
	}

	for _, item := range newOrder.Item {
		updateItem := `UPDATE items SET item_code = $3, description = $4, quantity = $5 WHERE item_id = $1 AND order_id = $2;`
		_, errItem := db.Exec(updateItem, item.LineItemID, id, item.ItemCode, item.Description, item.Quantity)
		if errItem != nil {
			panic(errItem)
		}
	}

	defer db.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    newOrder,
	})
}

func DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("orderId")
	db := db.GetDB()

	deleteOrder := `DELETE from orders WHERE order_id = $1;`
	_, err := db.Exec(deleteOrder, id)
	if err != nil {
		panic(err)
	}

	deleteItem := `DELETE from items WHERE order_id = $1;`
	_, err2 := db.Exec(deleteItem, id)
	if err != nil {
		panic(err2)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
	})
}
