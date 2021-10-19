package main

import (
	"reserver/pkg"
	"reserver/utils"
)

func main() {
	server := pkg.NewReServer(utils.ParseCommonArgs())
	server.Run()
}
