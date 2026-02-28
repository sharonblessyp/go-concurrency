The WaitGroup in the Go programming language is a synchronization primitive from the standard **sync package** used to wait for a collection of goroutines to finish their execution before the main program proceeds

The sync.WaitGroup exports three key methods:
- Add(delta int)
- Done()
- Wait()