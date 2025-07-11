package command

import "errors"

func CommandHandler(command *Command) ([]byte, error) {
	switch command.Name {
	case "ECHO":
		return []byte(command.Args[0]), nil
	}
	return nil, errors.New("??")
}