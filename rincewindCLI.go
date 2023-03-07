package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheScottBot/rincewind"
)

var tempStdout *os.File

func main() {

	translation := flag.String("t", "flag -t unset", "Value to be translated")
	targetLang := flag.String("ta", "flag -ta unset", "Target language")
	sourceLang := flag.String("s", "flag -s unset", "Source")

	flag.Parse()

	easterEgg(*translation)
	unassignStdout()

	changeDefaults(targetLang, sourceLang)

	translationRequest := rincewind.TranslationRequest{
		TranslateText:  *translation,
		SourceLanguage: *sourceLang,
		TargetLanguage: *targetLang,
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

func changeDefaults(target *string, source *string) {
	if *target == "flag -ta unset" {
		fmt.Println("default target detected, setting to DE\n")
		*target = "DE"
	}

	if *source == "flag -s unset" {
		fmt.Println("default source detected, setting to EN\n")
		*source = "EN"
	}
}

func easterEgg(text string) {
	if text == "wizard" {
		fmt.Println("I think you mean 'Wizzard'")
	}
}
