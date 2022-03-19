package main

import (
    "fmt"
    "time"
)
func main() {
    now := time.Now()

    channel1 := make(chan string)
    go func() {channel1 <- "Heello"}()

    fmt.Println(<-channel1)
    fmt.Println(time.Now().Sub(now))

}
