package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
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

func cryptoCompare(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("https://min-api.cryptocompare.com/data/v2/news/?lang=EN")
	if err != nil {
		fmt.Println("Error", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject cryptocompare
	json.Unmarshal(responseData, &responseObject)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)

	for i := 0; i < len(responseObject.Newscrypto); i++ {
		fmt.Println(responseObject.Newscrypto[i].ID)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/news", cryptoCompare).Methods("GET")

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	log.Fatal(http.ListenAndServe(":8000", r))
}
