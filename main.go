package main

import (
	"flag"

	"github.com/axseem/blomple/cmd"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "ssg":
		cmd.ServeSSG()
	case "ssr":
		cmd.ServeSSG()
	case "seed":
		cmd.SeedDB()
	case "help":
		cmd.Help()
	default:
		cmd.Help()
	}
}
