package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	emergencyFlag bool
	returnToHome  bool
	disarm        bool
)

type controlPacket struct {
	uniqueID   string
	throttle   int
	rudder     int
	aileron    int
	elevator   int
	motorPower int
	aux1       int
	aux2       int
	aux3       int
}

type missionPacket struct {
	latitude  float32
	longitude float32
	altitude  float32
	mode      string
}

type telemetryPacket struct {
	batteryVoltage  float32
	currentDraw     float32
	currentLocation string
	currentAttitude string
	temperature     float32
	motorRPM        int
}

func initSys() {
	fmt.Println("yiii")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/dbug", testPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func testPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, `{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-04-12 18:09:21","Message":"Requests","Qbitrate":"70","status":"info"}`)

}

func main() {
	now := time.Now()
	fmt.Println("The time is ", now)

	testPayload := controlPacket{uniqueID: "asd"}
	string1 := "{\"uniqueID\": \"" + testPayload.uniqueID + "\"}"
	fmt.Println(string1)

	handleRequests()

	then := time.Now()
	diff := then.Sub(now)
	fmt.Println("Execution time = ", diff)
}
