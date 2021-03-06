package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/jhunderdog/nmcoin/explorer"
	"github.com/jhunderdog/nmcoin/rest"
)

func usage() {
	fmt.Printf("Welocme to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port:	Start the PORT of the server\n")
	fmt.Printf("-mode:	Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
