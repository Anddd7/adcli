package main

import "fmt"

const ENGLISH_HELLO_FMT = "Hello, %s!"

func Hello(name string) string {
	target := "World"
	if name != "" {
		target = name
	}
	return fmt.Sprintf(ENGLISH_HELLO_FMT, target)
}
