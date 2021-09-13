package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/buger/jsonparser"
	"github.com/gorilla/mux"
)

//const LOGFILE2 = "/home/vagrant/qunu_zynq_json_bob"

func homePage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	listOfKeys := []string{"ts", "Message", "Qbitrate", "status", "RawDataRate", "Bytes", "Requests"}
	dict := logfileToList("bulk.json", listOfKeys)

	fmt.Fprintf(w, "{")
	count := 0
	sizeOfDict := len(dict) - 1

	for x, y := range dict {
		fmt.Fprintf(w, `"`+x+`":"`+y+`"`)
		if count < sizeOfDict {
			fmt.Fprint(w, ",")
		}
		count = count + 1
		fmt.Println(`{"`+x+`":"`+y+`"}`, len(dict))
	}
	fmt.Fprintf(w, "}")
}

func homePage2(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	listOfKeys := []string{"ts", "Message", "Qbitrate", "status", "RawDataRate", "Bytes", "Requests"}
	dict := logfileToList("bulk.json", listOfKeys)

	fmt.Fprintf(w, "[{")
	count := 0
	sizeOfDict := len(dict) - 1

	for x, y := range dict {
		fmt.Fprintf(w, `"`+x+`":"`+y+`"`)
		if count < sizeOfDict {
			fmt.Fprint(w, ",")
		}
		count = count + 1
		fmt.Println(`{"`+x+`":"`+y+`"}`, len(dict))
	}
	fmt.Fprintf(w, "}]")
}

func testPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w,`[{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-04-12 18:09:21","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:22","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:23","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:24","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:25","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:26","Message":"Bytes","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:27","Message":"Bytes","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:28","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:29","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:30","Message":"Qbitrate","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:31","Message":"Requests","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:32","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:33","Message":"Qbitrate","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:34","Message":"Requests","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:35","Message":"Qbitrate","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:36","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:37","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:38","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:39","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:40","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:41","Message":"Requests","Qbitrate":"70","status":"info"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:42","Message":"Requests","Qbitrate":"70","status":"critical"},
					{"RawDataRate":"70","Bytes":"70","Requests":"70","ts":"2021-05-12 18:09:43","Message":"Requests","Qbitrate":"70","status":"info"}]`)


}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/dbug", testPage).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8888", router))
}



//return the latest key:value
func logfileToList(path string, listOfKeys []string) map[string]string {

	m := make(map[string]string)
	for _, key := range listOfKeys {
		m[key] = "null" //init empty string for each key
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, key := range listOfKeys {
			extract, err0 := jsonparser.GetString([]byte(scanner.Text()), key)
			if err0 == nil {
				m[key] = extract
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}

func main() {

	//listOfKeys := []string{"ts", "Message","Qbitrate", "status","RawDataRate", "Bytes", "Requests"}
	//fmt.Println(logfileToList("bulk.json", listOfKeys))

	//handleRequests()
	//handleRequests2()
}
