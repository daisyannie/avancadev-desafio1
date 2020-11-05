package main

import (
	"io/ioutil"
	"github.com/hashicorp/go-retryablehttp"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5
	
	url := "http://localhost:3000/coupons/"+coupon
    response, err := retryClient.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
 
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	responseString := string(responseData)
	
	valid := "invalid"
	if responseString == "OK"{
		valid = "valid"
	}
	result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}