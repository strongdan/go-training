package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) { // done channel will notify another goroutine 
    fmt.Println("working...") // that this function's work is done
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true // send a value to notify that we're done
}

func main() {
    done := make(chan bool, 1) // start a worker goroutine, giving it the channel to notify on
    go worker(done)

    <-done // block until we recieve a notificatino from the worker on the channel
}
