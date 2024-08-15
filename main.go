package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SavelyDev/api-client/ipinfo"
)

const API_TOKEN = "YOUR_TOKEN"

func main() {
	ipinfoClient, err := ipinfo.NewClient(time.Second*5, API_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ipinfoClient.GetIPInfo("8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Info())
}
