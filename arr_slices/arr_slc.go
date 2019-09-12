package main

import "fmt"

func main() {
	//var names [3]string // array of length three
	//names[0] = "jim"
	//names[1] = "jon"
	//names[2] = "jan"
	//fmt.Println(names)

	names := make([]string, 4)
	names = append(names, "don")
	names = append(names, "jon")
	names = append(names, "ron")
	names[3] = "jin"

	fmt.Println(names)
	fmt.Println("names[2] is nil:", names[2] == "")
}
