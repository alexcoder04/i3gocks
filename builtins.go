package main

import (
	"strings"
	"time"
)

var BuiltIns map[string]func([]string) []string = map[string]func([]string) []string{
	"echo": func(args []string) []string {
		return []string{strings.Join(args, " ")}
	},

	"time": func(args []string) []string {
		var format string
		if len(args) < 1 {
			format = "02.01.2006 (Mon) 15:04:05"
		} else {
			format = args[0]
		}
		return []string{time.Now().Format(format)}
	},
}
