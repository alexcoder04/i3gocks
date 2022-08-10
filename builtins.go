package main

import (
	"strings"
	"time"
)

func ExecuteBuiltIn(name string, args []string) []string {
	switch name {
	case "echo", "print":
		return BuiltInEcho(args)
	case "date", "time", "datetime":
		return BuiltInDateTime(args)
	}
	return []string{" builtin not found"}
}

func BuiltInEcho(args []string) []string {
	return []string{strings.Join(args, " ")}
}

func BuiltInDateTime(args []string) []string {
	var format string
	if len(args) < 1 {
		format = "02.01.2006 (Mon) 15:04:05"
	} else {
		format = args[0]
	}
	return []string{time.Now().Format(format)}
}
