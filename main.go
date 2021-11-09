package main

import (
	"github.com/jhunderdog/nmcoin/explorer"
	"github.com/jhunderdog/nmcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(5000)

}
