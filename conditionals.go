package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["jim"] = 61
	if ages["jim"] < 18 {
		fmt.Println("jim can't vote yet")
	} else if ages["jim"] < 67 {
		fmt.Println("jim is not ready for retirement")
	} else {
		fmt.Println("enjoy retirement, jim!")
	}

	switch ages["jim"] {
	case 1, 2, 3, 4, 5, 11, 13, 17, 19:
		fmt.Println("kevin's age is a prime number")
	case 16:
		fmt.Println("jim can drive")
	case 18:
		fmt.Println("jim can vote")
	default:
		fmt.Println("jim can retire")
	}
}
