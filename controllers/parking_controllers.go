package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gitlab.mapan.io/playground/parking-lot-golang/configs"
	"gitlab.mapan.io/playground/parking-lot-golang/models"
	"gorm.io/gorm/clause"
)

func CreateSlotParkingController(c echo.Context) error {
	slotreceive := c.Param("n")
	slot, _ := strconv.Atoi(slotreceive)
	var ParkingInput models.ParkingRequest
	c.Bind(&ParkingInput)

	var createslotDB models.Parking

	for i := 1; i <= slot; i++ {

		createslotDB.ParkingNumber = uint(i)
		createslotDB.RegistrationNumber = ""
		createslotDB.Status = "empty"
		configs.DB.Create(&createslotDB)
		createslotDB.ID++
	}
	// err := configs.DB.Create(&createslotDB).Error
	if slot == 0 {
		return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Create SLot",
			Status:  "error",
		})
	}

	return c.JSON(http.StatusOK, models.ParkingResponseNotif{
		Code:    http.StatusOK,
		Message: "Created a parking lot with " + slotreceive + " slots",
		Status:  "success",
	})

}

func AddingVehicleParking(c echo.Context) error {
	nopolreceive := c.Param("nopol")
	colourreceive := c.Param("colour")

	var ParkingInput models.ParkingRequest
	c.Bind(&ParkingInput)

	var AddingVehicleParkingDB models.Parking
	err_row := configs.DB.Where("status = ?", "empty").First(&AddingVehicleParkingDB)
	err_cat := configs.DB.Preload(clause.Associations).Find(&AddingVehicleParkingDB, "status = ?", "empty").Error

	if err_cat != nil || err_row.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_cat.Error(),
			Status:  "Sorry,parking lot is full",
		})
	}

	AddingVehicleParkingDB.RegistrationNumber = nopolreceive
	AddingVehicleParkingDB.Colour = colourreceive
	AddingVehicleParkingDB.Status = "BOOKED"
	err := configs.DB.Save(&AddingVehicleParkingDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "Cannot Update",
		})
	}
	return c.JSON(http.StatusOK, models.ParkingResponseNotif{
		Code:    http.StatusOK,
		Message: "Allocated slot number: " + strconv.Itoa(int(AddingVehicleParkingDB.ParkingNumber)),
		Status:  "Success ",
	})
}

func LeaveParking(c echo.Context) error {
	numberleave := c.Param("parkingnumber")
	numberleaveconv, _ := strconv.Atoi(numberleave)
	var ParkingLeaveInput models.ParkingRequest
	c.Bind(&ParkingLeaveInput)
	var parkingDB models.Parking
	if numberleaveconv != 0 {

		row_park := configs.DB.Where("parking_number = ?", numberleave).Find(&parkingDB)
		// err_park := configs.DB.Preload(clause.Associations).Find(&parkingDB, "status = ?", "BOOKED").Error
		if row_park.RowsAffected != 0 && parkingDB.Status == "BOOKED" {
			parkingDB.Status = "empty"
			parkingDB.Colour = ""
			parkingDB.RegistrationNumber = ""
			configs.DB.Save(&parkingDB)
			return c.JSON(http.StatusOK, models.ParkingResponseNotif{
				Code:    http.StatusOK,
				Message: "Slot number " + strconv.Itoa(int(parkingDB.ParkingNumber)) + " is free ",
				Status:  "parking number " + strconv.Itoa(int(parkingDB.ParkingNumber)) + " is free ",
			})

		}
		return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: "Sorry, parking lot is full",
			Status:  "error",
		})

	}
	return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
		Code:    http.StatusInternalServerError,
		Message: "Nomor Parkir tidak terdeteksi",
		Status:  "error",
	})

}
func Getnopolbycolour(c echo.Context) error { //sek belum
	get_reg := make([]string, 0)
	var count int64
	var i int64
	receive_colour := c.Param("colour")
	var parkingDB []models.Parking
	if receive_colour != "" {

		row_park := configs.DB.Where("colour = ?", receive_colour).Find(&parkingDB).RowsAffected
		err_park := configs.DB.Preload(clause.Associations).Find(&parkingDB, "colour = ?", receive_colour).Error
		configs.DB.Where("colour = ?", receive_colour).Find(&parkingDB).Count(&count)
		if err_park != nil || row_park == 0 {
			return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
				Code:    http.StatusInternalServerError,
				Message: "Data Empty",
				Status:  "error",
			})
		}
		for i = 0; i < count; i++ {

			get_reg = append(get_reg, parkingDB[i].RegistrationNumber)

		}
		return c.JSON(http.StatusOK, models.ParkingResponseManyString{

			Data: get_reg,
		})
	}
	return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
		Code:    http.StatusInternalServerError,
		Message: "Input Not Null",
		Status:  "error",
	})

}

