package main

import (
	"fmt"
	"go_game/src/gnet"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("hello world")
	go func() {
		gnet.Bind()
		wg.Done()
	}()
	go func() {
		gnet.LoopConn()
		wg.Done()
	}()
	wg.Wait()
}
