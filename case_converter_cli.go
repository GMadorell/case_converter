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
* snake    (e.g.: some_example_sentence)
* camel    (e.g.: someExampleSentence)
* pascal   (e.g.: SomeExampleSentence)
* lower    (e.g.: some example sentence)
* upper    (e.g.: SOME EXAMPLE SENTENCE)
* title    (e.g.: Some Example Sentence)
* sentence (e.g.: Some example sentence)
* flat     (e.g.: someexamplesentence)  (not supported for -t flag)
* kebab    (e.g.: some-example-sentence)
* dot      (e.g.: some.example.sentence)
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
