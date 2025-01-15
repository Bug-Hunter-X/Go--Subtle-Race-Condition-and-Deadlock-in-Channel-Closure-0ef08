# Go: Subtle Race Condition and Deadlock

This repository demonstrates a subtle race condition and potential deadlock in a Go program involving channels and goroutines. The program appears simple but highlights the importance of careful synchronization when dealing with concurrent operations.

## Bug Description

A goroutine waits to receive from a channel, but the channel is closed before the goroutine has a chance to receive anything. This leads to a deadlock, halting program execution. The subtlety comes from the seemingly correct use of a WaitGroup, which masks the underlying race condition.