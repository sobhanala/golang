# Go Learning Exercises

A small repository for learning Go. The current source is `hay.go`, a goroutine and channel example.

## Run

Install Go, then run:

```sh
go run hay.go
```

The program starts two goroutines and prints `ping` after receiving from the first channel. The second goroutine sends to an unconsumed channel and stays blocked until the process exits; this is a learning example rather than a complete concurrency pattern.

There is no Go module or automated test suite in this snapshot.
