package main

import "fmt"

const (
	ENGLISH = "english"
	SPANISH = "spanish"
	FRENCH  = "french"
)

const ENGLISH_HELLO_FMT = "Hello, %s!"
const SPANISH_HELLO_FMT = "Hola, %s!"
const FRENCH_HELLO_FMT = "Bonjour, %s!"

func Hello(name, language string) string {
	return fmt.Sprintf(
		getPatternByLanguage(language),
		emptyOr(name, "World"),
	)
}

func getPatternByLanguage(language string) string {
	switch language {
	case SPANISH:
		return SPANISH_HELLO_FMT
	case FRENCH:
		return FRENCH_HELLO_FMT

	default:
		return ENGLISH_HELLO_FMT
	}
}

func emptyOr(val, df string) string {
	if val == "" {
		return df
	}
	return val
}
