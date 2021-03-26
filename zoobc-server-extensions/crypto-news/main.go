package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type cryptocompare struct {
	Type       uint8        `json:"Type"`
	Message    string       `json:"Message"`
	Promoted   []int        `json:"Promoted"`
	Newscrypto []newscrypto `json:"Data"`
}

type newscrypto struct {
	ID         string     `json:"id"`
	Link       string     `json:"guid"`
	Publish    uint32     `json:"published_on"`
	Image      string     `json:"imageurl"`
	Title      string     `json:"title"`
	URL        string     `json:"url"`
	Source     string     `json:"source"`
	Body       string     `json:"body"`
	Tags       string     `json:"tags"`
	Categories string     `json:"categories"`
	Upvotes    string     `json:"upvotes"`
	Downvotes  string     `json:"downvotes"`
	Lang       string     `json:"lang"`
	Sourceinfo sourceinfo `json:"source_info"`
}

type sourceinfo struct {
	Name string `json:"name"`
	Lang string `json:"lang"`
	Img  string `json:"img"`
}

var responseObject cryptocompare

func cryptoCompare(w http.ResponseWriter, r *http.Request) {
	//Get the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)

	for i := 0; i < len(responseObject.Newscrypto); i++ {
		fmt.Println(responseObject.Newscrypto[i].ID)
	}
}

func main() {
	//Get the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	path := os.Getenv("PATH")
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc(path, cryptoCompare).Methods("GET")

	log.Fatal(http.ListenAndServe(port, r))
}
