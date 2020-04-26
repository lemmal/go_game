package main

import (
	"fmt"
	"go_game/src/gnet"
)

func main() {
	fmt.Println("hello world")
	server := gnet.CreateServer("tcp", "127.0.0.1", 2046)
	server.Start()
	select {
	case needShutdown := <-server.GetShutdownChan():
		if needShutdown {
			server.Shutdown()
		}
	}
}
