package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/davidnarayan/go-extrahop"
)

// Command line flags
var (
	host   = flag.String("host", "extrahop", "ExtraHop host")
	apikey = flag.String("apikey", "", "ExtraHop API key")
)

func main() {
	flag.Parse()

	// Create new ExtraHop client
	eh := extrahop.NewClient(*host, *apikey)

	// Get information about this ExtraHop
	f, err := eh.Get("/extrahop")

	if err != nil {
		log.Fatal(err)
	}

	m := f.(map[string]interface{})

	fmt.Println("ExtraHop version:", m["version"])
	fmt.Println("ExtraHop ECM:", m["ecm"])
}
