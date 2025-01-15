```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)

        wg.Add(1)
        go func() {
                defer wg.Done()
                fmt.Println("Goroutine 1: ", <-ch)
        }()

        close(ch)
        wg.Wait()
}
```