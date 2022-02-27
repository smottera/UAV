//////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func initDrone() {

	resp, err := http.Get("http://localhost:8888/")
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	//error checking of the ioutil.ReadAll() request
	if err != nil {
		fmt.Println("ERROR....! ", err)
	}

	fmt.Println(string(bodyBytes))
}

func getMission() {
	fmt.Println("Get mission related data")
}

func updateTelemtryData() {
	fmt.Println("Telemetry Data Updated")
}

func updateEmergencyVariables() {
	fmt.Println("Emergency Variables set!")
}

//GETs latest control link data
func getControlLink() {
	fmt.Println("Get control link related data")

	resp, err := http.Get("http://localhost:8888/")
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	//error checking of the ioutil.ReadAll() request
	if err != nil {
		fmt.Println("ERROR....! ", err)
	}

	b1 := []byte(string(bodyBytes))
	var m map[string]string

	err = json.Unmarshal(b1, &m)
	if err != nil {
		panic(err)
	}

	cBuffer := controlLinkStringTObuffer(m["controlLink"])
	fmt.Println(cBuffer)
}

func controlLinkStringTObuffer(cl string) controlPacket {

	buffer := controlPacket{}
	s := strings.Split(cl, ",")
	buffer.uniqueID = s[0]
	buffer.throttle, _ = strconv.Atoi(s[1])
	buffer.rudder, _ = strconv.Atoi(s[2])
	buffer.aileron, _ = strconv.Atoi(s[3])
	buffer.elevator, _ = strconv.Atoi(s[4])
	buffer.aux1, _ = strconv.Atoi(s[5])
	buffer.aux2, _ = strconv.Atoi(s[6])
	buffer.aux3, _ = strconv.Atoi(s[7])

	return buffer
}

func main() {

	fmt.Println("Hai")

	then := time.Now()
	getControlLink()
	now := time.Now()
	diff := now.Sub(then)
	fmt.Println("The response time is = ", diff)
}
