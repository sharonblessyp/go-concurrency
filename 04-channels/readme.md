## Channels
Used for communication between goroutines.

-------------------------------
**Important rule**
Channels are blocking by default

The Go philosophy: "Do not communicate by sharing memory; instead, share memory by communicating."

------------------------------
## Types
**A. Unbuffered Channels (The Default)**
These have no capacity to hold data. They require both a sender and a receiver to be ready at the exact same time.

Behavior: The sender blocks until a receiver takes the value. The receiver blocks until a sender sends a value.


**B. Buffered Channels**
These have a pre-defined capacity.

Behavior: The sender only blocks when the "buffer" is full. The receiver only blocks when the buffer is empty.

-------------------------
**Directional Channels**

- Bidirectional	**chan Type**	Can send and receive.
- Send-only(Producer) **chan<- Type**	
- Receive-only (Consumer)	**<-chan Type**	Can 

---------------------------
**Closing Channels**

- Closing a channel signals that no more values will be sent.
- **Who closes?** Always the sender. 
- Check status: val, ok := <-ch. If ok is **false**, the channel is **closed and empty.**

----------------------
**The select Statement**

- The select statement is like a switch for channels. 
- It lets a goroutine wait on multiple channel operations. 
- It picks whichever one is "ready" first.

-------------------
The "Idiomatic" Way **(range)**
The loop internally does the val, ok := <-ch check 

---------------------
**Note**
A range loop over a channel will block if the channel is open but empty. It only terminates when the channel is explicitly closed by a sender.