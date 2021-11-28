package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//drone to cloud
type telemetry struct {
	batteryVoltage   float32
	avgCellVoltage   float32
	totalCurrentDraw float32
	latitude         string
	longitude        string
	destination      string
	home             string
	gyro             string
	accelerometer    string
}

type controlVehicle struct {
	speed         float32
	emergencySTOP bool
	takeOffPermit bool
	Altitude	float32
	rangeDistance	float32
	
}

func homePage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, `[{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-04-12 18:09:21","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:22","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:23","Message":"Requests","Qbitrate":"70","status":"info"}]`)
	fmt.Println("This is working!")
}

func homePage2(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, `[{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-04-12 18:09:21","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:22","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:23","Message":"Requests","Qbitrate":"70","status":"info"}]`)
	fmt.Println("This is working!")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func main() {
	handleRequests()
}
