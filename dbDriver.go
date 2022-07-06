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
