package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// EnqueuePayload represents enqueue payload data.
type EnqueuePayload struct {
	ItemIDs   []string `json:"itemIDs"`
	PodcastID string   `json:"podcastID"`
	TagIDs    []string `json:"tagIDs"`
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	activePlayers  = make(map[*websocket.Conn]string)
	allConnections = make(map[*websocket.Conn]string)
	connMutex      sync.RWMutex
)

var broadcast = make(chan Message) // broadcast channel

// Message represents message data.
type Message struct {
	Connection  *websocket.Conn `json:"-"`
	Identifier  string          `json:"identifier"`
	MessageType string          `json:"messageType"`
	Payload     string          `json:"payload"`
}

// Wshandler handles the wshandler request.
func Wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("Error closing websocket connection: %v\n", err)
		}
	}()
	for {
		var mess Message
		err := conn.ReadJSON(&mess)
		if err != nil {
			connMutex.Lock()
			isPlayer := activePlayers[conn] != ""
			if isPlayer {
				delete(activePlayers, conn)
				broadcast <- Message{
					MessageType: "PlayerRemoved",
					Identifier:  mess.Identifier,
				}
			}
			delete(allConnections, conn)
			connMutex.Unlock()
			break
		}
		mess.Connection = conn
		connMutex.Lock()
		allConnections[conn] = mess.Identifier
		connMutex.Unlock()
		broadcast <- mess
	}
}

// HandleWebsocketMessages handles the handle websocket messages request.
func HandleWebsocketMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		switch msg.MessageType {
		case "RegisterPlayer":
			connMutex.Lock()
			activePlayers[msg.Connection] = msg.Identifier
			connMutex.Unlock()

			connMutex.RLock()
			for connection := range allConnections {
				if err := connection.WriteJSON(Message{
					Identifier:  msg.Identifier,
					MessageType: "PlayerExists",
				}); err != nil {
					fmt.Printf("Error writing JSON to connection: %v\n", err)
				}
			}
			connMutex.RUnlock()
			fmt.Println("Player Registered")
		case "PlayerRemoved":
			connMutex.RLock()
			for connection := range allConnections {
				if err := connection.WriteJSON(Message{
					Identifier:  msg.Identifier,
					MessageType: "NoPlayer",
				}); err != nil {
					fmt.Printf("Error writing JSON to connection: %v\n", err)
				}
			}
			connMutex.RUnlock()
			fmt.Println("Player Registered")
		case "Enqueue":
			var payload EnqueuePayload
			fmt.Println(msg.Payload)
			err := json.Unmarshal([]byte(msg.Payload), &payload)
			if err == nil {
				items := getItemsToPlay(payload.ItemIDs, payload.PodcastID, payload.TagIDs)
				var player *websocket.Conn
				connMutex.RLock()
				for connection, id := range activePlayers {
					if msg.Identifier == id {
						player = connection
						break
					}
				}
				connMutex.RUnlock()
				if player != nil {
					payloadStr, marshalErr := json.Marshal(items)
					if marshalErr == nil {
						if writeErr := player.WriteJSON(Message{
							Identifier:  msg.Identifier,
							MessageType: "Enqueue",
							Payload:     string(payloadStr),
						}); writeErr != nil {
							fmt.Printf("Error writing JSON to connection: %v\n", writeErr)
						}
					}
				}
			} else {
				fmt.Println(err.Error())
			}
		case "Register":
			var player *websocket.Conn
			connMutex.RLock()
			for connection, id := range activePlayers {
				if msg.Identifier == id {
					player = connection
					break
				}
			}
			connMutex.RUnlock()

			if player == nil {
				fmt.Println("Player Not Exists")
				if err := msg.Connection.WriteJSON(Message{
					Identifier:  msg.Identifier,
					MessageType: "NoPlayer",
				}); err != nil {
					fmt.Printf("Error writing JSON to connection: %v\n", err)
				}
			} else {
				if err := msg.Connection.WriteJSON(Message{
					Identifier:  msg.Identifier,
					MessageType: "PlayerExists",
				}); err != nil {
					fmt.Printf("Error writing JSON to connection: %v\n", err)
				}
			}
		}
	}
}
