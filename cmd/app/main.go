package main

import (
	"github.com/marcoshuck/dagger-demo/pkg/adder"
	"log"
)

func main() {
	log.Println("Hello, world!")
	log.Println(adder.Add(500, 1))
}
