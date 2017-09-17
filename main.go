package main

import (
	"log"
	"os"
	"net/http"
	"time"
)

const (
	duration = 5 * time.Second
	offset   = 20 * time.Second
)

func ping(domain string, client *http.Client) int {
	resp, err := client.Head(domain)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	return resp.StatusCode
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please enter URL")
	}
	domain := os.Args[1:2][0]
	timeout := time.Duration(duration)
	client := http.Client{
		Timeout: timeout,
	}
	if ping(domain, &client) == 200 {
		log.Println("OK")
		os.Exit(0)
	}
	timeEnd := time.Now().Local().Add(time.Duration(offset))
	for x := range time.Tick(duration) {
		if ping(domain, &client) == 200 {
			log.Println("OK")
			break
		}
		log.Println("NO")
		if x.After(timeEnd) {
			break
		}
	}
}
