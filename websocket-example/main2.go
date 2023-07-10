package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePagefunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func wsfunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "hello from websocket")

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, error := upgrader.Upgrade(w, r, nil)
	if error != nil {
		log.Println(error)
	}
	error = ws.WriteMessage(1, []byte("hello client"))
	if error != nil {
		log.Println(error)
	}

	reader(ws)

}

func reader(ws *websocket.Conn) {
	for {
		messageType, packate, err := ws.ReadMessage()

		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(packate))

		if err := ws.WriteMessage(messageType, []byte("Hey")); err != nil {
			log.Println(err)
		}

	}
}

func main() {
	http.HandleFunc("/", homePagefunc)
	http.HandleFunc("/ws", wsfunc)

	http.ListenAndServe(":9090", nil)
}
