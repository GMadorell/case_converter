package case_converter

import (
	"testing"
)

func TestConvertCaseOnHappyPath(t *testing.T) {
	caseToFormattedText := map[string]string{
		"snake":    "some_example_sentence",
		"camel":    "someExampleSentence",
		"pascal":   "SomeExampleSentence",
		"lower":    "some example sentence",
		"upper":    "SOME EXAMPLE SENTENCE",
		"title":    "Some Example Sentence",
		"sentence": "Some example sentence",
		"flat":     "someexamplesentence",
		"kebab":    "some-example-sentence",
		"dot":      "some.example.sentence",
	}

	for _, from := range cases {
		for _, to := range cases {
			inputText := caseToFormattedText[from]
			expectedOutput := caseToFormattedText[to]
			actualOutput := ConvertCase(from, to, inputText)
			if actualOutput != expectedOutput {
				t.Errorf(
					"Conversion from '%s' to '%s' with input text '%s' failed, got: '%s', want: '%s'.",
					from,
					to,
					inputText,
					actualOutput,
					expectedOutput,
				)
			}
		}
	}
}

func TestEmptyInput(t *testing.T) {
	for _, from := range cases {
		for _, to := range cases {
			inputText := ""
			expectedOutput := ""
			actualOutput := ConvertCase(from, to, inputText)
			if actualOutput != expectedOutput {
				t.Errorf(
					"Conversion from '%s' to '%s' with input text '%s' failed, got: '%s', want: '%s'.",
					from,
					to,
					inputText,
					actualOutput,
					expectedOutput,
				)
			}
		}
	}
}
