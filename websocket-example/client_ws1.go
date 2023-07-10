package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := url.URL{Scheme: "ws", Host: "127.0.0.1:9090", Path: "/ws"}

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		log.Println(err)
	}

	interrupt := make(chan struct{})
	done := make(chan bool)

	defer conn.Close()
	//reader
	go func() {
		for {
			_, packate, err := conn.ReadMessage()

			if err != nil {
				log.Println(err)
			}
			log.Println(string(packate))
		}

	}()

	// writer
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				message := "hellow from client!!"
				if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("error sending message:", err)
					return
				}

				log.Println("write message: ", message)
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
