package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var listEndpoint string

func main() {
	//TODO: cleanly shutdown, healthcheck
	flag.StringVar(&listEndpoint, "listEndpoint", "", "the url for the transformer list")
	flag.Parse()

	if listEndpoint == "" {
		log.Fatal("listEndpoint must be provided")
	}

	messages := make(chan string, 128)

	go func() {
		fetchAll(messages)
		close(messages)
	}()

	for msg := range messages {
		fmt.Println(msg)
	}

}

func fetchAll(messages chan<- string) {
	urls := make(chan string, 128)
	go fetchUrlList(urls)

	readers := 32

	readWg := sync.WaitGroup{}

	for i := 0; i < readers; i++ {
		readWg.Add(1)
		go func(i int) {
			fetchMessages(messages, urls)
			log.Printf("fetcher %d exiting\n", i)
			readWg.Done()
		}(i)
	}

	readWg.Wait()
}

func fetchUrlList(urls chan<- string) {

	//	resp, err := httpClient.Get("http://ftaps39395-law1b-eu-t/transformers/organisations/")
	resp, err := httpClient.Get(listEndpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	type listEntry struct {
		APIUrl string `json:apiUrl`
	}

	var list []listEntry
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&list); err != nil {
		panic(err)
	}

	for _, listEntry := range list {
		urls <- listEntry.APIUrl
	}

	close(urls)
}

func fetchMessages(messages chan<- string, urls <-chan string) {
	for url := range urls {
		resp, err := httpClient.Get(url)
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}
		messages <- string(data)
	}

	close(messages)
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 32,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
	},
}
