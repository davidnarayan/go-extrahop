package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/davidnarayan/extrahop"
)

// Command line flags
var (
	host   = flag.String("host", "extrahop", "ExtraHop host")
	apikey = flag.String("apikey", "", "ExtraHop API key")
)

func main() {
	flag.Parse()

	// Create new ExtraHop client
	eh := extrahop.Client()

	// Get information about this ExtraHop
	f, err := eh.Get("/extrahop")

	if err != nil {
		log.Fatal(err)
	}

	m := f.(map[string]interface{})

	fmt.Println("ExtraHop version: %s", m["version"])
	fmt.Println("ExtraHop ECM: %s", m["ecm"])
}
