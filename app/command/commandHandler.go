package command

import (
	"errors"
	"strconv"
)

func CommandHandler(command *Command) ([]byte, error) {
	switch command.Name {
	case "ECHO":
		s := "$" + strconv.Itoa(len(command.Args[0])) + "\r\n" + command.Args[0] + "\r\n"
		return []byte(s), nil
	case "PING":
		s := "$4\r\nPONG\r\n"
		return []byte(s), nil
	}
	return nil, errors.New("??")
}
