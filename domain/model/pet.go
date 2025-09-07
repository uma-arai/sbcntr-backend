package model

import (
	"time"
)

// Pet ... entity for pet domain model
type Pet struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	Breed            string     `json:"breed"`
	Gender           string     `json:"gender"`
	Price            float64    `json:"price"`
	ImageURL         *string    `json:"image_url"`
	Likes            int        `json:"likes"`
	Shop             Shop       `json:"shop"`
	BirthDate        *time.Time `json:"birth_date"`
	ReferenceNumber  string     `json:"reference_number"`
	Tags             []string   `json:"tags"`
	ReservationCount int64      `json:"reservation_count"`
}

type Shop struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}
