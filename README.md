
# Case Converter
Haven't you always wanted to convert that pesky camelCase string to beautiful snake\_case? You just found the perfect CLI solution for it!

## Installation

```
go get github.com/gmadorell/case_converter
```

## Usage

```
$ case_converter -f format_from -t format_to some_text
```

Examples:
```
$ case_converter --from camel --to snake someText
-> some_text
```
```
$ case_converter -f pascal -t kebab SomeText
-> some-text
```

Accepted cases:

* snake    (e.g. some_example_sentence)
* camel    (e.g. someExampleSentence)
* pascal   (e.g. SomeExampleSentence)
* lower    (e.g. some example sentence)
* upper    (e.g. SOME EXAMPLE SENTENCE)
* title    (e.g. Some Example Sentence)
* sentence (e.g. Some example sentence)
* flat     (e.g. someexamplesentence)  (not supported for -f flag)
* kebab    (e.g. some-example-sentence)
* dot      (e.g. some.example.sentence)
