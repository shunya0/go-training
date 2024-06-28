package models

type Discount struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string `json:"description" bson:"description"`
	Percentage  int    `json:"percentage" bson:"percentage"`
}
