package main

import (
	"Mongo-GoClient/controllers"
	"Mongo-GoClient/database"
	"Mongo-GoClient/validators"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	defer database.Client.Disconnect(context.Background())
	r := gin.Default()
	// r.Use(controllers.SessionMiddleware)
	if customerValid, ok := binding.Validator.Engine().(*validator.Validate); ok {
		customerValid.RegisterValidation("customer_id", validators.CustomerID)
	}

	if productIdValid, ok := binding.Validator.Engine().(*validator.Validate); ok {
		productIdValid.RegisterValidation("product_id", validators.Products)
	}

	if orderIdValid, ok := binding.Validator.Engine().(*validator.Validate); ok {
		orderIdValid.RegisterValidation("order_id", validators.OrderID)
	}

	if shippingtIdValid, ok := binding.Validator.Engine().(*validator.Validate); ok {
		shippingtIdValid.RegisterValidation("shipping_id", validators.ShippingID)
	}

	if custoemrIdValidBill, ok := binding.Validator.Engine().(*validator.Validate); ok {
		custoemrIdValidBill.RegisterValidation("customer_id", validators.CustomerIDForBill)
	}

	if BillID, ok := binding.Validator.Engine().(*validator.Validate); ok {
		BillID.RegisterValidation("bill_id", validators.BillID)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/orders", controllers.SessionMiddleware, controllers.GetUserOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.DELETE("/orders", controllers.CancelOrder)

	r.GET("/bill", controllers.GetBillCustomerId)
	r.POST("/bill", controllers.BillCreation)
	r.Run()

}

//r.Use(gin.Logger()) for writing logs to gin.DefaulWriter
//r.Use(gin.Recovery()) Recovery middleware recovers from any panics and writes a 500 if there was one.