func GetSlotParkingByColour(c echo.Context) error {
	get_slot := make([]string, 0)
	var count int64
	var i int64
	receive_colour := c.Param("colour")
	var parkingDB []models.Parking
	if receive_colour != "" {

		row_park := configs.DB.Where("colour = ?", receive_colour).Find(&parkingDB).RowsAffected
		err_park := configs.DB.Preload(clause.Associations).Find(&parkingDB, "colour = ?", receive_colour).Error
		configs.DB.Where("colour = ?", receive_colour).Find(&parkingDB).Count(&count)
		if err_park != nil || row_park == 0 {
			return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
				Code:    http.StatusInternalServerError,
				Message: " Not Found",
				Status:  "error",
			})
		}
		for i = 0; i < count; i++ {

			get_slot = append(get_slot, strconv.Itoa(int(parkingDB[i].ParkingNumber)))

		}
		return c.JSON(http.StatusOK, models.ParkingResponseManyString{

			Data: get_slot,
		})
	}
	return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
		Code:    http.StatusInternalServerError,
		Message: "Input Not Null",
		Status:  "error",
	})

}

func GetSlotParkingByRegistrationNumber(c echo.Context) error { //return parking number
	get_parking_number := make([]string, 0)
	var count int64
	var i int64
	receive_regnumber := c.Param("reg")
	var parkingDB []models.Parking
	if receive_regnumber != "" {

		row_park := configs.DB.Where("registration_number = ?", receive_regnumber).Find(&parkingDB).RowsAffected
		configs.DB.Where("registration_number = ?", receive_regnumber).Find(&parkingDB).Count(&count)
		err_park := configs.DB.Preload(clause.Associations).Find(&parkingDB, "registration_number = ?", receive_regnumber).Error
		if err_park != nil || row_park == 0 {
			return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
				Code:    http.StatusInternalServerError,
				Message: "Not Found",
				Status:  "error",
			})
		}
		for i = 0; i < count; i++ {

			get_parking_number = append(get_parking_number, strconv.Itoa(int(parkingDB[i].ParkingNumber)))

		}
		// return c.JSON(http.StatusOK, models.ParkingResponseManyString{
		// 	Code:    http.StatusOK,
		// 	Message: "Success get parking number by registration number",
		// 	Status:  "success",
		// 	Data:    parkingDB,
		// })
		return c.JSON(http.StatusOK, models.ParkingResponseManyString{
			Data: get_parking_number,
		})
	}
	return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
		Code:    http.StatusInternalServerError,
		Message: "Input Not Null",
		Status:  "error",
	})

}

func GetStatus(c echo.Context) error {
	var allparkingDB []models.Parking

	var count int64
	var i int64
	get_parking_number, get_registration, get_colour := make([]string, 0), make([]string, 0), make([]string, 0)
	result := make([]string, 0)
	err := configs.DB.Where("status = ?", "BOOKED").Find(&allparkingDB).RowsAffected
	// configs.DB.Where("status = ?", "BOOKED").Find(&allparkingDB).Scan(&allparkingDB)

	configs.DB.Where("status = ?", "BOOKED").Select("status").Find(&allparkingDB).Count(&count)
	configs.DB.Preload(clause.Associations).Find(&allparkingDB, "status = ?", "BOOKED")
	for i = 0; i < count; i++ {

		get_parking_number = append(get_parking_number, strconv.Itoa(int(allparkingDB[i].ParkingNumber)))
		get_registration = append(get_registration, allparkingDB[i].RegistrationNumber)
		get_colour = append(get_colour, allparkingDB[i].Colour)

	}
	for i := 0; i < len(get_parking_number); i++ {
		result = append(result, get_parking_number[i]+" "+get_registration[i]+" "+get_colour[i])
	}
	if err != 0 {

		// return c.JSON(http.StatusOK, models.ParkingResponseManyStatus{
		// 	Code:    http.StatusOK,
		// 	Message: "Success get status parking",
		// 	Status:  "success",
		// 	Data:    allparkingDB[0].ParkingNumber,
		// })

		return c.JSON(http.StatusOK, models.ParkingResponseManyString{
			Data: result,
		})

	}
	return c.JSON(http.StatusInternalServerError, models.ParkingResponseNotif{
		Code:    http.StatusInternalServerError,
		Message: "Parking Data Is Empty",
		Status:  "error",
	})

}
