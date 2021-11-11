package main

import (
	"log"

	"github.com/HuCuiGang/bank/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
