package main

import "github.com/jhunderdog/nmcoin/cli"

func main() {
	// go explorer.Start(3000)
	// rest.Start(4000)
	// fmt.Println(os.Args)
	// if len(os.Args) < 2 {
	// 	usage()
	// }

	// rest := flag.NewFlagSet("rest", flag.ExitOnError)
	// portFlag := rest.Int("port", 4000, "Sets the port of the server")

	// switch os.Args[1] {
	// case "explorer":
	// 	fmt.Println("Start Explorere")
	// case "rest":
	// 	rest.Parse(os.Args[2:])
	// default:
	// 	usage()
	// }
	// if rest.Parsed() {
	// 	fmt.Println(portFlag)
	// 	fmt.Println("Start server")
	// }
	cli.Start()
}
