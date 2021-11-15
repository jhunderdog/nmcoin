package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welocme to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTMl Explorer\n")
	fmt.Printf("rest:	Starts the REST API(recommended)\n\n")
	os.Exit(0)
}

func main() {
	// go explorer.Start(3000)
	// rest.Start(4000)
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorere")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}
	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start server")
	}
}
