package service

import(
	"log"

	"../models"
	"../repository"
	)


//UpdateBusLocation Handles UpdateBusLocation
func UpdateBusLocation(Location models.Location, LicenseNo string)(int64, error){
	rowsAffected, err := repository.UpdateLocation(Location, LicenseNo)
	if err != nil{
		log.Println("ERR: Updating The Bus")
		return rowsAffected, err
	}
	return rowsAffected, nil
	
}

//UpdateBusAvailability update the Bus Availability
func UpdateBusAvailability(LicenseNo string, data int)(int64, error){
	rowsAffected, err := repository.UpdateBusAvailability(LicenseNo, data)
	if err != nil{
		return rowsAffected, err
	}
	return rowsAffected, nil
}

//GetBusLocation gets location data
func GetBusLocation(LicenseNo string)(models.Location, error){
	result, err := repository.GetBusLocation(LicenseNo)
	if err != nil{
		return result, err
	}
	
	return result, nil
}

//GetBusAttributes gets attribute data
func GetBusAttributes(LicenseNo string)(models.Attributes, error){
	result, err := repository.GetBusAttributes(LicenseNo)
	if err != nil{
		log.Println("ERR: Getting BusAttributes")
		return result, err
	}
	
	return result, nil
}

//CreateBus creates a new bus
func CreateBus(bus models.Bus)(int64, error){
	rowsAffected, err := repository.CreateBus(bus)
	if err != nil{
		log.Println("ERR: Creating the bus")
		return rowsAffected, err
	}
	return rowsAffected, nil


}