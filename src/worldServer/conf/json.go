package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/eric2918/leaf/log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string

	MongodbAddr       string
	MongodbSessionNum int

	ServerName      string
	ListenAddr      string
	ConnAddrs       map[string]string
	PendingWriteNum int
}

func init() {
	data, err := ioutil.ReadFile("conf/worldServer.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
