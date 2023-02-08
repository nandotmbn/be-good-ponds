package views

import "time"

type PayloadRecord struct {
	Acidity     float32 `json:"acidity,omitempty"`
	Salinity    float32 `json:"salinity,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	Oxygen      float32 `json:"oxygen,omitempty"`
}

type LastRecord struct {
	Acidity     float32   `json:"acidity,omitempty"`
	Salinity    float32   `json:"salinity,omitempty"`
	Temperature float32   `json:"temperature,omitempty"`
	Oxygen      float32   `json:"oxygen,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type FinalRecord struct {
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Acidity     float32     `json:"acidity,omitempty"`
	Salinity    float32     `json:"salinity,omitempty"`
	Temperature float32     `json:"temperature,omitempty"`
	Oxygen      float32     `json:"oxygen,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
