package main

import (
	"fmt"
	"go_game/src/gnet"
)

func main() {
	fmt.Println("hello world")
	server := gnet.CreateServer("tcp", "0.0.0.0", 12001)
	server.Start()
	select {
	case needShutdown := <-server.GetShutdownChan():
		if needShutdown {
			server.Shutdown()
		}
	}
}
