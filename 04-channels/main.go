package main

func main() {
	// A. unbuffered
	ch := make(chan string)

	go func() {
		// blocks here until main receives
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	println(msg)

	// B. buffered channels
	bufCh := make(chan string, 2)
	bufCh <- "First"
	bufCh <- "Second"
	// sender should close the channel
	close(bufCh)

	for msg := range bufCh {
		println(msg)
	}
}
