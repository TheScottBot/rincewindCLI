package main

import (
	"fmt"
	"os"

	"github.com/TheScottBot/rincewind"
	"github.com/jessevdk/go-flags"
)

var tempStdout *os.File

type Options struct {
	TranslationText string `short:"t" long:"text" description:"Text to be translated" required:"true"`
	SourceLang      string `short:"s" long:"source" description:"Source text language code" optional:"yes" optional-value:"DE"`
	TargetLang      string `short:"i" long:"intended" description:"Target translation language code" optional:"yes" optional-value:"EN"`
}

var options Options

func main() {
	_, err := flags.Parse(&options)

	if err != nil {
		fmt.Println(err)
	}

	easterEgg(options.TranslationText)
	unassignStdout()

	translationRequest := rincewind.TranslationRequest{
		TranslateText:  options.TranslationText,
		SourceLanguage: options.SourceLang,
		TargetLanguage: options.TargetLang,
	}

	value, _ := rincewind.Translate(translationRequest)

	reassignStdout()

	if len(value.Translations) < 1 {
		fmt.Println("Something went wrong and there was no translation response")
		return
	}
	fmt.Println(value.Translations[0].Text)
}

func unassignStdout() {
	tempStdout = os.Stdout
	os.Stdout = nil
}

func reassignStdout() {
	os.Stdout = tempStdout
}

func easterEgg(text string) {
	if text == "wizard" {
		fmt.Println("I think you mean 'Wizzard'")
	}
}
