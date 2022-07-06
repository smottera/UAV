//middleware between backend services and postgres (and/or redis)
//https://golangdocs.com/golang-postgresql-example
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

func initTables() error {

	//Main Table
	query1 := `CREATE TABLE mainTable(
		userID INT PRIMARY KEY,
		firstName VARCHAR(255),
		lastName VARCHAR(255),
		password VARCHAR(255),
		dateOfBirth VARCHAR(255),
		defaultAddress VARCHAR(255),
		phoneNumber INT,
		viewCount INT,
		userType VARCHAR(255),
		verified BOOL,
		listOfProperties VARCHAR(255),
		listOfMissions VARCHAR(255),
		listOfPayments VARCHAR(255),
		thumbnail VARCHAR(255)
		);`

	//Property Details
	query2 := `CREATE TABLE propertyDetails(
		uniqueID INT PRIMARY KEY,
		ownerID INT,
		propertyName VARCHAR(255),
		propertyType VARCHAR(255),
		registeredDate VARCHAR(255),
		purchaseType VARCHAR(255),
		address VARCHAR(255),
		area INT,
		value INT,
		rating INT,
		reviews VARCHAR(255),
		description VARCHAR(255),
		);`

	//Payment History
	query4 := `CREATE TABLE paymentHistory(
		uniqueID INT PRIMARY KEY,
		userID INT,
		timestamp VARCHAR(255),
		paymentMode INT,
		amount VARCHAR(255),
		status VARCHAR(255),
		subscription VARCHAR(255),
		currency VARCHAR(255),
		location VARCHAR(255),
		memo VARCHAR(255)
		);`

	//Mission Details (ATC)
	query3 := `CREATE TABLE missionDetails(
		name VARCHAR(255) NOT NULL PRIMARY KEY,
		numOfPanels INT
		);`

	//Blackbox / flight logs
	query5 := `CREATE TABLE telemetryBlackbox(
		timestamp VARCHAR(255),
		UAVid VARCHAR(255),
		Altitude VARCHAR(255),
		Attitude VARCHAR(255),
		temperature VARCHAR(255),
		pressure VARCHAR(255),
		gyro VARCHAR(255),
		accel VARCHAR(255),
		speed VARCHAR(255),
		latitude VARCHAR(255),
		longitude VARCHAR(255),
		batteryVoltage VARCHAR(255),
		currentDraw VARCHAR(255),
		signalStrength VARCHAR(255),
		throttle VARCHAR(255),
		rudder VARCHAR(255),
		elevator VARCHAR(255),
		aileron VARCHAR(255)
		);`

	//Website logs and stats
	query8 := `CREATE TABLE websiteLogsAndStats(
		name VARCHAR(255) NOT NULL PRIMARY KEY,
		numOfPanels INT
		);`

	query6 := `ALTER TABLE mainTable add FOREIGN KEY(uniqueID) REFERENCES mainTable(uniqueID);`

	query7 := `ALTER TABLE propertyDetails add FOREIGN KEY(ownerID) REFERENCES mainTable(uniqueID);`

	fmt.Println(query1, query2, query3, query4, query5, query6, query7, query8)

	return nil
}

func connectToPSQL() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func connectToRedis() {
	fmt.Println("Go Redis!")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func main() {
	//connectToPSQL()
	connectToRedis()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
