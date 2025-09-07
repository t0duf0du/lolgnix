package main

import (
	server "github.com/codecrafters-io/http-server-starter-go/internal/server"
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("hello world")
	s, err := server.NewServer("localhost", 9779)
	if err != nil {
		pp.Println("OH SHITTTTT")
	}
	err = s.Run()
	if err != nil {
		pp.Println(err)
	}
}
