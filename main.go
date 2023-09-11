package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Parameters struct {
	Result    string `json:"result"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	email := os.Getenv("email")
	password := os.Getenv("password")
	const and = "&"
	const domainType = "Free"

	url := "https://api.freenom.com/v2/domain/search"
	url += "?domain=mk0iscool.tk"
	url += and + email
	url += and + password
	url += and + domainType

	fmt.Println(url)

	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Parse XML response
	params := Parameters{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&params)
	if err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}

	// Print the result
	fmt.Printf("Result: %v\n", params.Result)
	fmt.Printf("Status: %v\n", params.Status)
	fmt.Printf("Timestamp: %v\n", params.Timestamp)
}
