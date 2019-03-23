package case_converter

import (
	"fmt"
)

var cases = [...]string{
	"snake",
	"camel",
	"pascal",
	"lower",
	"upper",
	"title",
	"sentence",
	"flat",
	"kebab",
	"dot",
}

func ConvertCase(from string, to string, text string) string {
	fmt.Printf("From: %s, to: %s, text: %s\n", from, to, text)
	return "Test"
}
