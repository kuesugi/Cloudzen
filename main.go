package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
	Url      string   `json:"url"`
	Type     string   `json:"type"`
	Face     float32  `json:"face"`
}

type Statics struct {
	PORT string `json:"PORT"`
}

func main() {
	var PORT = readPort()
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func readPort() string {
	jsonFile, err := os.Open("statics.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var statics Statics
	json.Unmarshal(byteValue, &statics)
	return statics.PORT
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}
