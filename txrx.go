package main

import (
	"fmt"
	"time"
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
func main() {
	now := time.Now()
	fmt.Println("The time is ", now)

	testPayload := controlPacket{uniqueID: "asd"}
	string1 := "{uniqueID: \"" + testPayload.uniqueID + "\"}"
	fmt.Println(string1)

	then := time.Now()
	diff := then.Sub(now)
	fmt.Println("Execution time = ", diff)
}
