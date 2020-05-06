package repository

import (
	"testing"
	
	"../models"
)

func TestConnect(t *testing.T){
	DB, err := Connect("root:lqsym319@tcp(localhost:3306)/bustracker")
	if err != nil{
		t.Errorf("Error calling Connect()")
	}
	if DB == nil {
		t.Errorf("DB not connected")
	}
}



func BenchmarkConnect(b *testing.B){
	var err error
	b.ResetTimer()
	for i := 0; i <b.N; i++{
		DB, err = Connect("root:lqsym319@tcp(localhost:3306)/bustracker")
		if err != nil {
			b.Fail()
		} 
   }
}

func BenchmarkUpdateStatus(b *testing.B){
	val := models.Location{
			Latitude: 4.234546,
			Longitude: 75.678234,
			Speed: 23,
			Heading: 289,
			}
	var err error
	DB, err = Connect("root:lqsym319@tcp(localhost:3306)/bustracker")
	if err != nil{
		b.FailNow()
	}

	err = StmtPrepare()
	if err != nil{
		b.FailNow()
	}

	b.ResetTimer()
	for i := 0; i <b.N; i++{
	    UpdateLocation(val, "NC-7783")
   }
}



