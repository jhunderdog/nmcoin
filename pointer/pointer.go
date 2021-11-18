package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	fmt.Printf("funcinside = %d\n", arg.value)
	arg.value = 999
	fmt.Printf("funcinside = %d\n", arg.value)
	fmt.Print(arg.value)
	arg.data[100] = 999

}

func main() {
	var data Data
	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
	// fmt.Print(data)
}
