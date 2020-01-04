package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	endpoint = flag.String("endpoint", "", "The endpoint of the Google Home device")
	port     = flag.String("port", "8008", "The port of the Google Home device")
)

func main() {
	flag.Parse()
	if *endpoint == "" {
		log.Fatal("Endpoint must be set")
	}
	log.Printf("[main] Device (%s:%s)", *endpoint, *port)
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	httpClient := &http.Client{
		Transport: transport,
	}
	client := NewClient(httpClient, *endpoint, *port)
	info, err := client.DeviceInfo()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", info)
}
