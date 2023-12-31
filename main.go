// Connect multiple client to single websocket
// Boardcast one client message to all and recieve
// Build front end with chat sceeen, input and send button80o;lo
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize:  1024,
}

var wsConns []*websocket.Conn
var history [][]byte;

func reader(conn *websocket.Conn){
	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		history = append(history, p)

		for _, val := range wsConns {
			// if val == conn{
			// 	continue
			// }
			if err := val.WriteMessage(messageType, p); err != nil{
				log.Println(err)
				return
			}	
		}
		// if err := conn.WriteMessage(messageType, p); err != nil{
		// 	log.Println(err)
		// 	return
		// }
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func (r *http.Request )bool {
		return true;
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}
	wsConns = append(wsConns, ws);
	
	log.Println("Client sucessfully connected...")
	reader(ws)
}

func setupRoutes(){
	http.Handle("/", http.FileServer(http.Dir("./static")));
	http.HandleFunc("/ws", wsEndpoint);
}

func main(){
	fmt.Println("Starting connection at port 8080")
	setupRoutes();
	http.ListenAndServe("localhost:8080", nil);
}