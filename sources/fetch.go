package sources

import (

	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"net/http"
	
)


func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

func Fetch(w http.ResponseWriter, r *http.Request){
	var sign Input

	err := json.NewDecoder(r.Body).Decode(&sign)

	if err != nil {

		var err Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	

url := "https://yahoo-finance97.p.rapidapi.com/stock-info"
	

signature:="symbol="+sign.Symbol

payload := strings.NewReader(signature)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", "5b4d024045msha50d8de10c2e5a2p12d7d5jsna90afb4b0417")
	req.Header.Add("X-RapidAPI-Host", "yahoo-finance97.p.rapidapi.com")

	res, err:= http.DefaultClient.Do(req)
	if err!=nil{
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

fmt.Println(res,"this is the response")

data_obj := Stockinfo{}

jsonErr := json.Unmarshal(body, &data_obj)

if jsonErr != nil {
     log.Fatal(jsonErr) 
}
var info Output
info.Message=data_obj.Message
info.Status=data_obj.Status
info.Data=data_obj.Data["longBusinessSummary"]
fmt.Println(info)
json.NewEncoder(w).Encode(info)
}