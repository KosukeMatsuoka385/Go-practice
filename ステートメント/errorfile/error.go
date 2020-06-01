package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("./range.go")
	if err != nil {
		log.Fatal("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
}