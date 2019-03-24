package main

import (
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
)

var inputCases = []string{
	"snake",
	"camel",
	"pascal",
	"lower",
	"upper",
	"title",
	"sentence",
	"kebab",
	"dot",
}

var outputCases = []string{
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

func ConvertCase(from string, to string, text string) (string, error) {
	if !contains(inputCases, from) {
		return "", fmt.Errorf("Unknown 'from' case: %s", from)
	}
	if !contains(outputCases, to) {
		return "", fmt.Errorf("Unknown 'to' case: %s", to)
	}

	tokens := inputCaseConfigurations[from].tokenize(text)
	var convertedTokens []string
	for i, token := range tokens {
		convertedTokens = append(convertedTokens, outputCaseConfigurations[to].convertToken(i, token))
	}

	return strings.Join(convertedTokens, outputCaseConfigurations[to].joinWith), nil
}

type inputCaseConfiguration struct {
	tokenize func(text string) []string
}

type outputCaseConfiguration struct {
	convertToken func(index int, token string) string
	joinWith     string
}

var inputCaseConfigurations = map[string]inputCaseConfiguration{
	"snake": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, "_")
		},
	},
	"camel": inputCaseConfiguration{
		tokenize: camelcase.Split,
	},
	"pascal": inputCaseConfiguration{
		tokenize: camelcase.Split,
	},
	"lower": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, " ")
		},
	},
	"upper": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, " ")
		},
	},
	"title": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, " ")
		},
	},
	"sentence": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, " ")
		},
	},
	"kebab": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, "-")
		},
	},
	"dot": inputCaseConfiguration{
		tokenize: func(text string) []string {
			return strings.Split(text, ".")
		},
	},
}

var outputCaseConfigurations = map[string]outputCaseConfiguration{
	"snake": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return strings.ToLower(token)
		},
		joinWith: "_",
	},
	"camel": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			if index == 0 {
				return strings.ToLower(token)
			} else {
				return capitalize(token)
			}
		},
		joinWith: "",
	},
	"pascal": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return capitalize(token)
		},
		joinWith: "",
	},
	"lower": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return strings.ToLower(token)
		},
		joinWith: " ",
	},
	"upper": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return strings.ToUpper(token)
		},
		joinWith: " ",
	},
	"title": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return capitalize(token)
		},
		joinWith: " ",
	},
	"sentence": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			if index == 0 {
				return capitalize(token)
			} else {
				return strings.ToLower(token)
			}
		},
		joinWith: " ",
	},
	"flat": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return strings.ToLower(token)
		},
		joinWith: "",
	},
	"kebab": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return strings.ToLower(token)
		},
		joinWith: "-",
	},
	"dot": outputCaseConfiguration{
		convertToken: func(index int, token string) string {
			return strings.ToLower(token)
		},
		joinWith: ".",
	},
}

func capitalize(input string) string {
	return strings.Title(strings.ToLower(input))
}

func contains(array []string, element string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}
	return false
}
