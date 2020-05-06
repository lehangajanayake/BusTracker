package main

import (
	"log"

	"./controller"
	"./repository"
	// "./models"
)

//Con checking go routines
func main() {
	log.Println("INFO: Starting Bus tracker server v1.0")

	var err error

	repository.DB, err = repository.Connect("root:lqsym319@tcp(localhost:3306)/bustracker")
	if err != nil {
		log.Fatal(err.Error())

	}

	//Prepare the Statement
	err = repository.StmtPrepare()
	if err != nil {
		log.Fatal(err.Error())
	}
	//Start Listening
	err = controller.StartServer(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
