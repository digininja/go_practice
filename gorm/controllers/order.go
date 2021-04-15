package controllers

import (
	"github.com/digininja/go_practice/gorm/models"
)

func GetOrder(orderID uint) models.Order {
	var order models.Order
	models.DB.First(&order, orderID)
	return order
}

func GetAllOrders() []models.Order {
	var orders []models.Order
	models.DB.Order("created_at").Find(&orders)
	return orders
}

func GetRecentOrders(count int) []models.Order {
	var orders []models.Order
	models.DB.Limit(count).Order("created_at desc").Find(&orders)
	return orders
}

func GetItems(orderID uint) []models.Item {
	var retrievedItems []models.Item
	models.DB.Where("order_id = ?", orderID).Find(&retrievedItems)
	return retrievedItems
}
