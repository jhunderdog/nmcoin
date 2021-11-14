package main

import (
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
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorere")
	case "rest":
		fmt.Println("Start Rest API")
	default:
		usage()
	}

}
