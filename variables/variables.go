package main

import "fmt"

func main(){
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)

		var myInt int = 16
		var val, ok = "yes", true

		fmt.Println("myInt is:", myInt)
		fmt.Println("myInt times two:", myInt*2)
		fmt.Println("val is:", val)
		fmt.Println("ok is:" ok)
}
