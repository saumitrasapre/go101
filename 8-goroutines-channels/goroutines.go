/*
 * Further Reading:
 *
 * Goroutines - https://go101.org/article/control-flows-more.html
 * Channels - https://go101.org/article/channel.html
 */

package main

import (
	"fmt"
	"net"
	"os"
)

/* CONSTANTS */

const (
	Proto          = "tcp"       // transport protocol to be used
	Host           = "localhost" // server host name
	Port           = "8080"      // server port number
	WorldSize      = 7           // number of clients that we expect to receive connections from
	MaxMessageSize = 100         // the maximum size of a message we expect to receive from any client
)

/* FUNCTION DEFINITIONS */

func listenForClientConnections(
	write_only_ch chan<- string, // a channel to which client messages will be sent
) {
	server_address := Host + ":" + Port
	fmt.Println("Starting " + Proto + " server on " + server_address)

	listener, err := net.Listen(Proto, server_address)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connecting: ", err.Error())
			return
		}
		fmt.Println("Client " + conn.RemoteAddr().String() + " connected, spawning goroutine to handle connection")

		// Spawn a new goroutine to handle this client's connections
		// and go back to listening for additional connections
		go handleClientConnection(conn, write_only_ch)
	}
}

func handleClientConnection(
	conn net.Conn, // the connection to handle
	write_only_ch chan<- string, // a channel to which incoming messages from this client will be written
) {
	/*
	 * In this simple example, we issue a single conn.Read() call without accounting
	 * for the fact that the client may have more data to send. For the assignment,
	 * please implement code that reads bytes until a record with stream_complete=1
	 * has been received.
	 */

	client_msg_buf := make([]byte, MaxMessageSize)
	bytes_read, err := conn.Read(client_msg_buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Conn::Read: err %v\n", err)
		os.Exit(1)
	}
	message := string(client_msg_buf[0:bytes_read])
	write_only_ch <- message
	conn.Close()
}

func receiveDataFromClients(
	read_only_ch <-chan string, // a channel from client messages will be read
) []string {
	/*
	 * Here, we wait for exactly 'WorldSize' messages to be received
	 * by performing a blocking read on the channel 'read_only_ch'.
	 *
	 * Note that since this function was called synchronously from main()
	 * i.e. not within a separate goroutine, the main goroutine will wait
	 * till all messages have been received before exiting. Please make
	 * sure that in your implementation, the main goroutine waits for
	 * the completion of all other goroutines it had spawned before
	 * exiting.
	 */
	var client_messages []string
	messages_received := 0
	for messages_received < WorldSize {
		/*
		 * Since we're using an unbuffered channel, the receive operation
		 * will block till some other goroutine writes to this channel
		 */
		message := <-read_only_ch
		fmt.Println("Received message", message, ", appending to slice")
		client_messages = append(client_messages, message)
		messages_received++
	}

	return client_messages
}

func main() {

	/*
	 * Create a bi-directional channel whose:
	 * 		a) send-side will be used by the goroutine that handles
	 * 		each client's connection to send the client's message
	 *
	 *		b) receive-side will be used by the main goroutine to
	 *		accumulate all client messages into a slice
	 */
	bidirectional_ch := make(chan string)
	defer close(bidirectional_ch)

	/*
	 * Spawn a new goroutine that concurrently listens for connections from
	 * clients and pass the send side of this channel to this go-routine.
	 *
	 * Though the channel is bidirectional, the function's parameter
	 * is of type "send-only", as a result of which, it is converted
	 * into a "send-only" channel when the function is invoked
	 */
	go listenForClientConnections(bidirectional_ch)

	/*
	 * Wait until all client messages have been received.
	 *
	 * Though the channel is bidirectional, the function's parameter
	 * is of type "receive-only", as a result of which, it is converted
	 * into a "receive-only" channel when the function is invoked
	 */

	client_messages := receiveDataFromClients(bidirectional_ch)

	// Print client messages
	fmt.Println("\nClient messages:")
	for slice_index, message := range client_messages {
		fmt.Println(slice_index, message)
	}

}