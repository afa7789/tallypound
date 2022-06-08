package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/afa7789/tallypound/cmd"
	"github.com/afa7789/tallypound/internal/domain"
)

var flags domain.Flags

// init of the project, it loads a port from the command line input
func init() {
	flags.Port = flag.Int("port", 8080, "port number to listen")
}

// main function
func main() {
	flag.Parse()
	if err := cmd.ServerExecute(flags); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("")
		log.Printf("SERVER HERE: http://localhost:%d\n", *flags.Port)
	}
}
