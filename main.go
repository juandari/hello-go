package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language , e.g. en, el, fr, he, ur, vi")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello world",
	"el": "Γειά σου κόσμε",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào thế giới",
	"id": "Halo dunia",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
