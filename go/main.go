package main

import (
	"go-api/log"
	"go-api/server"
)

func main() {

	infoChan, errChan := log.Start()

	server.Start(infoChan, errChan) // blocks execution

}
