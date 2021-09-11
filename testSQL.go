package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
)

type params struct{
    param       string
    separateAxis bool
    rightAxis bool
    panelID     int     //FKey

}

type panel struct {
    paramList   []params
    panelID     int         //Primary Key
    name         string      //FKey
    numOfPanels     int
    streamingType bool `default: false` //true == live mode
    quadrant    int 
    startTime   string
    endTime     string
    timeDelta   string
}

type frontend struct{
    name    string //Pkey
    panels []panel
    numOfPanels int
} 

func B2i(b bool) int {
    if b {
        return 1
    }
    return 0
}

func testProgram(numOfRows int, f1 frontend) error {

    for i := 0; i < numOfRows; i++ {
        addPanels(f1)
    }

    return nil
}

func initTables() error{

    query1 := `CREATE TABLE params(
    panelID INT PRIMARY KEY AUTO_INCREMENT,
    param VARCHAR(255),
    separateAxis boolean not null default 0,
    rightAxis boolean not null default 0
    );`
    
    query2 := `CREATE TABLE panels(
    startTime VARCHAR(255),
    endTime VARCHAR(255),
    deltaTime VARCHAR(255),
    numOfPanels INT,
    quadrant INT,
    streamingType boolean not null default 0,
    name VARCHAR(255),
    panelID INT PRIMARY KEY AUTO_INCREMENT
    );`
    
    query3 := `CREATE TABLE frontend(
    name VARCHAR(255) NOT NULL PRIMARY KEY,
    numOfPanels INT
    );`

    query4 := `ALTER TABLE panels add FOREIGN KEY(name) REFERENCES frontend(name);`
    
    query5 := `ALTER TABLE params add FOREIGN KEY(panelID) REFERENCES panels(panelID);`

    db, err := sql.Open("mysql", "newuser:password@tcp(127.0.0.1:3306)/qunu")
    defer db.Close()

    rows, err := db.Query(`SELECT * FROM params INNER JOIN panels ON params.panelID=panels.panelID;`)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(rows)

    fmt.Println(query1, query2, query3, query4, query5)
    return nil
}

func delete(rowID int) error{

    deleteID := `DELETE frontend, panels FROM frontend INNER JOIN panels ON frontend.panelID=panels.panelID WHERE panelID=` + strconv.Itoa(rowID) + `;`
    
    db, err := sql.Open("mysql", "newuser:password@tcp(127.0.0.1:3306)/qunu")
    defer db.Close()

    if rowID == -1 {
    
        _, err = db.Query(`DELETE FROM params;`)
        _, err = db.Query(`DELETE FROM panels;`)
        _, err = db.Query(`DELETE FROM frontend;`)
    
    if err != nil {
            panic(err)
        }
    
    }else{
        _, err = db.Query(deleteID)
            if err != nil {
            panic(err)
        }
    }


    return  nil
}

func addPanels(row frontend) error{

    for _, pan := range row.panels{

        for _, b := range pan.paramList{
            insert3 := `INSERT INTO params(param, separateAxis, rightAxis, panelID) VALUES ("`+b.param+`", ` +strconv.Itoa(B2i(b.separateAxis)) +`, ` + strconv.Itoa(B2i(b.rightAxis))+`, `+strconv.Itoa(b.panelID)+`);`

            //fmt.Println(insert3)

            db, err := sql.Open("mysql", "newuser:password@tcp(127.0.0.1:3306)/qunu")
            defer db.Close()

            _, err = db.Query(insert3)
            if err != nil {
                panic(err)
            }
        }
        insert2 := `INSERT INTO panels(startTime, endTime, deltaTime, numOfPanels, quadrant, streamingType, name, panelID) VALUES ("`+pan.startTime+`", "`+pan.endTime+`", "`+pan.timeDelta+`", `+strconv.Itoa(pan.numOfPanels)+`, `+strconv.Itoa(pan.quadrant)+`,`+ strconv.Itoa(B2i(pan.streamingType)) +`, "`+pan.name+`", `+strconv.Itoa(pan.panelID)+`);`

        db, err := sql.Open("mysql", "newuser:password@tcp(127.0.0.1:3306)/qunu")
            defer db.Close()

            _, err = db.Query(insert2)
            if err != nil {
                panic(err)
            }

    }
    
    insert1 := `INSERT INTO frontend(name, numOfPanels) VALUES("`+ row.name +`", `+strconv.Itoa(row.numOfPanels) +`);`
        
    fmt.Println(insert1)

    db, err := sql.Open("mysql", "newuser:password@tcp(127.0.0.1:3306)/qunu")
            defer db.Close()

            _, err = db.Query(insert1)
            if err != nil {
                panic(err)
            }
    return nil
}

func main() {

    /*db, err := sql.Open("mysql", "newuser:password@tcp(127.0.0.1:3306)/qunu")
    defer db.Close()

    rows, err := db.Query(`SELECT * FROM frontend;`)
    
    if err != nil {
        panic(err)
    }
    
    fmt.Println(rows)
*/
    pID := 72
    name := `blah12`
    par := []params{{"a", true, true, pID}, {"b", true, false, pID+2}, {"c", true, true, pID+3}}
    panel := []panel{{par, pID, name, 3, false, 0, "startt", "endt", "deltat"}}
    test := frontend{name, panel, 2}

    addPanels(test) 

    //delete(-1)




}

