package main

import "fmt"

const (
	ENGLISH = "english"
	SPANISH = "spanish"
)

const ENGLISH_HELLO_FMT = "Hello, %s!"
const SPANISH_HELLO_FMT = "Hola, %s!"

func Hello(name, language string) string {
	target := "World"
	if name != "" {
		target = name
	}
	pattern := ENGLISH_HELLO_FMT
	switch language {
	case SPANISH:
		pattern = SPANISH_HELLO_FMT
	}
	return fmt.Sprintf(pattern, target)
}
