package proto

type command struct {
	id        ID
	recipient string
	sender    string
	body      []byte
}
