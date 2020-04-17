package proto

type channel struct {
	name    string
	clients map[*client]bool
}

func (c *channel) broadcast(s string, m []byte) {
	msg := append([]byte(s), ": "...)
	msg = append(msg, m...)
	msg = append(msg, '\n')

	for cl := range c.clients {
		cl.conn.Write(msg)
	}
}

func newChannel(name string) *channel {
	c := new(channel)
	c.name = name
	return c
}
