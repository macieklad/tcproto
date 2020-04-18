package proto

import "net"

type Hub struct {
	channels        map[string]*channel
	clients         map[string]*client
	commands        chan command
	deregistrations chan *client
	registrations   chan *client
}

func NewHub() *Hub {
	return &Hub{
		registrations:   make(chan *client),
		deregistrations: make(chan *client),
		clients:         make(map[string]*client),
		channels:        make(map[string]*channel),
		commands:        make(chan command),
	}
}

func (h *Hub) MakeClient(conn net.Conn) *client {
	return NewClient(
		conn,
		h.commands,
		h.registrations,
		h.deregistrations,
	)
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.registrations:
			h.register(client)
		case client := <-h.deregistrations:
			h.unregister(client)
		case cmd := <-h.commands:
			switch cmd.id {
			case JOIN:
				h.joinChannel(cmd.sender, cmd.recipient)
			case LEAVE:
				h.leaveChannel(cmd.sender, cmd.recipient)
			case MSG:
				h.message(cmd.sender, cmd.recipient, cmd.body)
			case USRS:
				h.listUsers(cmd.sender)
			case CHNS:
				h.listChannels(cmd.sender)
			default:
				panic("Unknown command passed, code " + string(cmd.id))
			}
		}
	}
}

func (h *Hub) register(c *client) {
	if _, exists := h.clients[c.username]; exists {
		c.username = ""
		c.conn.Write([]byte("ERR username taken \n"))
	} else {
		h.clients[c.username] = c
		c.conn.Write([]byte("OK\n"))
	}

}

func (h *Hub) unregister(c *client) {
	if _, exists := h.clients[c.username]; exists {
		delete(h.clients, c.username)

		for _, channel := range h.channels {
			delete(channel.clients, c)
		}
	}
}

func (h *Hub) joinChannel(u string, c string) {
	if client, ok := h.clients[u]; ok {
		if channel, ok := h.channels[c]; ok {
			channel.clients[client] = true
		} else {
			h.channels[c] = newChannel(c)
			h.channels[c].clients[client] = true
		}
	}
}

func (h *Hub) leaveChannel(sender string, recipient string) {

}

func (h *Hub) message(u string, r string, m []byte) {
	if sender, ok := h.clients[u]; ok {
		switch r[0] {
		case '#':
			if channel, ok := h.channels[r]; ok {
				if _, ok := channel.clients[sender]; ok {
					channel.broadcast(sender.username, m)
				}
			}
		case '@':
			if user, ok := h.clients[r]; ok {
				user.conn.Write(append(m, '\n'))
			}
		}
	}
}

func (h *Hub) listUsers(sender string) {

}

func (h *Hub) listChannels(sender string) {

}
