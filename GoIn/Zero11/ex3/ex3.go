package ex3

import "fmt"

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

type Logger interface {
	Log(message string)
}

type Log struct {
	Level LogLevel
}

func (l *Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Println("ERROR:", message)
	case Info:
		fmt.Println("INFO:", message)
	default:
		fmt.Println("UNKNOWN LEVEL:", message)
	}
}