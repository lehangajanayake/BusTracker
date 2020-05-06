package repository

import(
	"log"
	"database/sql"
	"fmt"

	//Used to work with mysql DB
	_ "github.com/go-sql-driver/mysql"

	"../models"
)
//SQL Statements
var (
	updateLocationStmt, 
	getLocationStmt,
	updateAvailabilityStmt, 
	getAvailabilityStmt *sql.Stmt
)

//DB Database obj
var DB *sql.DB

//Connect connects to the db
func Connect(path string) (*sql.DB, error){
	//TODO: Commit feature
	log.Println("INFO: Connecting to the DB")
	db, err := sql.Open("mysql", path)
	log.Println("INFO: Connected to the DB")
	return db, err
}


//StmtPrepare prepares stmt
func StmtPrepare()error{
	log.Println("INFO: Preparing the DB")
	var err error

	updateLocationStmt, err = DB.Prepare("UPDATE buslocation SET Latitude=?, Longitude=?, Speed=?, Heading=? WHERE LicenseNo=?")
	if err != nil{
		return fmt.Errorf("ERR: Preparing updateLocationStmt : ", err)
	}

	getLocationStmt, err = DB.Prepare("SELECT Latitude, Longitude, Heading, Speed from buslocation WHERE LicenseNo=?")
	if err != nil{
		return fmt.Errorf("ERR: Preparing getLocationStmt : ", err)
	}

	updateAvailabilityStmt, err = DB.Prepare("UPDATE busattributes SET Availability=? WHERE LicenseNo=?")
	if err != nil{
		return fmt.Errorf("ERR: Preparing updateAvailabilityStmt : ", err)
	}

	getAvailabilityStmt, err = DB.Prepare("SELECT Availability FROM busattributes WHERE LicenseNo=?")
	if err != nil{
		return fmt.Errorf("ERR: Preparing getAvailabilityStmt : ", err)
	}
	// createBusLocationStmt, err = DB.Prepare(`INSERT INTO buslocation(LicenseNo, Latitude, Longitude, Speed, Heading) 
	// 												VALUES(?,?,?,?,?) 
	// 												ON DUPLICATE KEY UPDATE Latitude=?, Longitude=?, Speed=?, Heading=?`)
	// if err != nil{
	// 	return fmt.Errorf("ERR: Preparing createbusLocationStmt : ", err)
	// }
	// createBusAttributesStmt, err  = DB.Prepare(`INSERT INTO busattributes (LicenseNo, PathNo, Availability, AC) 
	// 											VALUES(?,?,?,?) 
	// 											ON DUPLICATE KEY UPDATE PathNo=?, Availability=?, AC=?`)
	// if err != nil {
	// 	return fmt.Errorf("ERR: Preparing createBusAttributesStmt")
	// }

	log.Println("INFO: Prepared the DB")

	return nil
}

//UpdateLocation : Updates the location of  the database
func UpdateLocation(data models.Location, LicenseNo string)(int64, error){
	result, err := updateLocationStmt.Exec(
		data.Latitude,
		data.Longitude,
		data.Speed, 
		data.Heading, 
		LicenseNo,
	)
	
	rowsAffected , _ := result.RowsAffected()
	return rowsAffected , err
}

//UpdateBusAvailability update Bus Availability
func UpdateBusAvailability(LicenseNo string, data int)(int64, error){
	result, err := updateAvailabilityStmt.Exec(data, LicenseNo)
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

//GetBusLocation Gets the Bus Location
func GetBusLocation(LicenseNo string)(models.Location, error){
	var result models.Location
	err := getLocationStmt.QueryRow(LicenseNo).Scan( 
		&result.Latitude, 
		&result.Longitude, 
		&result.Heading,
		&result.Speed)
	if err != nil{
		return result, err
	}
	return result, nil
}

//GetBusAttributes gets bus attributes data
func GetBusAttributes(LicenseNo string)(models.Attributes, error){
	var result models.Attributes
	err := DB.QueryRow("SELECT PathNo, AC, Availability FROM busattributes WHERE LicenseNo=?", LicenseNo).Scan(
		&result.PathNo,
		&result.AC,
		&result.Availability,
	)
	if err != nil{
		return result, err
	}
	return result, nil
}

//CreateBus creates a new bus
func CreateBus(bus models.Bus)(int64, error){
	tx, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	var result, result2 sql.Result
	result,  err = tx.Exec("INSERT INTO buslocation (LicenseNo, Latitude, Longitude, Speed, Heading) VALUES(?,?,?,?,?) ",
		bus.LicenseNo,
		bus.Location.Latitude,
		bus.Location.Longitude,
		bus.Location.Speed,
		bus.Location.Heading,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return 0, err
	}

	result2, err = tx.Exec("INSERT INTO busattributes (LicenseNo, PathNo, Availability, AC) VALUES(?,?,?,?) ",
	bus.LicenseNo, 
	bus.Attributes.PathNo, 
	bus.Attributes.Availability, 
	bus.Attributes.AC,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return 0, err
	}
	tx.Commit()

	finalresult, _ := result.RowsAffected()
	finalresult2, _ := result2.RowsAffected()
	rowsAffected := finalresult + finalresult2
	return rowsAffected, nil
}
