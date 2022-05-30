package main

import (
    "fmt"

    "github.com/wmy09527/mytool/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("mytool")
    fmt.Println(message)
}
