package main

import (
	server "github.com/codecrafters-io/http-server-starter-go/internal/server"
	"github.com/k0kubun/pp"
)

func main() {
	allowedPaths := []string{
		"/lol",
		"/lol1",
	}
	s, err := server.NewServer("localhost", 9779, allowedPaths)
	if err != nil {
		pp.Println("OH SHITTTTT")
	}
	err = s.Run()
	if err != nil {
		pp.Println(err)
	}
}
