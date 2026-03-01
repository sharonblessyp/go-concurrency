## Why Worker Pool?
- say we have 10,000 tasks and we don’t want 10,000 goroutines running at once.
- use when controlled parallelism is needed.
- Fixed number of workers, Processing unlimited tasks

--------------

**Problems:**

- Too many goroutines
- Memory pressure
- CPU thrashing
- DB connection overload
- API rate limits

----------------------
**Components:**

- Job queue (channel)
- N workers (goroutines)
- WaitGroup for shutdown

-------------
**Note**
Always pass sync primitives as pointers: it holds state

*sync.WaitGroup
*sync.Mutex
*sync.RWMutex
*sync.Once