```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 1) // Buffered channel to avoid deadlock

        wg.Add(1)
        go func() {
                defer wg.Done()
                fmt.Println("Goroutine 1: ", <-ch)
        }()

        ch <- 1 // Send data to the channel before closing
        close(ch)
        wg.Wait()
}
```