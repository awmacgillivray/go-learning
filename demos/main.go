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
			case "6":
				six()
			case "7":
				seven()
			case "8":
				eight()
			case "9":
				nine()
			case "10":
				ten()
			case "11":
				eleven()
			case "12":
				twelve()
		}
	}
}
