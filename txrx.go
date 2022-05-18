//..................THINGS TO DO

/*
2. store all other RCLink, log, telemetry shit in buffer
4. test speed/reliability of emergency functions and disarm
5. drone.go should output sbus
6. Get serial data from USB joystick, send it (make it sendable) to drone.go
7. Implement dynamic + random test inputs for drone.go json buffered output
9. The gRPC framework is a high-performance, platform-neutral standard for making distributed function calls across a network.
10. Use pointers instead of generating new variables often. Reduces garbage collection effort
11. gRPC server/client or bidirectional setup could be 5x faster than net/http JSON
13. Implement a watchdog function
14. Implement security measures for backend (software security. Hackproof)
15. Implement video transfer protocol
16. Implement camera zoom, focus and gimbal control functions
20. Implement Drone Social Network (drone ID, drone specific data, authorization)
21. Implement Unmanned Traffic Management function with the help of the above info
23. Implement necessary backend storage and caching services as needed
24. Design and build a scheduler?


*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//priority 0
var (
	timestamp     bool
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
	altitude       float32
	temperature    float32
	motorRPM       int
}

func droneDummyDataGenerator(iterate int, delay int) {

	for j := 0; j <= iterate; j++ {

		time.Sleep(time.Millisecond * time.Duration(delay))
		//init these variables

		rand1 := float32(rand.Float64())
		timestamp := time.Now()
		emergencyFlag = false
		disarm = true
		returnToHome = false

		mode = " startup"
		homeLat = 123.321 * rand1
		homeLon = 987.789 * rand1
		homeAlt = 6900.32 * rand1

		t1 := telemetryPacket{12.3 * rand1, 123.2 * rand1, 123123.3123123 * rand1, 3333.555 * rand1, 23.46, 69.4 * rand1, 9999}
		m1 := missionPacket{123, 32.55 * rand1, 112.34 * rand1, 12.99 * rand1, false}
		c1 := controlPacket{"niggaOne", 0, 1, 2, 3, 4, 5, 6, 7}

		fmt.Println("Yoo  ", rand1, timestamp, t1, m1, c1, emergencyFlag, mode, disarm, returnToHome, homeAlt, homeLat, homeLon)
	}
}

func initSys() {
	fmt.Println("Helllloooo")
	batteryService()
	//get list of of all peripherals connected
	//check wireless connections
	//check communication with HQ/cloud
	//check if flight controls are working correctly
	//check if data is correctly being logged

	fmt.Println("All Systems Go!")
}

func batteryService() {
	//check health of all cells
	//log data
	//calculate range, lifespan, next service date
	fmt.Println("battery okay ...")
}

func missionService() {
	fmt.Println("Ready for any mission ...!")
	//inputs: mission start time, deadline (time), description, distance,
	// path[[gcode]], landing bool?, payload info, attempts, failure?complete?,
	//return address
	//backup return address
	//check mission queue

}

func landingService() {
	fmt.Println("Preparing for landing ...")
	//stop motion, reset attitude
	//plan path for landing
	//begin path to decend
	//check sensors for possibilty of landing
	//continue phase 2 of decent
	//touch down
	//poweroff main motor
	//Disarm
	fmt.Println("Landing complete.")
}

func PathPlanneringProMax() {
	fmt.Println("Path planned.")
	//inputs:starting,ending points,
	//regions to avoid, regions to pass through
	//collion prevention with other UAVs
	//path dependent on time?
	//output: path coordinates, keep track of timing

}

func commsHeartBeat() {
	//check ping/latency with all comms
	//measure signal strength, in proximity UAVs and HQs
	//log everything
	fmt.Println("All comms okay.")

	//check if connection with HQ is okay
	//check if sbus output to microcontroller is okay
	//in case of signal loss: continue with current path? path plan a new Return to home then land
}
func getJoystickData() {
	fmt.Println("NIGGA PWEASE")
}

func main() {

	fmt.Println("Lets implement some shit")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(2000))

	droneDummyDataGenerator(10, 100)
}
