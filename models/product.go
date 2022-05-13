package models

import "time"

//TODO: ADD MIN, MAX (ALL STRINGS,QUANTITY,PRICE)
type Product struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Name        string `json:"name" validate:"min=4, max=30"`
	Description string `json:"description,omitempty" validate:"min=4, max=30"`
	// Variation       string    `json:"variation,omitempty", validate:"min=4, max=30"`
	Stock           uint      `json:"stock"`
	Price           float64   `json:"price"`
	Discount        float64   `json:"discount"`
	DiscountEndDate time.Time `json:"discountEndDate"`
	//Location        string    `json:"location,omitempty"`(location id)
	Image string `json:"image,omitempty"`
	//Code			string
}

//Omitempty:Daca fieldu e gol nu apare in baza de date
