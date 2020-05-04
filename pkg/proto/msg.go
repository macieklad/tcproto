package proto

import "fmt"


func Error(msg string) []byte {
	return []byte(fmt.Sprintf("ERR %s \n", msg))
}

func Ok() []byte {
	return []byte("OK\n")
}