package validators

import (
	"Mongo-GoClient/services"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CustomerIDForBill validator.Func = func(fl validator.FieldLevel) bool {
	customer_id, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	customer_details, err := services.GetSingleCustomerService(customer_id)
	if err != nil || len(customer_details) == 0 {
		return false
	}

	if string(customer_details[0].ID) != customer_id {

		return false
	}

	return true
}

var BillID validator.Func = func(fl validator.FieldLevel) bool {
	bill_id, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	bill_id_obj, err := primitive.ObjectIDFromHex(bill_id)
	if err != nil {
		return false
	}
	bill_details, err := services.GetBill(bill_id_obj)
	if err != nil || len(bill_details[0].CustomerID) == 0 {
		return false
	}
	return true

}
