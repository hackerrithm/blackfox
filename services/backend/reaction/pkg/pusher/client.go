// Copyright 2019 kemar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pusher

import "github.com/gorilla/websocket"

type Client struct {
	hub      *Hub
	id       int
	socket   *websocket.Conn
	outbound chan []byte
}

func newClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (client *Client) write() {
	for {
		select {
		case data, ok := <-client.outbound:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (client Client) close() {
	client.socket.Close()
	close(client.outbound)
}
