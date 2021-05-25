package models

import (
	"time"

	"gorm.io/gorm"
)

//DB Products
type Parking struct {
	ID                 uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt          time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt          time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt          gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	ParkingNumber      uint           `json:"parkingnumber", form:"parkingnumber"`
	RegistrationNumber string         `json:"registrationnumber", form:"registrationnumber"`
	Colour             string         `json:"colour", form:"colour"`
	Status             string         `json:"status", form:"status"`
}
type ParkingRequest struct {
	ParkingNumber      uint   `json:"parkingnumber", form:"parkingnumber"`
	RegistrationNumber string `json:"registrationnumber", form:"registrationnumber"`
	Colour             string `json:"colour", form:"colour"`
	Status             string `json:"status", form:"status"`
}

type ParkingStatus struct {
	ParkingNumber      uint   `json:"parkingnumber", form:"parkingnumber"`
	RegistrationNumber string `json:"registrationnumber", form:"registrationnumber"`
	Colour             string `json:"colour", form:"colour"`
	Status             string `json:"status", form:"status"`
}

type SlotParkingResponse struct {
	ParkingNumber uint `json:"parkingnumber", form:"parkingnumber"`
}
type ParkingResponse struct {
	ParkingNumber      uint   `json:"id_category", form:"id_category"`
	RegistrationNumber string `json:"name", form:"name"`
	Colour             string `json:"description", form:"description"`
}
type ParkingResponseAny struct {
	Code    int     `json:"code", form:"code"`
	Message string  `json:"message", form:"message"`
	Status  string  `json:"status", form:"status"`
	Data    Parking `json:"data", form:"data"`
}
type ParkingResponseMany struct {
	Code    int       `json:"code", form:"code"`
	Message string    `json:"message", form:"message"`
	Status  string    `json:"status", form:"status"`
	Data    []Parking `json:"data", form:"data"`
}

type ParkingResponseManyString struct {
	Data []string `json:"data", form:"data"`
}
type ParkingResponseNotif struct {
	Code    int    `json:"code", form:"code"`
	Message string `json:"message", form:"message"`
	Status  string `json:"status", form:"status"`
}
