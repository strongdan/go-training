package main

import "fmt"
import "math"

const s string = "constant"

func main(){
    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}
