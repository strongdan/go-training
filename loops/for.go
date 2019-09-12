package main

import "fmt"

func main() {

	ages := map[string]int{}
	ages["Kevin"] = 11
	ages["Keith"] = 28
	ages["James"] = 67
	ages["Michael"] = 18
	ages["Leigha"] = 16

	for i := 1; i <= 10; i++ {
		fmt.Println("we're counting:", i)
	}
	a := 0
	for a < 10 { // while loop
		if a%2 == 0 {
			a++
			continue
		} else if a == 5 {
			break
		}
		fmt.Println("we're counting (again):", a)
		a++
	}
	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(fmt.Sprintf("%s's age is a prime number!", name))
		case 16:
			fmt.Println(name, "can drive now!")
		case 18:
			fmt.Println(name, "can vote now!")
		case 67:
			fmt.Println(name, "can retire now!")
		default:
			fmt.Println(fmt.Sprintf("There's nothing special about %s's current age.", name))
		}
	}
}
