package command

type Command struct {
	Name string
	Args []string
}

type CommandType int

const (
	PING CommandType = iota
	ECHO
)