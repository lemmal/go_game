package main

import (
	"fmt"
	"go_game/src/net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("hello world")
	go func() {
		net.Bind()
		wg.Done()
	}()
	go func() {
		net.LoopConn()
		wg.Done()
	}()
	wg.Wait()
}
