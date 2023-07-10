package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan struct{})
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:9090", Path: "/ws"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	// Start a goroutine to receive messages from the WebSocket server
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message from server:", err)
				return
			}
			fmt.Printf("Received message: %s\n", message)
		}
	}()

	// Start a goroutine to send messages to the WebSocket server
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				message := "Hello from the client!"
				err := conn.WriteMessage(websocket.TextMessage, []byte(message))
				if err != nil {
					log.Println("Error sending message to server:", err)
					return
				}
				fmt.Printf("Sent message: %s\n", message)
			case <-interrupt:
				// Gracefully close the connection when interrupted
				err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("Error closing connection:", err)
					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return
			}
		}
	}()

	// Wait for an interrupt signal (e.g., Ctrl+C) to stop the client
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt)
	<-interruptSignal

	log.Println("Stopping the client...")
	close(interrupt)

	// Wait for the goroutines to finish
	select {
	case <-done:
	case <-time.After(time.Second):
		log.Println("Timed out waiting for goroutines to finish")
	}
}
