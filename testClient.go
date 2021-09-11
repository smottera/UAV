package main

import (
	"context"
	"regexp"
	"fmt"
	"log"
	"time"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"net/http"
	"github.com/elastic/go-elasticsearch/v8"
    "github.com/buger/jsonparser"
)

type keyValue struct {
	fieldName  string
	fieldValue string
}

type fieldRange struct{
	fieldName 	string
	gt 			string
	lte 		string
}

type query_info struct {
	ObservableFields   []string
	FieldMatchings     []keyValue
	SortingFields      []keyValue
	RangeFields		   []fieldRange
	TotalHits			int `default: 100`
}

type dateValue struct{
    date  string
    value float64
}

type queryDict struct{
    key           string
    startTime     string `default: "0000-00-00 00:00:00"`
    endtime       string `default: "2030-00-00 00:00:00"`
    document      []dateValue
}

//---------------------------------
type parameter_base struct {
	Parameter_name                string `json:"parameter_name"`
	Parameter_data_type           string `json:"parameter_type"`
	Parameter_min_allowable_value string `json:"parameter_min_allowable_value"`
	Parameter_max_allowable_value string `json:"parameter_max_allowable_value"`
	Parameter_precision           string `json:"parameter_precision"`
	Parameter_unit                string `json:"parameter_unit"`
}

const MAX_PARAMETER = 100
const MAX_SUBSYSTEMS = 100
const MAX_SYSTEMS = 10

type subsystem_info struct {
	Parameter_list   [] /*MAX_PARAMETER*/ parameter_base
	Subsystem_name   string `json:"subsystem_name"`
	subsystem_health string
}
type system_info struct {
	Subsystem_list [] /*MAX_SUBSYSTEMS*/ subsystem_info `json:"subsystem_list"`
	System_name    string                               `json:"system_name"`
	system_health  string

}

type QKD_network_info struct {
	System_list    []system_info `json:"system_information"`
	network_health string
	indexPattern	string
}

type QKDinfo struct {
	host 	string  `default: "127.0.0.1"`
	port		int `default: 9200`
	indexName string
	qElement string `default: "Alice"`
}

type qSystem struct{
	qList []QKDinfo
	uniqueID string
}

func initSystems() int {
    
    jsonFile, err := os.Open("system_info.json")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users QKD_network_info

	json.Unmarshal(byteValue, &users)


	for i := 0; i < len(users.System_list); i++ {
		fmt.Println("system name: " + users.System_list[i].System_name)
		for j := 0; j < len(users.System_list[i].Subsystem_list); j++ {
			fmt.Println("subsystem name: " + users.System_list[i].Subsystem_list[j].Subsystem_name)
		}
	}
	return 0
}

func initElastic(info QKDinfo, inputQuery string) string{
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://"+ info.host +":" + strconv.Itoa(info.port),
			//"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	r := strings.NewReader(inputQuery)

	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(info.indexName),
		es.Search.WithBody(r),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	return res.String()[9:]
}

func queryBuilder(q1 query_info) string{
	
	s := `{"query": {"bool": { "must": [`

	if len(q1.FieldMatchings) == 0 && len(q1.RangeFields) == 0{
		s += `]`
	}else{
		for i := 0; i < len(q1.FieldMatchings); i++{
			if len(q1.FieldMatchings) == 1 || i == (len(q1.FieldMatchings) -1) {
				s += `{"match":{"`+q1.FieldMatchings[i].fieldName +`":"`+q1.FieldMatchings[i].fieldValue +`"}},`
			} 
		}
	}

	if len(q1.RangeFields) == 0 {
		s = s[:len(s)-1] //get rid of previous comma
	}else{
		for i := 0; i < len(q1.RangeFields); i++{
			if len(q1.RangeFields) == 1 || i == (len(q1.RangeFields) -1) {
				s += `{"range":{"`+q1.RangeFields[i].fieldName +`":{"gt":`+q1.RangeFields[i].gt +`,"lte":` + q1.RangeFields[i].lte + `}}}`
			} else {
				s += `{"range":{"`+q1.RangeFields[i].fieldName +`":{"gt":`+q1.RangeFields[i].gt +`,"lte":` + q1.RangeFields[i].lte + `}}},`
			}
		}
	}

	s += `]}},"_source":[`

	if len(q1.ObservableFields) == 0{
		s += `]`
	}else{
		for i := 0; i < len(q1.ObservableFields); i++{
			if len(q1.ObservableFields) == 1 || i == (len(q1.ObservableFields) -1) {
				s += `"`+q1.ObservableFields[i] +`"]`
			} else {
				s += `"`+q1.ObservableFields[i] +`",`
			}
		}
	}

	s += `,"sort":[`

	if len(q1.SortingFields) > 0 {
		for i := 0; i < len(q1.SortingFields); i++{
			if len(q1.SortingFields) == 1 || i == (len(q1.SortingFields) -1) {
				s += `{"`+q1.SortingFields[i].fieldName +`":{"order":"`+q1.SortingFields[i].fieldValue +`"}}`
			} else {
				s += `{"`+q1.SortingFields[i].fieldName +`":{"order":"`+q1.SortingFields[i].fieldValue +`"}},{`
			}
		}
	}


	s += `],"size" : ` +  strconv.Itoa(q1.TotalHits)
	s += `}`
	return s
}

