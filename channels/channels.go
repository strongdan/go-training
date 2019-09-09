package main

import "fmt"

func main() {
    messages := make(chan string) // create a new channel

    go func() { messages <- "ping" }() // send a value into the channel with <-

    msg := <-messages // the <-channel syntax recieves a value from the channel
    fmt.Println(msg)

}
