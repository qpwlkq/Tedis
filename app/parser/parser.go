package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/command"
)

func Parse(b []byte) (*command.Command, error) {
	if len(b) <= 0 {
		return nil, errors.New("empty input")
	}
	result, err := ParseRESP(string(b))
	command := command.Command{}
	command.Name = strings.ToUpper(result[0])
	command.Args = result[1:]
	return &command, err
}

/*
RESP data type	Minimal protocol version	Category	First byte
Simple strings	RESP2	Simple	+
Simple Errors	RESP2	Simple	-
Integers	RESP2	Simple	:
Bulk strings	RESP2	Aggregate	$
Arrays	RESP2	Aggregate  	*
Nulls	RESP3	Simple	_
Booleans	RESP3	Simple	#
Doubles	RESP3	Simple	,
Big numbers	RESP3	Simple	(
Bulk errors	RESP3	Aggregate	!
Verbatim strings	RESP3	Aggregate	=
Maps	RESP3	Aggregate	%
Attributes	RESP3	Aggregate	|
Sets	RESP3	Aggregate	~
Pushes	RESP3	Aggregate	>
*/
func ParseRESP(input string) ([]string, error) {
	if len(input) <= 0 {
		return nil, errors.New("empty input")
	}
	firstByte := input[0]
	switch firstByte {
	case '+':
	case '-':
	case '$':
		return ParseBulkString(input)
	case '*':
		return ParseArrays(input)
	case '_':
	case '#':
	case ',':
	case '(':
	case '!':
	case '=':
	case '%':
	case '|':
	case '~':
	case '>':
	default:
		return nil, errors.New("")
	}
	return nil, errors.New("Parse() is under development")
}

// *<number-of-elements>\r\n<element-1>...<element-n>
func ParseArrays(input string) ([]string, error) {
	splits := strings.Split(input, "\r\n")
	num, err := strconv.Atoi(splits[0][1:])
	if err != nil {
		return nil, errors.New("parse failed")
	}
	fmt.Print(num)
	result := []string{}
	for _, resp := range splits[1:] {
		pr, err := ParseRESP(resp)
		if err != nil {
			return nil, err
		}
		for _, r := range pr {
			result = append(result, r)
		}
	}
	return result, nil
}

// $<length>\r\n<data>\r\n
func ParseBulkString(input string) ([]string, error) {
	splits := strings.Split(input, "\r\n")
	return []string{splits[1]}, nil
}