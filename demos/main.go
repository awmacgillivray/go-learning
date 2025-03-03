package main

import (
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for _, element := range argsWithoutProg {
		switch element {
			case "1":
				one()
			case "2":
				two()
			case "3":
				three()
			case "4":
				four()
			case "5":
				five()
		}
	}
}
