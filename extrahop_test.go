package extrahop

import (
	"errors"
	"flag"
	"fmt"
	"testing"
)

// These flags allow the connection information to be specified as command
// line options
var host = flag.String("extrahop.host", "extrahop", "ExtraHop host")
var apikey = flag.String("extrahop.apikey", "", "ExtraHop API key")

func checkFlags() error {
	if *host == "" {
		return errors.New("Host not set. Use -extrahop.host=<host>")
	}

	if *apikey == "" {
		return errors.New("API key not set. Use -extrahop.apikey=<apikey>")
	}

	return nil
}

func TestNewClient(t *testing.T) {
	if err := checkFlags(); err != nil {
		t.Fatal(err)
	}

	eh := NewClient(*host, *apikey)

	if eh.Host != *host {
		t.Errorf("Host not set correctly:\nwant\t%s\ngot\t%s", *host, eh.Host)
	}

	if eh.ApiKey != *apikey {
		t.Errorf("ApiKey not set correctly:\nwant\t%s\ngot\t%s", *apikey, eh.ApiKey)
	}
}

func TestGet(t *testing.T) {
	if err := checkFlags(); err != nil {
		t.Fatal(err)
	}

	eh := NewClient(*host, *apikey)
	path := "/extrahop"

	f, err := eh.Get(path)

	if err != nil {
		t.Fatalf("Error getting %s: %s", path, err.Error())
	}

	fmt.Println(f)

	// A valid response should have the version and ecm keys
	m := f.(map[string]interface{})

	if _, ok := m["version"]; !ok {
		t.Fatalf("Missing key 'version' in response from %s", path)
	}

	if _, ok := m["ecm"]; !ok {
		t.Fatalf("Missing key 'ecm' in response from %s", path)
	}
}

/*
func TestExplore(t *testing.T) {
	if err := checkFlags(); err != nil {
		t.Fatal(err)
	}

	eh := NewClient(*host, *apikey)
	eh.Explore()
}
*/
