package modelstest

import (
	"Mongo-GoClient/models"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func checkUserAddress(userAddress *models.UserAddress) error {
	if len(userAddress.City) == 0 || len(userAddress.Zip) == 0 || len(userAddress.Zip) != 6 || userAddress.City == "" {
		return errors.New("Invalid Address")
	}
	return nil
}

func checkCreatBill(t *models.CreateBill) error {
	if len(t.CustomerId) == 0 {
		return errors.New("Invalid Bill Details")
	}
	if len(t.Products) == 0 {
		return errors.New("Invalid Bill Details")
	}
	for _, product := range t.Products {
		if err := checkProductDetailsBill(&product); err != nil {
			return err
		}
	}
	return nil
}

func checkGetBill(t *models.GetBill) error {
	if t == nil {
		return errors.New("Invalid Bill")
	}
	if t.BillId == primitive.NilObjectID.Hex() {
		return errors.New("Invalid Bill")
	}
	_, err := primitive.ObjectIDFromHex(t.BillId)
	if err != nil {
		return errors.New("Invalid Bill")
	}
	return nil
}

func checkGetBillByCustomer(t *models.GetBillByCustomer) error {
	if t == nil {
		return errors.New("Invalid Bill")
	}
	if t.CustomerID == primitive.NilObjectID.Hex() {
		return errors.New("Invalid Bill")
	}
	_, err := primitive.ObjectIDFromHex(t.CustomerID)
	if err != nil {
		return errors.New("Invalid Bill")
	}
	return nil
}

func checkProductDetailsBill(t *models.ProductDetailsBill) error {
	if len(t.ProductId) == 0 {
		return errors.New("Invalid Bill Details")
	}
	if t.Quantity <= 0 {
		return errors.New("Invalid Bill Details")
	}
	return nil
}

func checkbillgen(t *models.BillGen) error {
	if t.OrderID == primitive.NilObjectID || t.CustomerID == primitive.NilObjectID {
		return errors.New("Invalid BillGen")
	}
	if len(t.ShippingStatus) == 0 {
		return errors.New("Invalid BillGen")
	}
	if err := checkUserAddress(&t.Address); err != nil {
		return err
	}
	for _, item := range t.Items {
		if err := checkItemOrdered(&item); err != nil {
			return err
		}
	}
	return nil
}

func checkItemOrdered(t *models.ItemOrdered) error {
	if t.ProductId != primitive.NilObjectID && t.Quantity > 0 && t.Price > 0 && t.DiscountId != primitive.NilObjectID {
		return nil
	}
	return errors.New("Invalid Item")
}

func Test_UserAddress(t *testing.T) {
	tests := []struct {
		name         string
		user_address *models.UserAddress
		expected     error
	}{
		{"valid address", &models.UserAddress{City: "New york", Zip: "310314"}, nil},
		{"empty address", &models.UserAddress{City: "", Zip: ""}, errors.New("Invalid Address")},
		{"Invalid address", &models.UserAddress{City: "New York", Zip: "1234"}, errors.New("Invalid Address")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkUserAddress(tc.user_address)
			if tc.expected == nil {
				if err != nil {
					t.Errorf("expected no error, got %q", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					t.Errorf("expected error %q, got %q", tc.expected, err)
				}
			}
		})
	}
}

func Test_ValidItemOrdered(t *testing.T) {
	tests := []struct {
		name     string
		item     *models.ItemOrdered
		expected error
	}{
		{"valid item ordered", &models.ItemOrdered{ProductId: primitive.NewObjectID(), Quantity: 1, Price: 1000, DiscountId: primitive.NewObjectID()}, nil},
		{"empty item ordered", &models.ItemOrdered{ProductId: primitive.NilObjectID, Quantity: 0, Price: 0, DiscountId: primitive.NilObjectID}, errors.New("Invalid Item")},
		{"invalid item ordered", &models.ItemOrdered{ProductId: primitive.NilObjectID, Quantity: 21, Price: 1321012, DiscountId: primitive.NilObjectID}, errors.New("Invalid Item")},
		{"invalid item ordered", &models.ItemOrdered{ProductId: primitive.NilObjectID, Quantity: 21, Price: 1321012, DiscountId: primitive.NilObjectID}, errors.New("Invalid Item")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkItemOrdered(tc.item)
			if tc.expected == nil {
				if err != nil {
					t.Errorf("expected no error, got %q", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					t.Errorf("expected error %q, got %q", tc.expected, err)
				}
			}
		})
	}
}

func Test_BillGen(t *testing.T) {
	tests := []struct {
		name     string
		bill     *models.BillGen
		expected error
	}{
		{"valid billgen", &models.BillGen{
			OrderID:        primitive.NewObjectID(),
			CustomerID:     primitive.NewObjectID(),
			ShippingStatus: "shipped",
			Address: models.UserAddress{
				City: "New York",
				Zip:  "310314",
			},
			Items: []models.ItemOrdered{
				{
					ProductId:  primitive.NewObjectID(),
					Quantity:   1,
					Price:      1000,
					DiscountId: primitive.NewObjectID(),
				},
			},
		}, nil},
		{"invalid billgen", &models.BillGen{
			OrderID:        primitive.NilObjectID,
			CustomerID:     primitive.NilObjectID,
			ShippingStatus: "shipped",
			Address: models.UserAddress{
				City: "",
				Zip:  "102",
			},
			Items: []models.ItemOrdered{
				{
					ProductId:  primitive.NilObjectID,
					Quantity:   1,
					Price:      -12,
					DiscountId: primitive.NilObjectID,
				},
			},
		}, errors.New("Invalid BillGen")},
		{"empty billgen", &models.BillGen{
			OrderID:        primitive.NilObjectID,
			CustomerID:     primitive.NilObjectID,
			ShippingStatus: "shipped",
			Address: models.UserAddress{
				City: "",
				Zip:  "",
			},
			Items: []models.ItemOrdered{
				{
					ProductId:  primitive.NilObjectID,
					Quantity:   0,
					Price:      0,
					DiscountId: primitive.NilObjectID,
				},
			},
		}, errors.New("Invalid BillGen")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkbillgen(tc.bill)
			if tc.expected == nil {
				if err != nil {
					t.Errorf("expected no error, got %q", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					t.Errorf("expected error %q, got %q", tc.expected, err)
				}
			}
		})
	}
}

func Test_CreateBil(t *testing.T) {
	tests := []struct {
		name       string
		createbill *models.CreateBill
		expected   error
	}{
		{"valid bill", &models.CreateBill{CustomerId: primitive.NewObjectID().Hex(), Products: []models.ProductDetailsBill{{ProductId: primitive.NewObjectID().Hex(), Quantity: 100}}}, nil},
		{"invalid bill", &models.CreateBill{CustomerId: primitive.NilObjectID.Hex(), Products: []models.ProductDetailsBill{{ProductId: primitive.NilObjectID.Hex(), Quantity: -10}}}, errors.New("Invalid Bill Details")},
		{"empty bill", &models.CreateBill{CustomerId: primitive.NilObjectID.Hex(), Products: []models.ProductDetailsBill{{ProductId: primitive.NilObjectID.Hex(), Quantity: 0}}}, errors.New("Invalid Bill Details")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkCreatBill(tc.createbill)
			if tc.expected == nil {
				if err != nil {
					t.Errorf("expected no error, got %q", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					t.Errorf("expected error %q, got %q", tc.expected, err)
				}
			}
		})
	}
}

func Test_GetBill(t *testing.T) {

	tests := []struct {
		name     string
		getbill  *models.GetBill
		expected error
	}{
		{"valid getbill", &models.GetBill{BillId: primitive.NewObjectID().Hex()}, nil},
		{"invalid getbill", &models.GetBill{BillId: primitive.NilObjectID.Hex()}, errors.New("Invalid Bill")},
		{"nil getbill", nil, errors.New("Invalid Bill")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkGetBill(tc.getbill)

			if tc.expected == nil {
				if err != nil {
					t.Errorf("expected no error, got %q", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					t.Errorf("expected error %q, got %q", tc.expected, err)
				}
			}
		})
	}
}
func Test_GetBillByCustomer(t *testing.T) {

	tests := []struct {
		name     string
		getbill  *models.GetBillByCustomer
		expected error
	}{
		{"valid getbill", &models.GetBillByCustomer{CustomerID: primitive.NewObjectID().Hex()}, nil},
		{"invalid getbill", &models.GetBillByCustomer{CustomerID: primitive.NilObjectID.Hex()}, errors.New("Invalid Bill")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkGetBillByCustomer(tc.getbill)
			if tc.expected == nil {
				if err != nil {
					t.Errorf("expected no error, got %q", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					t.Errorf("expected error %q, got %q", tc.expected, err)
				}
			}
		})
	}
}
