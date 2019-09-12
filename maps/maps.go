package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["K2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	birthdays := map[string]string{
		"jim": "02/06/1990",
		"jon": "01/01/1957",
		"jin": "06/24/1975",
	}

	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["jim"] = 28 // adding to maps
	ages["jon"] = 61
	ages["jin"] = 43

	fmt.Println(ages, ages["jon"])

	delete(ages, "jon") // removing elements from maps
}
