package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("server")
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer func() {
			log.Println("disconnect")
			connection.Close()
		}()
		for {
			messageType, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				break
			}
			log.Printf("receive: %s\n", message)
			err = connection.WriteMessage(messageType, message)
			if err != nil {
				log.Println("write: ", err)
				break
			}
		}
	})
	log.Println("server start at : 8899")
	log.Fatal(http.ListenAndServe(":8888", nil))

}
