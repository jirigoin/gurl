package main

import (
	"log"
	"os"
)

var missingUrlErr = "You must pass a url value!"

func main() {

	if !(len(os.Args) > 1) {
		log.Fatalln(missingUrlErr)
	}
	url := os.Args[1]

	handler := HandlerFactory()
	if err := handler.Store(url); err != nil {
		log.Fatalln("can't download the file: ", err)
	}
}
