package handlers

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/services"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createBill struct {
	CustomerId string               `json:"customer_id" bson:"customer_id"`
	Products   []productDetailsBill `json:"products"`
}
type productDetailsBill struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func BillCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post request allowed", http.StatusBadRequest)
		return
	}

	var bill_body_data createBill

	if err := json.NewDecoder(r.Body).Decode(&bill_body_data); err != nil {
		http.Error(w, "error in reading request body", http.StatusBadRequest)
		return
	}

	if bill_body_data.CustomerId == "" {
		http.Error(w, "incorrect request body", http.StatusBadRequest)
		return
	}

	customer_obj_id, err := primitive.ObjectIDFromHex(bill_body_data.CustomerId)
	if err != nil {
		http.Error(w, "not a valid user", http.StatusBadRequest)
		return
	}
	orderDetails, err := services.GetOrderDetialsService(customer_obj_id) // order details
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	customerDetails, err := services.GetSingleCustomerService(bill_body_data.CustomerId) //customer details

	var bill_details models.BillGen

	bill_details.CustomerID, err = primitive.ObjectIDFromHex(bill_body_data.CustomerId)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	bill_details.OrderID, err = primitive.ObjectIDFromHex(orderDetails[0].ID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	bill_details.ShippingStatus = orderDetails[0].Status
	bill_details.Address.City = customerDetails[0].Address.City
	bill_details.Address.Zip = customerDetails[0].Address.Zip
	var items []models.ItemOrdered
	for _, product := range orderDetails[0].Items {
		var item models.ItemOrdered
		item.ProductId = product.ProductId
		item.Price = int(product.Price)
		item.DiscountId = product.DiscountId
		item.Quantity = product.Quantity
		for _, request_product := range bill_body_data.Products {
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
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Bill created sucessfully\nBill Id: ", bill_creation)

}
