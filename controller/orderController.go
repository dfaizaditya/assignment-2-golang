package controllers

import (
	"assignment-2/database"
	"assignment-2/helpers"
	"assignment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()
	Order := models.Order{}

	if err := c.ShouldBindJSON(&Order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseError{
			Error:   "Bad Request",
			Message: err.Error(),
		})
		return
	}

	err := db.Debug().Create(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseError{
			Error:   "Bad request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Order)

}

func GetOrders(c *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}

	err := db.Preload("Items").Find(&Orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseError{
			Error:   "Bad request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Orders)
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	orderID, err := strconv.Atoi(c.Param("OrderID"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
			Error:   "Internal server error",
			Message: "Invalid param orderId",
		})
		return
	}

	db.Delete(models.Item{}, "order_id", orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
			Error:   "Error deleting item",
			Message: err.Error(),
		})
		return
	}

	err = db.Delete(models.Order{}, "order_id", orderID).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
			Error:   "Error deleting order",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete order",
	})
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDB()
	OrderID, err := strconv.Atoi(c.Param("OrderID"))
	UpdateOrder := models.Order{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
			Error:   "Internal server error",
			Message: "Invalid param orderId",
		})
		return
	}
	if err := c.ShouldBindJSON(&UpdateOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseError{
			Error:   "Bad request",
			Message: err.Error(),
		})
		return
	}

	for i := range UpdateOrder.Items {
		err = db.Model(&UpdateOrder.Items[i]).Where("item_id=?", UpdateOrder.Items[i].ItemID).Updates(&UpdateOrder.Items[i]).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
				Error:   "Error updating item",
				Message: err.Error(),
			})
			return
		}
	}

	err = db.Model(&UpdateOrder).Where("order_id=?", OrderID).Omit("Items").Updates(&UpdateOrder).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
			Error:   "Error updating order",
			Message: err.Error(),
		})
		return
	}

	err = db.Preload("Items").Where("order_id=?", OrderID).Find(&UpdateOrder).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ResponseError{
			Error:   "Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Message: "Successfuly update data",
		Data:    UpdateOrder,
	})
}
