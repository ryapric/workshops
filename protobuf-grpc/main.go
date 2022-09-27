package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	args := flag.CommandLine.Args()
	if len(args) == 0 {
		log.Fatalln("You must provide at least one arg, and you provided 0")
	}

	if args[0] == "server" {
		runServer()
	} else {
		if len(args) > 1 {
			runClient(args[1])
		} else {
			log.Fatalln("You must provide a URI arg")
		}
	}
}
