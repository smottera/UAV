//..................THINGS TO DO
//2. store all other RCLink, log, telemetry shit in buffer
//4. test speed/reliability of emergency functions and disarm
//5. drone.go should output sbus
//6. Get serial data from USB joystick, send it (make it sendable) to drone.go

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
	throttle   int //Channel 0
	rudder     int
	aileron    int
	elevator   int
	motorPower int
	aux1       int
	aux2       int
	aux3       int //Channel 7
}

//Control syntax
//{controlLink: "UAV007,1024,0,4096,3123,1123,412,4000,23"}

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

func getJoystickData() {
	fmt.Println("NIGGA PWEASE")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", testPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func testPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, testInputs())

}

//test ONLY CONTROL, ONLY STATUS, ONLY MISSION, ALL dATA, all permutations
func testInputs() string {

	string1 := "{"
	string2 := "\"Timestamp\": \"" + time.Now().String() + "\""
	string3 := ",\"controlLink\" : \"UAV007,1024,0,4096,3123,1123,412,4000,23\""

	dispatch := string1 + string2 + string3 + "}"

	fmt.Println(dispatch)

	return dispatch
}

func main() {

	testInputs()
	handleRequests()

}
