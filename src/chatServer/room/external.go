package room

import (
	"chatServer/conf"
	"chatServer/room/internal"
	"math"

	"github.com/eric2918/leaf/chanrpc"
	"github.com/eric2918/leaf/module"
)

var (
	modules = []*internal.Module{}
)

type Module interface {
	GetChanRPC() *chanrpc.Server
}

func CreateModules() []module.Module {
	results := []module.Module{}
	for i := 0; i < conf.Server.RoomModuleCount; i++ {
		module := internal.NewModule(i)
		modules = append(modules, module)
		results = append(results, module)
	}
	return results
}

func GetBestModule() (module *internal.Module) {
	var minCount int32 = math.MaxInt32
	for _, _module := range modules {
		count := _module.GetClientCount()
		if count < minCount {
			module = _module
			minCount = count
		}
	}
	return
}
