package controllers

import (
	"restapi2/model"
	"restapi2/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostOrder(c *gin.Context) {
	db := model.GetDb()
	order := model.Order{}

	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	newOrder := model.Order{
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Items:        order.Items,
	}

	if err := db.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	addedItems := make([]response.Item, len(newOrder.Items))

	for i, _ := range newOrder.Items {
		addedItems[i] = response.Item{
			Description: newOrder.Items[i].Description,
			ItemID:      newOrder.Items[i].ID,
			Quantity:    newOrder.Items[i].Quantity,
		}
	}

	addedOrder := response.Order{
		CustomerName: newOrder.CustomerName,
		Items:        addedItems,
		ID:           newOrder.ID,
		OrderedAt:    newOrder.OrderedAt,
	}

	c.JSON(http.StatusOK, response.Response{Message: "add order success", Data: addedOrder})
}

func PutOrder(c *gin.Context) {
	db := model.GetDb()
	order := model.Order{}
	item := model.Item{}

	if err := db.Where("order_id = ?", c.Param("orderId")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	db.Unscoped().Where("order_id = ?", order.ID).Delete(item)

	if err := db.Unscoped().Where("order_id = ?", order.ID).Delete(item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := db.Save(order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	updatedItem := make([]response.Item, len(order.Items))

	for i, _ := range order.Items {
		updatedItem[i] = response.Item{
			Description: order.Items[i].Description,
			ItemID:      order.Items[i].ID,
			Quantity:    order.Items[i].Quantity,
		}
	}

	updatedOrder := response.Order{
		CustomerName: order.CustomerName,
		Items:        updatedItem,
		ID:           order.ID,
		OrderedAt:    order.OrderedAt,
	}

	c.JSON(http.StatusOK, response.Response{Message: "update order success", Data: updatedOrder})
}

func DeleteOrder(c *gin.Context) {
	db := model.GetDb()
	order := model.Order{}

	if err := db.Where("order_id = ?", c.Param("orderId")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	if err := db.Select(clause.Associations).Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete order success"})
}

func GetOrders(c *gin.Context) {
	db := model.GetDb()
	orders := []model.Order{}

	if err := c.ShouldBind(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := db.Model(&model.Order{}).Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ordersData := make([]response.Order, len(orders))

	for i, _ := range orders {
		itemsData := make([]response.Item, len(orders[i].Items))

		for j, _ := range orders[i].Items {
			itemsData[j] = response.Item{
				Description: orders[i].Items[j].Description,
				ItemID:      orders[i].Items[j].ID,
				Quantity:    orders[i].Items[j].Quantity,
			}

			ordersData[i] = response.Order{
				CustomerName: orders[i].CustomerName,
				Items:        itemsData,
				ID:           orders[i].ID,
				OrderedAt:    orders[i].OrderedAt,
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": ordersData})
}