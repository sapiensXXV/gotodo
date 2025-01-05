package main

import (
	"log"
	"os"
)

func main() {

	args := os.Args[1:]

	Validator(args)
	manager := NewTaskManager()
	if len(args) == 0 {
		log.Fatalln("Usage: gotodo <command> [arguments]")
	}
	manager.handleCommand(args)

}
