package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}

func main() {
	MakeRequest()
}
