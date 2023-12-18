package main

import (
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	receiver := NewDataReceiver()
	http.HandleFunc("/ws", receiver.handleWS)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("new obu connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Printf("read error: %v\n", err)
			continue
		}

		fmt.Printf("received data from obu [%d]: <lat %.2f, long %.2f>\n", data.OBUID, data.Latitude, data.Longitude)

		dr.msgch <- data
	}
}
