package main

import (
	"chatClient/client"
	"chatClient/common"
	"chatClient/common/msg"

	"github.com/eric2918/leaf"
)

func main() {
	common.Init()
	client.Init(msg.Processor)

	leaf.Run(
		client.Module,
	)
}
