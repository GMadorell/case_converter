
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
$ case_converter -f camel -t snake someText
-> some_text
```
```
$ case_converter -f pascal -t kebab SomeText
-> some-text
```

Accepted cases:

* snake    (ej: some_example_sentence)
* camel    (ej: someExampleSentence)
* pascal   (ej: SomeExampleSentence)
* lower    (ej: some example sentence)
* upper    (ej: SOME EXAMPLE SENTENCE)
* title    (ej: Some Example Sentence)
* sentence (ej: Some example sentence)
* flat     (ej: someexamplesentence)  (not supported for -f flag)
* kebab    (ej: some-example-sentence)
* dot      (ej: some.example.sentence)
