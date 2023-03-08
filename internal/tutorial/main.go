package main

import "fmt"

func Hello(name string) string {
	target := "World"
	if name != "" {
		target = name
	}
	return fmt.Sprintf("Hello, %s!", target)
}
