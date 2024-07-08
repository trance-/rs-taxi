package main

import (
    "fmt"
    "os"
    "syscall"
    "os/signal"
)

func controlLoop(ch chan string) {
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGTERM)

    for {
        select {
        case data := <-ch:
            fmt.Println("Received from control loop:", data)
            // Perform actions based on received data
        case <-done:
            fmt.Println("Received termination signal, exiting...")
            // Perform cleanup tasks before exiting
            return
        }
    }
}

func main() {
    ch := make(chan string)
    go controlLoop(ch)
    // ... rest of your application code
}

