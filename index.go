package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olivere/elastic"
)

const (
	POST_INDEX = "post"
)

type Statics struct {
	ES_URL string `json:"ES_URL"`
}

func main() {
	var ES_URL = readESURL()
	client, err := elastic.NewClient(elastic.SetURL(ES_URL))
	if err != nil {
		panic(err)
	}

	exists, err := client.IndexExists(POST_INDEX).Do(context.Background())

	if err != nil {
		panic(err)
	}
	if !exists {
		mapping := `{
                        "mappings": {
                                "properties": {
                                        "user":     { "type": "keyword", "index": false },
                                        "message":  { "type": "keyword", "index": false },
                                        "location": { "type": "geo_point" },
                                        "url":      { "type": "keyword", "index": false },
                                        "type":     { "type": "keyword", "index": false },
                                        "face":     { "type": "float" }
                                }
                        }
                }`
		_, err := client.CreateIndex(POST_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Post index is created.")
}

func readESURL() string {
	jsonFile, err := os.Open("statics.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var statics Statics
	json.Unmarshal(byteValue, &statics)
	return statics.ES_URL
}
