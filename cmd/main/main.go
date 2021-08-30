package main

import (
	"fmt"

	"github.com/ark-go/httpfiber/internal/server"
)

var versionProg string

func init() {
}

func main() {
	fmt.Println("Version:", versionProg)
	server.StartServer()
}
