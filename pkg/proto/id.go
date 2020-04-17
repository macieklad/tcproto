package proto

type ID int

const (
	REG ID = iota
	JOIN
	LEAVE
	MSG
	CHNS
	USRS
)
