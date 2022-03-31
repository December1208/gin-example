package main

import (
	"go-example/server"
)

func main() {
	server.Init()
	defer func() {
		err := recover()
		if err != nil {

		}
	}()
}

