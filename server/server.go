package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // */*
}

func rootfunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Success", "Desc": "RootPage Endpoint Hitted"})

}

func readAndEcho(ws *websocket.Conn) {
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("The Received Message type is %v \n", msgType)
		log.Printf("The Received Message is %v \n", msg)
		log.Println("Echo the same messgae")
		// echo the same msg
		if err := ws.WriteMessage(msgType, []byte(fmt.Sprintf("The Echo : %s", msg))); err != nil {
			log.Println(err)
		}
		// type ResJson struct {
		// 	Msg string
		// }
		// ws.WriteJSON(ResJson{Msg: "Message"})

	}
}

func writeMsg(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second * 10)
		log.Println("Writing the message to the client")
		msg := time.Now()
		log.Println(msg.GoString())
		if err := ws.WriteMessage(1, []byte(msg.GoString())); err != nil {
			log.Println(err)
			return
		}
	}
}

func getWSConn(w http.ResponseWriter, req *http.Request) {
	log.Println("Received new Client connection")
	ws, err := wsUpgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client upgrade to WS success")
	go writeMsg(ws)
	go readAndEcho(ws)

}

func main() {
	http.HandleFunc("/", rootfunc)
	http.HandleFunc("/ws", getWSConn)

	log.Fatal(http.ListenAndServe(":8090", nil))

}
