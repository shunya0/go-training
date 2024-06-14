package handlers

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"Mongo-GoClient/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type createOrderProductBody struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type createOrderRequestBody struct {
	CustomerId string                   `json:"customer_id"`
	Products   []createOrderProductBody `json:"products"`
}

const ORDER_COLLECTION string = "orders"
const DISCOUNT_COLLECTION string = "discount"
const CUSTOMERS_COLLECTION string = "customers"
const SHIPPING_COLLECTION string = "shipping"
const PRODUCTS_COLLECTION string = "products"
const CUSTOMER_COLLECTION string = "customers"

func getOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	col, err := database.GetCollection("orders")
	if err != nil {
		http.Error(w, "failed to get collection (main.go/getOrders) ", http.StatusInternalServerError)
		return
	}
	fmt.Println(database.Client)

	fmt.Println(col)
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Err", err)
		http.Error(w, "Error finding collection (main.go/getOrders)", http.StatusInternalServerError)

		return
	}

	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error decoding order (main.go/getOrders) ", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)

	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Error in iterating cursor (main.go/getOrders) ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, "Error encoding orders (main.go/getOrders)", http.StatusInternalServerError)
		return

	}
}
func GetOrderById(w http.ResponseWriter, r *http.Request) {
	order_id_str := r.URL.Query().Get("orderid")
	order_id, err := primitive.ObjectIDFromHex(order_id_str)
	if err != nil {
		http.Error(w, "Error converting string to objectId (order.go)", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	ordersCol, err := database.GetCollection(ORDER_COLLECTION)
	if err != nil {
		http.Error(w, "Error getting collection: (order.go/orders)", http.StatusInternalServerError)
		return
	}

	pipeline := mongo.Pipeline{

		{
			bson.E{"$match", bson.D{{"_id", order_id}}},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "customers"}, {"localField", "customer_id"}, {"foreignField", "_id"}, {"as", "customer_details"}}},
		},
		{
			bson.E{"$unwind", "$customer_details"},
		},
		{
			bson.E{"$unwind", "$items"},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "products"}, {"localField", "items.product_id"}, {"foreignField", "_id"}, {"as", "product_details"}}},
		},
		{
			bson.E{"$unwind", "$product_details"},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "discount"}, {"localField", "items.discount_id"}, {"foreignField", "_id"}, {"as", "discount_details"}}},
		},
		{
			bson.E{"$unwind", "$discount_details"},
		},
		{
			bson.E{"$group", bson.D{{"_id", "$_id"}, {"customer_name", bson.D{{"$first", "$customer_details.name"}}},
				{"customer_email", bson.D{{"$first", "$customer_details.email"}}}, {"join_date", bson.D{{"$first", "$customer_details.join_date"}}},
				{"address", bson.D{{"$first", "$customer_details.address"}}}, {"ordered_date", bson.D{{"$first", "$order_date"}}}, {"status", bson.D{{"$first", "$status"}}},
				{"order_details", bson.D{{"$push", bson.D{{"name", "$product_details.name"}, {"quantity", "$items.quantity"}, {"original_price_per_product", "$items.price"},
					{"discount", bson.D{{"$multiply", bson.A{"$items.price", bson.D{{"$divide", bson.A{"$discount_details.percentage", 100}}}}}}}}}}}}},
		},
	}

	cursor, err := ordersCol.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Fprintln(w, err)
		http.Error(w, "Error aggreagating orders: (order.go/cursor)", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var orders []models.OrderAggergate

	for cursor.Next(ctx) {
		var order models.OrderAggergate
		errS := cursor.Decode(&order)
		if errS != nil {
			http.Error(w, "Error decoding cursor: (order.go/cursor.next)", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "Error encoding orders ,(order.go/encode)", http.StatusInternalServerError)
		return
	}

}

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only post request allowed!", http.StatusBadRequest)
		return
	}

	var body_data createOrderRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body_data); err != nil {
		http.Error(w, "Error in reading request body", http.StatusBadRequest)
		return
	}

	if body_data.CustomerId == "" || len(body_data.Products) == 0 {
		http.Error(w, "Incorrect request body", http.StatusBadRequest)
		return
	}

	customer_obj_id, err := primitive.ObjectIDFromHex(body_data.CustomerId)
	if err != nil {
		http.Error(w, "invalid customer id", http.StatusBadRequest)
		return
	}

	for _, product := range body_data.Products {

		if len(product.ProductId) == 0 || reflect.TypeOf(product.Quantity) != reflect.TypeOf(int(0)) {
			http.Error(w, "Incorrect products found", http.StatusBadRequest)
			return
		}
	}
	customer_obj, err := services.GetSingleCustomerService(body_data.CustomerId)
	fmt.Println(customer_obj)

	fmt.Println(customer_obj[0].Name)
	if err != nil || customer_obj[0].Name == "" {
		http.Error(w, "customer id not found", http.StatusNotFound)
		return
	}

	var product_ids []string

	for _, product := range body_data.Products {
		product_ids = append(product_ids, product.ProductId)
	}

	product_detail_arr, err := services.GetProductsService(product_ids)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Println(product_detail_arr)

	discount_detail_arr, err := services.GetDiscountServices()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Println(discount_detail_arr)

	var items []models.Items
	for _, product := range product_detail_arr {
		var item models.Items
		product_obj_id, err := primitive.ObjectIDFromHex(product.ID)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		item.ProductId = product_obj_id
		item.Price = int(product.Price)
		random_discount_id, err := primitive.ObjectIDFromHex(discount_detail_arr[rand.Intn(len(discount_detail_arr))].ID)
		if err != nil {
			fmt.Println("error in getting random discount id", err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		item.DiscountId = random_discount_id
		for _, request_product := range body_data.Products {
			if product.ID == request_product.ProductId {
				item.Quantity = int(request_product.Quantity)
				break
			}

		}
		items = append(items, item)
	}

	fmt.Println(items)

	shipping_creation_output, err := services.CreateShippingDetailsService()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if len(shipping_creation_output) == 0 {
		fmt.Println("empty return from shipping service")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Println(shipping_creation_output)

	if err := services.UpdateInventoryService(items); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
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
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if err := services.UpdateShippingOrderIdService(shipping_creation_output[0], order_creation_output[0]); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "order created successfully \norder id:", order_creation_output[0], "\nshipping id: ", shipping_creation_output[0])

}
