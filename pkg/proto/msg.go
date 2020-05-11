package proto

import "fmt"

func Error(msg string) []byte {
	return []byte(fmt.Sprintf("ERR %s \n", msg))
}

func Ok() []byte {
	return []byte("OK\n")
}

func OkMsg(msg string) []byte {
	return []byte(fmt.Sprintf("OK %s \n", msg))
}
