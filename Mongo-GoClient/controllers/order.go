package controllers

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/services"
	"Mongo-GoClient/utils"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserOrders(c *gin.Context) {

	CustomerIdStr := c.Request.URL.Query().Get("customer_id")
	CustomerId, err := primitive.ObjectIDFromHex(CustomerIdStr)
	if err != nil {
		fmt.Println("error customer id obj getUserOrder")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	if CustomerIdStr != utils.CUSTOMER_LOGGED {
		fmt.Println("Not authorized: ", CustomerId, "\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	orders, err := services.GetOrderDetialsByCustomerIDService(CustomerId)
	if err != nil {
		fmt.Println("error getting order detail for user: ", CustomerId, "\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	fmt.Fprintln(c.Writer, "=========ORDER DETAILS=========")
	fmt.Fprintln(c.Writer, orders[0].CustomerId, " : Customer ID")
	fmt.Fprintln(c.Writer, orders[0].ID, " : Order ID")
	fmt.Fprintln(c.Writer, orders[0].Items, " : Items")
	fmt.Fprintln(c.Writer, orders[0].OrderDate, " : Order Date")
	fmt.Fprintln(c.Writer, orders[0].ShippingID, " : Shipping ID")
	fmt.Fprintln(c.Writer, orders[0].Status, " : Status")
	fmt.Fprintln(c.Writer, "===============================")

}

func CreateOrder(c *gin.Context) {
	var json models.CreateOrderRequestBody
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	customer_obj_arr, err := services.GetSingleCustomerService(json.CustomerId)
	if err != nil {
		fmt.Println("customerARR")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	customer_obj_id, err := primitive.ObjectIDFromHex(customer_obj_arr[0].ID)
	if err != nil {
		fmt.Println("customerID")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	var product_id []string

	for _, product := range json.Products {
		product_id = append(product_id, product.ProductId)
	}

	product_detail_arr, err := services.GetProductsService(product_id)
	if err != nil {
		fmt.Println("prdctARR")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	discounts, err := services.GetDiscountServices()
	if err != nil {
		fmt.Println("discount")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	var items []models.Items
	for _, product := range product_detail_arr {
		var item models.Items
		product_obj_id, err := primitive.ObjectIDFromHex(product.ID)
		if err != nil {
			fmt.Println("productObjID")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}
		item.ProductId = product_obj_id
		item.Price = int(product.Price)
		random_discount_id, err := primitive.ObjectIDFromHex(discounts[rand.Intn(len(discounts))].ID)
		if err != nil {
			fmt.Println("randomDiscount")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}
		item.DiscountId = random_discount_id
		for _, request_product := range json.Products {
			if product.ID == request_product.ProductId {
				item.Quantity = int(request_product.Quantity)
				break
			}

		}
		items = append(items, item)
	}

	shipping_creation_output, err := services.CreateShippingDetailsService()
	if err != nil {
		fmt.Println("shippingCreation")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := services.UpdateInventoryService(items); err != nil {
		fmt.Println("updateService")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	var new_order models.Order
	new_order.OrderDate = primitive.NewDateTimeFromTime(time.Now())
	new_order.CustomerId = customer_obj_id
	new_order.Status = "pending"
	new_order.Items = items
	new_order.ShippingID = shipping_creation_output[0]

	order_creation_output, err := services.CreateOrderService(new_order)
	if err != nil {
		fmt.Println("orderCreation")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := services.UpdateShippingOrderIdService(shipping_creation_output[0], order_creation_output[0]); err != nil {
		fmt.Println("updateShippingID")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
	fmt.Fprintln(c.Writer, "order created successfully \norder id:", order_creation_output[0], "\nshipping id: ", shipping_creation_output[0])

}

func CancelOrder(c *gin.Context) {
	var json models.CancelOrderBody
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	customer_obj_id, err := primitive.ObjectIDFromHex(json.CustomerId)
	if err != nil {
		fmt.Println("err converting customerID: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	order_obj_id, err := primitive.ObjectIDFromHex(json.OrderId)
	if err != nil {
		fmt.Println("err converting orderId: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	shipping_obj_id, err := primitive.ObjectIDFromHex(json.ShippingId)
	if err != nil {
		fmt.Println("err converting shippingID: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	shippment_cancel_err := services.CancelShippmentService(order_obj_id)
	if shippment_cancel_err != nil {
		fmt.Println("err in shippment_cancel: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	delete_shipping_err := services.DeleteShippmentService(shipping_obj_id)
	if delete_shipping_err != nil {
		fmt.Println("err in delete_shippment: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	var canceled_order models.Cancel
	canceled_order.CustomerID = customer_obj_id
	canceled_order.OrderID = order_obj_id
	canceled_order.ShippingID = shipping_obj_id
	canceled_order.ShippingStatus = "Canceled!"

	c.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(c.Writer, "User: ", customer_obj_id, " has canceled order with id: ", order_obj_id)

}
