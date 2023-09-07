package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jimbo8702/toll-calc/types"
	"github.com/gorilla/websocket"
)

// var kafkaTopic = "obudata"

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":3000", nil)
}

type DataReceiver struct {
	msgch 	chan types.OBUData
	conn	*websocket.Conn
	prod 	DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p DataProducer
		err error
		kafkaTopic = "obudata"
	)
	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}
	p = NewLogMiddleware(p)
	return &DataReceiver{
		msgch: 	make(chan types.OBUData, 128),
		prod: 	p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}
//gets from the web socket and produces it to kafka

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU cleint connected !")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}
		if err := dr.produceData(data); err != nil {
			log.Println("kafka produce error:", err)
			continue
		}
	}
}