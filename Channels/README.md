# Channels

Channels are pipes that connect concurrently running goroutines. Channels can be used for bi directional communication between goroutines.

Bidirectional communication means, using only one channel we are able to **send** and **recieve** data. Code execution will be stopped until the write and read operations are done successfully. This allows goroutines to synchronize without explicit locks or condition variables.

## Different types of channels

1.  **Unbuffered Channels**: These require both the sender and receiver to be present to be a successful operation. By default channels are unbuffered.
2.  **Buffered Channels**: These channels have the capacity to store the vaules for further processing, until the reader goroutine is available to perform operations on the values sent by the sender. The sender is not blocked until it becomes full and it doesn't necessarily need a reader to complete the synchronization with every operation.
    
    If a space in the array is available, the sender can send its value to the channel and complete its send operation immediately.

    After its execution, if a receiver comes, the channel will start sending values to the receiver and it will start its operation once it receives the values. As the sender and receiver are operating at different times, this is called *asynchronous communication*.



> Syntax to create a channel`make(chan val-type).`Channels are typed by the values they convey.
    
> Declaration of channels based on directions: </br>
    1. Bidirectional channel : chan T </br>
    2. Send only channel: chan <- T </br>
    3. Receive only channel: <- chan T