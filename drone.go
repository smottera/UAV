package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func getMission() {
	fmt.Println("Get mission related data")
}

func getControlLink() {
	fmt.Println("Get control link related data")

}

func main() {
	now := time.Now()
	fmt.Println("The time is ", now)

	resp, err := http.Get("http://localhost:8888/dbug")
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	//error checking of the ioutil.ReadAll() request
	if err != nil {
		fmt.Println("ERROR....! ", err)
	}

	fmt.Println(json.NewDecoder(bodyBytes))
}
