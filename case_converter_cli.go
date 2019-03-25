package main

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	From string `short:"f" long:"from" required:"true" name:"Format to convert from"`
	To   string `short:"t" long:"to"   required:"true" name:"Format to convert to"`
}

const help string = `
Usage:
$ case_converter -f format_from -t format_to some_text

Examples:
$ case_converter --from camel --to snake someText
-> some_text
$ case_converter -f pascal -t kebab SomeText
-> some-text

Options:
     -f, --from format of the input  
     -t, --to   format of the output

Format options (accepted by the -f and -t flags):
* snake    (ej: some_example_sentence)
* camel    (ej: someExampleSentence)
* pascal   (ej: SomeExampleSentence)
* lower    (ej: some example sentence)
* upper    (ej: SOME EXAMPLE SENTENCE)
* title    (ej: Some Example Sentence)
* sentence (ej: Some example sentence)
* flat     (ej: someexamplesentence)  (not supported for -t flag)
* kebab    (ej: some-example-sentence)
* dot      (ej: some.example.sentence)
`

func main() {
	args, err := flags.Parse(&opts)

	if err != nil {
		fmt.Println(help)
		log.Fatal(err)
	}

	if len(args) != 1 {
		fmt.Println(help)
		log.Fatal("case_converter accepts a single argument")
	}

	converted, err := ConvertCase(opts.From, opts.To, args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(converted)
}
