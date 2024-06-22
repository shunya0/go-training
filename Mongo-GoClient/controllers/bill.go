package controllers

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BillCreation(c *gin.Context) {
	var json models.CreateBill
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("error: json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	customer_obj_id, err := primitive.ObjectIDFromHex(json.CustomerId)
	if err != nil {
		fmt.Println("error > customer object id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
	}

	order_details, err := services.GetOrderDetialsByCustomerIDService(customer_obj_id)
	if err != nil {
		fmt.Println("error > order_details: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	customer_details, err := services.GetSingleCustomerService(json.CustomerId)
	if err != nil {
		fmt.Println("error > customer_details: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	var bill_details models.BillGen

	bill_details.CustomerID, err = primitive.ObjectIDFromHex(json.CustomerId)
	if err != nil {
		fmt.Println("error > bill_details_customer_id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	bill_details.OrderID, err = primitive.ObjectIDFromHex(order_details[0].ID)
	if err != nil {
		fmt.Println("error > bill_details_order_id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	bill_details.ShippingStatus = order_details[0].Status
	bill_details.Address.City = customer_details[0].Address.City
	bill_details.Address.Zip = customer_details[0].Address.Zip

	var items []models.ItemOrdered
	for _, product := range order_details[0].Items {
		var item models.ItemOrdered
		item.ProductId = product.ProductId
		item.Price = int(product.Price)
		item.DiscountId = product.DiscountId
		item.Quantity = product.Quantity
		for _, request_product := range json.Products {
			if product.ProductId.Hex() == request_product.ProductId {
				item.Quantity = int(request_product.Quantity)
				break
			}
		}
		items = append(items, item)
	}

	bill_details.Items = items

	bill_creation, err := services.CreateBill(customer_obj_id, bill_details)
	if err != nil {
		fmt.Println("error > bill_creation: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
	fmt.Fprintln(c.Writer, "Bill created sucessfully\nBill Id: ", bill_creation)

}

func GetBill(c *gin.Context) {
	var json models.GetBill
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("error: json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	bill_id_obj, err := primitive.ObjectIDFromHex(json.BillId)
	if err != nil {
		fmt.Println("error > bill obj id\nERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	bill_details, errs := services.GetBill(bill_id_obj)
	if errs != nil {
		fmt.Println("error > bill details\nERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	fmt.Fprintln(c.Writer, "**BILL DETAILS**\n")
	// fmt.Fprintln(c.Writer, "Bill Id: ", bill_details[0].)
	fmt.Fprintln(c.Writer, "Order Id: ", bill_details[0].OrderID)
	fmt.Fprintln(c.Writer, "Customer Id: ", bill_details[0].CustomerID)
	fmt.Fprintln(c.Writer, "Shipping Status: ", bill_details[0].ShippingStatus)
	fmt.Fprintln(c.Writer, "Address: ", bill_details[0].Address)
	fmt.Fprintln(c.Writer, "Orders: [")
	for _, product := range bill_details[0].Items {
		fmt.Fprintln(c.Writer, "{")
		fmt.Fprintln(c.Writer, "	Product Id: ", product.ProductId)
		fmt.Fprintln(c.Writer, "	Quantity: ", product.Quantity)
		fmt.Fprintln(c.Writer, "	Price: ", product.Price)
		fmt.Fprintln(c.Writer, "	Discount Id: ", product.DiscountId)
		fmt.Fprintln(c.Writer, "}\n")
	}
	fmt.Fprintln(c.Writer, "]")

}

func GetBillCustomerId(c *gin.Context) {
	var json models.GetBillByCustomer
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json: ", json.CustomerID)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	customer_id_obj, err := primitive.ObjectIDFromHex(json.CustomerID)
	if err != nil {
		fmt.Println("error > bill obj id\nERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	bill_details, errs := services.GetBillByCustomerId(customer_id_obj)
	if errs != nil {
		fmt.Println("error > bill details\nERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	fmt.Fprintln(c.Writer, "**BILL DETAILS**\n")
	// fmt.Fprintln(c.Writer, "Bill Id: ", bill_details[0].)
	fmt.Fprintln(c.Writer, "Order Id: ", bill_details[0].OrderID)
	fmt.Fprintln(c.Writer, "Customer Id: ", bill_details[0].CustomerID)
	fmt.Fprintln(c.Writer, "Shipping Status: ", bill_details[0].ShippingStatus)
	fmt.Fprintln(c.Writer, "Address: ", bill_details[0].Address)
	fmt.Fprintln(c.Writer, "Orders: [")
	for _, product := range bill_details[0].Items {
		fmt.Fprintln(c.Writer, "{")
		fmt.Fprintln(c.Writer, "	Product Id: ", product.ProductId)
		fmt.Fprintln(c.Writer, "	Quantity: ", product.Quantity)
		fmt.Fprintln(c.Writer, "	Price: ", product.Price)
		fmt.Fprintln(c.Writer, "	Discount Id: ", product.DiscountId)
		fmt.Fprintln(c.Writer, "}\n")
	}
	fmt.Fprintln(c.Writer, "]")

}
