package main

import (
	"reflect"
	"testing"
)

func TestTokenizers(t *testing.T) {
	var examples = []struct {
		testCase       string
		input          string
		expectedOutput []string
	}{
		{"snake", "test", []string{"test"}},
		{"snake", "snake_case", []string{"snake", "case"}},
		{"snake", "c_c_c_c", []string{"c", "c", "c", "c"}},
		{"camel", "test", []string{"test"}},
		{"camel", "camelCase", []string{"camel", "Case"}},
		{"camel", "cCCC", []string{"c", "CCC"}},
		{"camel", "pleaseSplitHTMLCorrectly", []string{"please", "Split", "HTML", "Correctly"}},
		{"pascal", "Test", []string{"Test"}},
		{"pascal", "PascalCase", []string{"Pascal", "Case"}},
		{"pascal", "CCCC", []string{"CCCC"}},
		{"pascal", "PleaseSplitHTMLCorrectly", []string{"Please", "Split", "HTML", "Correctly"}},
		{"lower", "test", []string{"test"}},
		{"lower", "lower case", []string{"lower", "case"}},
		{"lower", "c c c c", []string{"c", "c", "c", "c"}},
		{"upper", "TEST", []string{"TEST"}},
		{"upper", "UPPER CASE", []string{"UPPER", "CASE"}},
		{"upper", "C C C C", []string{"C", "C", "C", "C"}},
		{"title", "Test", []string{"Test"}},
		{"title", "Title Case", []string{"Title", "Case"}},
		{"title", "C C C C", []string{"C", "C", "C", "C"}},
		{"sentence", "Test", []string{"Test"}},
		{"sentence", "Sentence case", []string{"Sentence", "case"}},
		{"sentence", "C c c c", []string{"C", "c", "c", "c"}},
		{"kebab", "test", []string{"test"}},
		{"kebab", "kebab-case", []string{"kebab", "case"}},
		{"kebab", "c-c-c-c", []string{"c", "c", "c", "c"}},
		{"dot", "test", []string{"test"}},
		{"dot", "dot.case", []string{"dot", "case"}},
		{"dot", "c.c.c.c", []string{"c", "c", "c", "c"}},
	}

	for _, example := range examples {
		actualOutput := inputCaseConfigurations[example.testCase].tokenize(example.input)
		if !reflect.DeepEqual(actualOutput, example.expectedOutput) {
			t.Errorf("Failed testing %s case tokenizer: got %s for given input %s, expected %s", example.testCase, actualOutput, example.input, example.expectedOutput)
		}
	}
}

func TestGuardInputCaseIsValid(t *testing.T) {
	_, error := ConvertCase("randomString", "camel", "text")
	if error == nil {
		t.Error("Expected failure for input case: 'randomString', didn't get a failure.")
	}
}

func TestGuardOutputCaseIsValid(t *testing.T) {
	_, error := ConvertCase("camel", "randomString", "text")
	if error == nil {
		t.Error("Expected failure for output case: 'randomString', didn't get a failure.")
	}
}

func TestConvertAllCasesOnHappyPath(t *testing.T) {
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

	for _, from := range inputCases {
		for _, to := range outputCases {
			inputText := caseToFormattedText[from]
			expectedOutput := caseToFormattedText[to]
			actualOutput, _ := ConvertCase(from, to, inputText)
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
	for _, from := range inputCases {
		for _, to := range outputCases {
			inputText := ""
			expectedOutput := ""
			actualOutput, _ := ConvertCase(from, to, inputText)
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
