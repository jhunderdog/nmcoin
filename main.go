package main

import (
	"fmt"

	"github.com/jhunderdog/nmcoin/person"
)

func main() {
	jongho := person.Person{}
	jongho.SetDetails("jongho", 34)
	nico := person.Person{}
	nico.SetDetails("nico", 29)
	fmt.Println("Main 'jongho'", jongho)
	fmt.Println("Main 'nico'", nico)
}
