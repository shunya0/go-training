// func getCustomers(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("customers")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getCustomers)", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getCustomers)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var customers []models.Customer
// 	for cursor.Next(ctx) {
// 		var customer models.Customer
// 		err := cursor.Decode(&customer)
// 		if err != nil {
// 			http.Error(w, "Error decoding customer (main.go/getCustomers)", http.StatusInternalServerError)
// 			return
// 		}

// 		customers = append(customers, customer)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getCustomers)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(customers)
// 	if err != nil {
// 		http.Error(w, "Error encoding customer (main.go/getCustomers)", http.StatusInternalServerError)
// 	}
// }

// func getDiscounts(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("discount")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getDiscount): ", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getDiscount)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var discounts []models.Discount
// 	for cursor.Next(ctx) {
// 		var discount models.Discount
// 		err := cursor.Decode(&discount)
// 		if err != nil {
// 			http.Error(w, "Error decoding discount (main.go/getDiscount)", http.StatusInternalServerError)
// 			return
// 		}

// 		discounts = append(discounts, discount)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getDiscount)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(discounts)
// 	if err != nil {
// 		http.Error(w, "Error encoding discount (main.go/getDiscount)", http.StatusInternalServerError)
// 	}
// }

// func getWhishlist(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("whishlists")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getWhishlist): ", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getWhishlist)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var whishlists []models.Whishlist
// 	for cursor.Next(ctx) {
// 		var whishlist models.Whishlist
// 		err := cursor.Decode(&whishlist)
// 		if err != nil {
// 			http.Error(w, "Error decoding whishlists (main.go/getWhishlist)", http.StatusInternalServerError)
// 			return
// 		}

// 		whishlists = append(whishlists, whishlist)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getWhishlist)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(whishlists)
// 	if err != nil {
// 		http.Error(w, "Error encoding whishlists (main.go/getWhishlist)", http.StatusInternalServerError)
// 	}
// }

// func getInventory(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("inventory")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getInventory): ", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getInventory)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var inventories []models.Inventory
// 	for cursor.Next(ctx) {
// 		var inventory models.Inventory
// 		err := cursor.Decode(&inventory)
// 		if err != nil {
// 			http.Error(w, "Error decoding inventories (main.go/getInventory)", http.StatusInternalServerError)
// 			return
// 		}

// 		inventories = append(inventories, inventory)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getInventory)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(inventories)
// 	if err != nil {
// 		http.Error(w, "Error encoding inventories (main.go/getInventory)", http.StatusInternalServerError)
// 	}
// }

// func getProducts(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("products")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getProducts): ", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getProducts)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var products []models.Product
// 	for cursor.Next(ctx) {
// 		var product models.Product
// 		err := cursor.Decode(&product)
// 		if err != nil {
// 			http.Error(w, "Error decoding product (main.go/getProducts)", http.StatusInternalServerError)
// 			return
// 		}

// 		products = append(products, product)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getProducts)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(products)
// 	if err != nil {
// 		http.Error(w, "Error encoding product (main.go/getProducts)", http.StatusInternalServerError)
// 	}
// }

// func getReviews(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("reviews")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getReviews): ", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getReviews)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var reviews []models.Review
// 	for cursor.Next(ctx) {
// 		var review models.Review
// 		err := cursor.Decode(&review)
// 		if err != nil {
// 			http.Error(w, "Error decoding reviews (main.go/getReviews)", http.StatusInternalServerError)
// 			return
// 		}

// 		reviews = append(reviews, review)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getReviews)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(reviews)
// 	if err != nil {
// 		http.Error(w, "Error encoding reviews (main.go/getReviews)", http.StatusInternalServerError)
// 	}
// }

// func getShipping(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	col, err := database.GetCollection("shipping")
// 	if err != nil {
// 		http.Error(w, "failed to get collection (main.go/getShipping): ", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(database.Client)

// 	fmt.Println(col)
// 	cursor, err := col.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Error finding collection (main.go/getShipping)", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var shippings []models.Shipping
// 	for cursor.Next(ctx) {
// 		var shipping models.Shipping
// 		err := cursor.Decode(&shipping)
// 		if err != nil {
// 			http.Error(w, "Error decoding shippings (main.go/getShipping)", http.StatusInternalServerError)
// 			return
// 		}

// 		shippings = append(shippings, shipping)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, "Error in iterating cursor (main.go/getShipping)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(shippings)
// 	if err != nil {
// 		http.Error(w, "Error encoding shippings (main.go/getShipping)", http.StatusInternalServerError)
// 	}
// }

// func getAllCommands(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "orders")
// 	fmt.Fprintln(w, "customers")
// 	fmt.Fprintln(w, "whishlist")
// 	fmt.Fprintln(w, "discounts")
// 	fmt.Fprintln(w, "inventory")
// 	fmt.Fprintln(w, "products")
// 	fmt.Fprintln(w, "review")
// 	fmt.Fprintln(w, "shipping")
// }