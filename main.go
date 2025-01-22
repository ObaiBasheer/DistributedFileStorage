package main

import (
	"fmt"
	"log"

	"github.com/ObaiBasheer/Dfs/p2p"
)

func main() {
	// Code
	fmt.Println("Starting ...")

	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
