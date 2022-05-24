package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/eric2918/leaf/log"
)

var Client struct {
	LoginAddr string
}

func init() {
	data, err := ioutil.ReadFile("conf/client.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Client)
	if err != nil {
		log.Fatal("%v", err)
	}
}
