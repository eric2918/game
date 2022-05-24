package common

import (
	"chatServer/conf"
	"sync/atomic"

	"github.com/eric2918/leaf/cluster"
)

var (
	clientCount int32
)

func GetClientCount() int {
	return int(atomic.LoadInt32(&clientCount))
}

func AddClientCount(delta int32) {
	count := atomic.AddInt32(&clientCount, delta)
	cluster.Go("world", "UpdateChatInfo", conf.Server.ServerName, int(count))
}
