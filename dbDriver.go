//middleware between backend services and postgres (and/or redis)
//https://golangdocs.com/golang-postgresql-example
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

func main() {
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
	defer rows.Close()
for rows.Next() {
    var name string
    var roll int
 
    err = rows.Scan(&name, &roll)
    CheckError(err)
 
    fmt.Println(name, roll)
}
}



func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
