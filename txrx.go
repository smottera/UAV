//..................THINGS TO DO
//
//2. store all other RCLink, log, telemetry shit in buffer
//4. test speed/reliability of emergency functions and disarm
//5. drone.go should output sbus
//6. Get serial data from USB joystick, send it (make it sendable) to drone.go
//7. Implement dynamic + random test inputs for drone.go json buffered output
//8. Use UDP for faster data transmission. (where latest data arrival matters timewise)
//9. The gRPC framework is a high-performance, platform-neutral standard for making distributed function calls across a network.
//10. Use pointers instead of generating new variables often. Reduces garbage collection effort

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//priority 0
var (
	emergencyFlag bool
	returnToHome  bool
	disarm        bool
	mode          string
	homeLat       float32
	homeLon       float32
	homeAlt       float32
)

//priority 1
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

//priority 2
type missionPacket struct {
	missionID uint
	latitude  float32
	longitude float32
	altitude  float32
	repeat    bool
}

//priority 3
type telemetryPacket struct {
	batteryVoltage float32
	currentDraw    float32
	longitude      float32
	latitude       float32
	altitude       string
	temperature    float32
	motorRPM       int
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
