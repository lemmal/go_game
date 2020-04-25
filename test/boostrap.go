package main

import "go_game/test/client"

func main() {
	client.Connect("tcp", "127.0.0.1 : 2046")
	client.Call("hello")
	client.Close()
}
