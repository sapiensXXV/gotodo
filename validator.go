package main

import (
	"log"
)

func Validator(args []string) {
	validateNumberOfArgs(args)
	cmd := args[0]

	switch cmd {
	case "add":
		validateAdd(args)
	case "remove":
		validateRemove(args)
	case "list":
		validateList(args)
	case "status":
		validateStatus(args)
	default:
		log.Fatalln("허용되지 않은 명령어 입니다. 사용가능한 명령어: add, remove, lise")
	}
}

func validateNumberOfArgs(args []string) {
	if len(args) <= 0 {
		log.Fatalln("Usage: gotodo <command> [arguments]")
	}
}

func validateAdd(args []string) {
	if len(args) < 3 {
		log.Fatalln("usage: gotodo add <title> <category>")
	}
}

func validateRemove(args []string) {
	if len(args) < 2 {
		log.Fatalln("usage: gotodo remove <task_number>")
	}
}

func validateList(args []string) {
	if len(args) < 1 {
		log.Fatalln("usage: gotodo list")
	}
}

func validateStatus(args []string) {
	if len(args) < 3 {
		log.Fatalln("usage: gotodo status <task_number> <new_status>")
	}
}
