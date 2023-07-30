# Go Routines
A goroutine is a lightweight thread of execution.

A goroutine will always execute concurrently with respect to the calling function. We can also start goroutine for an anonymous function call.

In go, a program starts with one goroutine, which executes the `main` function. Additional goroutines can be created using the `go` keyword followed by a function call. This starts a new goroutine that runs concurrently with the original goroutine.

The go keyword makes the function call to return immediately, while the function starts running in the background as a goroutine and the rest of the program continues its execution.

> Goroutines are very lightweight, and it's possible to create thousands or even millions of them in a single program without significant overhead. This makes it easy to write concurrent programs in Go that take advantage of multiple CPU cores and can perform many tasks simultaneously.

goroutines are managed by the Go runtime, they are automatically scheduled and can communicate with each other using `channels`. This makes it easy to write complex concurrent programs without worrying about low-level details such as locking and synchronization.

## WaitGroups

In Go, `WaitGroup` is a synchronization primitive provided by the `sync package`. It is used to wait for a collection of goroutines to finish their execution before proceeding further in the program.

The `WaitGroup` has 3 essential methods : 

1. `Add(delta int)`: Adds or subtracts the number of goroutines to wait for. The `delta` parameter can be positive or negative. Typically we add the number of goroutines we are about to start, then for each goroutine we call `Done()` when it finishes it's execution.
2. `Done()`: Decrements the `WaitGroup` counter by one. This is the usually called at the end of each goroutine's execution.
3. `Wait()`: Blocks the execution of the current goroutine until the WaitGroup counter becomes zero.
