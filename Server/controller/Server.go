package controller

import(
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"

	"../service"
	"../models"
)

//StartServer starts listening to paths
func StartServer(port string)error{
	//create a instance of a mux router
	server := mux.NewRouter()

	//Update Bus Location 
	server.HandleFunc("/updateBusLocation", func (w http.ResponseWriter, r *http.Request){
		
		var bus models.Bus

		body, err := ioutil.ReadAll(r.Body)
		if err!= nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &bus)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			log.Println("ERR: while reading the request")
			log.Println("ERR: ", err)
			return
		}

		rowsAffected, err2 := service.UpdateBusLocation(bus.Location, bus.LicenseNo)
		if err2 != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}else if rowsAffected < 1{
			w.WriteHeader(304)
			return 
		}
		w.WriteHeader(200)
		
				
	}).Methods("PUT")

	//Update Bus Availability
	server.HandleFunc("/updateBusAvailability", func (w http.ResponseWriter, r *http.Request){
		var bus models.Bus
		
		body, err := ioutil.ReadAll(r.Body)
		if err!= nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &bus)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			log.Println("ERR: while reading the request")
			log.Println("ERR: ", err)
			return
		}

		rowsAffected, err2 := service.UpdateBusAvailability(bus.LicenseNo, bus.Attributes.Availability)
		if err2 != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}else if rowsAffected < 1{
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.WriteHeader(http.StatusOK)

	}).Methods("PUT")


	//Get Bus Location
	server.HandleFunc("/getBusLocation/{LicenseNo}", func (w http.ResponseWriter, r *http.Request){
		LicenseNo := mux.Vars(r)["LicenseNo"]
		result, err := service.GetBusLocation(LicenseNo)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	})

	//Get Bus Attributes
	server.HandleFunc("/getBusAttributes/{LicenseNo}", func (w http.ResponseWriter, r *http.Request){
		LicenseNo := mux.Vars(r)["LicenseNo"]
		result, err := service.GetBusAttributes(LicenseNo)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	})

	//Create  new Bus
	server.HandleFunc("/createBus", func (w http.ResponseWriter, r *http.Request){
		var bus models.Bus
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		err = json.Unmarshal(body, &bus)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			log.Println("ERR: while reading the request")
			log.Println("ERR: ", err)
			return
		}
		rowsAffected, err2 := service.CreateBus(bus)
		if err2 != nil{
			if strings.Contains(err2.Error(),"Error 1406"){
				http.Error(w, err2.Error(), http.StatusBadRequest)
				return
			}else if strings.Contains(err2.Error(),"Error 1062"){
				http.Error(w, err2.Error(), http.StatusConflict)
				return
			}
			log.Println(err2)
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}else if rowsAffected < 1{
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")


	err := http.ListenAndServe(port, server)
	return err
}


