package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":  {"discrete math"},
	"databases":        {"data structures"},
	"discrete math":    {"intro to programming"},
	"formal languages": {"disctrete math"},
	"networks":         {"operatiing systems"},
	"operatiing systems": {
		"data structures",
		"computer organization"},
	"programming languages": {
		"data structures",
		"computer organization"},
}

func main() {
	for _, key := range prereqs {
		fmt.Println(key)
	}
}
