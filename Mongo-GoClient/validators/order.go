package validators

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/services"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CustomerID validator.Func = func(fl validator.FieldLevel) bool {
	customer_id, ok := fl.Field().Interface().(string)

	var customers []models.Customer

	customers, err := services.GetSingleCustomerService(customer_id)
	if err != nil || len(customers) == 0 {

		return !ok
	}
	if !ok {
		return false
	}
	return true
}

var Products validator.Func = func(fl validator.FieldLevel) bool {
	var arr_product_id []string
	field_arr := fl.Field()
	if field_arr.Kind() == reflect.Slice {
		for i := 0; i < field_arr.Len(); i++ {

			product_struct := field_arr.Index(i).Interface().(models.CreateOrderProductBody)
			arr_product_id = append(arr_product_id, product_struct.ProductId)
		}

	} else {
		return false
	}

	product_details, err := services.GetProductsService(arr_product_id)
	if err != nil || len(product_details) != len(arr_product_id) {
		return false
	}

	return true
}

var OrderID validator.Func = func(fl validator.FieldLevel) bool {
	order_id_str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	order_id_obj, err := primitive.ObjectIDFromHex(order_id_str)
	if err != nil {
		return false
	}
	order_details, err := services.GetOrderDetialsService(order_id_obj)
	if err != nil || len(order_details[0].ID) == 0 {
		return false
	}
	return true
}

var ShippingID validator.Func = func(fl validator.FieldLevel) bool {
	shipping_id_str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	shipping_id_obj, err := primitive.ObjectIDFromHex(shipping_id_str)
	if err != nil {
		return false
	}

	shipping_details, err := services.GetShippingDetailsService(shipping_id_obj)
	// fmt.Println(shipping_details, "shippingID\n")
	if err != nil || len(shipping_details[0].ID) == 0 {
		fmt.Println("error shiipingDetails", err)
		return false
	}
	return true
}
