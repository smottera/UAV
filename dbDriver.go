//middleware between backend services and postgres (and/or redis)
//https://golangdocs.com/golang-postgresql-example

/*
postgres notes:
 \l lists all databases
 \du+ lists all roles
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO <userName>;
*/

/*
Things to do:
create table relationships -> Test insertion of values with random data ->


User stories:
1. UAV automatically uploads image/video data to cloud
2. UAV uploads blackbox logs
3. Customer searches for a property
4. Customer signs up, logs in, out, controls account
5. Mission controller has updated missions
*/
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = " "
	dbname   = "MyDB"
)

var ctx = context.Background()

func pingPSQL() {
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

func rowsToStrings(rows *sql.Rows) [][]string {
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	pretty := [][]string{cols}
	results := make([]interface{}, len(cols))
	for i := range results {
		results[i] = new(interface{})
	}
	for rows.Next() {
		if err := rows.Scan(results[:]...); err != nil {
			panic(err)
		}
		cur := make([]string, len(cols))
		for i := range results {
			val := *results[i].(*interface{})
			var str string
			if val == nil {
				str = "NULL"
			} else {
				switch v := val.(type) {
				case []byte:
					str = string(v)
				default:
					str = fmt.Sprintf("%v", v)
				}
			}
			cur[i] = str
		}
		pretty = append(pretty, cur)
	}
	return pretty
}

func psqlMapping(option1 int, customQuery string) error {
	////////////        Hardcoded Queries         ////////////////
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
		propertyID INT PRIMARY KEY,
		userID INT,
		propertyName VARCHAR(255),
		propertyType VARCHAR(255),
		registeredDate VARCHAR(255),
		purchaseType VARCHAR(255),
		address VARCHAR(255),
		area INT,
		value INT,
		rating INT,
		reviews VARCHAR(255),
		description VARCHAR(255)
		);`

	//Payment History
	query3 := `CREATE TABLE paymentHistory(
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
	query4 := `CREATE TABLE missionDetails(
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
	query6 := `CREATE TABLE websiteLogsAndStats(
				name VARCHAR(255) NOT NULL PRIMARY KEY,
				numOfPanels INT
				);`

	//query6 := `ALTER TABLE mainTable add FOREIGN KEY(uniqueID) REFERENCES mainTable(uniqueID);`

	query7 := `SELECT * FROM information_schema.tables;`
	//////////////////////////////////////////////////////////////////////////////////
	//open psql connection
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	//Query multiplexing
	switch option1 {
	//INSERT, UPDATE
	case 0:
		_, err2 := db.Exec(customQuery)
		fmt.Println("zero", err2)

	//CREATE tables
	case 1:
		_, err2 := db.Exec(query1)
		fmt.Println("Creating  table mainTable ... ", err2)
		_, err3 := db.Exec(query2)
		fmt.Println("Creating  table propertyDetails ... ", err3)
		_, err4 := db.Exec(query3)
		fmt.Println("Creating  table paymentHistory ... ", err4)
		_, err5 := db.Exec(query4)
		fmt.Println("Creating  table missionDetails ... ", err5)
		_, err6 := db.Exec(query5)
		fmt.Println("Creating  table telemetryBlackbox ... ", err6)
		_, err7 := db.Exec(query6)
		fmt.Println("Creating  table websiteLogs ... ", err7)

	//show all tables
	case 2:
		rows, err2 := db.Query(query7)
		outp := rowsToStrings(rows)
		fmt.Println("Executing Query 7 ... ", outp, err2)

	//DROP ALL tables
	case 3:
		_, err2 := db.Exec(`DROP TABLE IF EXISTS mainTable, propertydetails, 
		paymenthistory, missiondetails, telemetryblackbox, websitelogsandstats CASCADE;`)
		fmt.Println("Dropping ALL tables ... ", err2)

	case 4:
		fmt.Println("Yo")
	}

	// close database
	defer db.Close()

	fmt.Println(option1)
	return nil
}

func pingRedis() error {
	fmt.Println("Go Redis!")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	return nil
}

func redisMapping(customQuery string) error {
	fmt.Println(customQuery)
	return nil
}

func pingElasticsearch() error {
	return nil
}

func elasticsearchMapping(customJSON string) error {
	fmt.Println("Yoo", customJSON)
	return nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	pingPSQL()
	psqlMapping(1, "")
	//pingRedis()

}
