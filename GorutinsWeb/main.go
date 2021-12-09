package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// wg := sync.WaitGroup{}
	// wr := make(chan struct{})
	for i := 0; i < 100; i++ {

		go MakeRequest()
	}
	time.Sleep(4 * time.Second)
}

func MakeRequest() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
