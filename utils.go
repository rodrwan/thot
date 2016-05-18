package main

import (
	"log"
	"os"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
