package routes

import (
	"github.com/labstack/echo/v4"
	"gitlab.mapan.io/playground/parking-lot-golang/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	//create slot parking with number parameter
	e.POST("/create_parkig_lot/:n", controllers.CreateSlotParkingController)
	//add vehicle to parking and allocated it
	e.POST("/park/:nopol/:colour", controllers.AddingVehicleParking)
	//delete data parking with number parking parameter
	e.POST("/leave/:parkingnumber", controllers.LeaveParking)
	//getting all registration number with colour of vehicle
	e.GET("/cars_registration_numbers/colour/:colour", controllers.Getnopolbycolour)
	//getting all parking number with colour of vehicle
	e.GET("/cars_slot/colour/:colour", controllers.GetSlotParkingByColour)
	//getting parking number with registration number of vehicle
	e.GET("/slot_number/car_registration_number/:reg", controllers.GetSlotParkingByRegistrationNumber)
	//getting all data parking
	e.GET("/status", controllers.GetStatus)

	return e
}
