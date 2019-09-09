package main

import "fmt"

func main() {
    // buffered channels accept a limited number of values without a corresponding reciever for those values
    messages := make(chan string, 2) // make a channel of strings buffering up to 2 values

    messages <- "buffered" // because it's buffered we can send the two values into the channel
    messages <- "channel" // without a corresponding concurrent recieve

    fmt.Println(<-messages)
    fmt.Println(<-messages)


}
