package mongodb

import (
	"loginServer/conf"

	"github.com/eric2918/leaf/db/mongodb"
	"github.com/eric2918/leaf/log"
)

var (
	Context *mongodb.DialContext
)

func init() {
	var err error
	Context, err = mongodb.Dial(conf.Server.MongodbAddr, conf.Server.MongodbSessionNum)
	if err != nil {
		log.Fatal("mongondb init is error(%v)", err)
	}
}