func extractor(queryResult string, hits int, param []string) []queryDict{
    result := []queryDict{}

    if len(param) > 0 {
        data := []byte(queryResult)

        for p := 0; p < len(param); p++{
            x := []dateValue{}
            for i := 0; i < hits; i++{
                hiterator := strconv.Itoa(i)            
                    
                    timestamp, err0 := jsonparser.GetString(data, "hits", "hits", "[" +hiterator+ "]" , "_source", "@timestamp")
                    hit, err1 := jsonparser.GetFloat(data, "hits", "hits", "[" +hiterator+ "]" , "_source", param[p])
                    
                    if(err0 != nil || err1 != nil){
                        //fmt.Println(err0, err1)
                    	continue
                    }else{
                        
                        modTime := timestamp[:10] + ` ` +timestamp[11:23] //adjust time precision here
                        packet := []dateValue{{modTime, hit}}
                        x = append(x, packet[0])

                    }        
            }

    //populate []queryDict{}
    if len(x) > 0{
            q1 := []queryDict{{param[p], x[0].date, x[len(x)-1].date, x}}
            result = append(result, q1[0])
    		}

    	}
	}
    return result
}

func getILM(info QKDinfo) error {
    
    url := "http://"+ info.host +":" + strconv.Itoa(info.port)
    resp, err := http.Get(url + "/_ilm/policy/")
    
    if err != nil {
        fmt.Println("ERROR....! ", err)
    }

    defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)

    //error checking of the ioutil.ReadAll() request
    if err != nil {
        fmt.Println("ERROR....! ", err)
    }

    bodyString := string(bodyBytes)
    fmt.Println(bodyString)    

    return nil
}

func getIndexAge(info QKDinfo) (string, string) {

	// get all indices, find min/max in each index
	minQuery := `{"_source":["_id","@timestamp"],"sort":[{"@timestamp":{"order":"asc"}}],"size":1}`
	maxQuery := `{"_source":["_id","@timestamp"],"sort":[{"@timestamp":{"order":"desc"}}],"size":1}`

	minResult := initElastic(info, minQuery)
	maxResult := initElastic(info, maxQuery)
    
	oldest, _ := jsonparser.GetString([]byte(minResult), "hits", "hits", "[0]" , "_source", "@timestamp")
	newest, _ := jsonparser.GetString([]byte(maxResult), "hits", "hits", "[0]" , "_source", "@timestamp")

	//fmt.Println("Oldest document: ",oldest)
	//fmt.Println("Newest document: ",newest)

    return oldest, newest
}

//this function just grabs host from QKDinfo, regardless of the index name supplied
//Assuming we're only after filebeat indices (or indices starting with 'f')
func getAllIndicesAge(info QKDinfo, pattern string) (string, string) {
	
	url := "http://"+ info.host +":" + strconv.Itoa(info.port)
    resp, err := http.Get(url + "/_aliases")
    
    if err != nil {
        fmt.Println("ERROR....! ", err)
    }

    defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)

    //error checking of the ioutil.ReadAll() request
    if err != nil {
        fmt.Println("ERROR....! ", err)
    }

    bodyString := string(bodyBytes)

    // a map container to decode the JSON structure into
    c := make(map[string]json.RawMessage)

    // unmarschal JSON
    e := json.Unmarshal([]byte(bodyString), &c)

    // panic on error
    if e != nil {
        panic(e)
    }

    listOfIndices := make([]string, len(c))

    i := 0
    hottestIndex := `` //will store index Name
    tempTime := `` //will store datetime of newest document in index

    for indexName, _ := range c {
        listOfIndices[i] = indexName
        i++
        tempInfo := QKDinfo{info.host, 9200, indexName, "Alice"}  
		_, newest := getIndexAge(tempInfo)
				
		r, _ := regexp.Compile("(?i)" + pattern)

		if newest != `` && r.MatchString(indexName){
			//fmt.Println(indexName)
			if(tempTime < newest){
				tempTime = newest
				hottestIndex = indexName
				
			}
		}
    }

    return hottestIndex, tempTime
}



func main() {
	t1 := time.Now()

	obsFields := []string{"ts", "Message", "Qb_error", "status", "@timestamp"}
	testRange := []fieldRange{{"Qb_error", "0.01", "99"}}
	termMatch := []keyValue{{"status", "info"}}
	sortF := 	 []keyValue{{"@timestamp", "desc"}}
	
	numHits := 10
	
	testQuery := query_info{obsFields, termMatch, sortF, testRange, numHits}
	builtQuery := queryBuilder(testQuery)

	t2 := time.Now()

	info := QKDinfo{"192.168.10.181", 9200, "filebeat-7.2.0-2021.01.12-000001", "Alice"}  
	//info := QKDinfo{"127.0.0.1", 9200, "filebeat-7.2.0-2021.01.12-000001", "Alice"}  
	result := initElastic(info, builtQuery)
    parameters := []string{"Qb_error"}
    
    fmt.Println(extractor(result, numHits, parameters))


	t3 := time.Now()
	
	fmt.Println("___________________________________________")
	fmt.Println("Time Taken by elasticsearch =",t2.Sub(t1))
	fmt.Println("Time Taken by extractor     =",t3.Sub(t2))
	fmt.Println("Time Taken overall          = ",t3.Sub(t1))


	fmt.Println("-------------")

	fmt.Println(getAllIndicesAge(info, "filebeat"))

}