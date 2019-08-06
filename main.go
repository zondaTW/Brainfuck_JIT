package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func great() {
	fmt.Printf("++ptr;")
}

func less() {
	fmt.Printf("--ptr;")
}

func plus() {
	fmt.Printf("++*ptr;")
}

func hyphen() {
	fmt.Printf("--*ptr;")
}

func dot() {
	fmt.Printf("putchar(*ptr);")
}

func comma() {
	fmt.Printf("*ptr = getchar();")
}

func open_bracket() {
	fmt.Printf("while (*ptr) {")
}

func close_bracket() {
	fmt.Printf("}")
}

func main() {
	Brainfuck := map[string]func(){
		">": great,
		"<": less,
		"+": plus,
		"-": hyphen,
		".": dot,
		",": comma,
		"[": open_bracket,
		"]": close_bracket,
	}

	content, err := ioutil.ReadFile("hello_world.bf")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range content {
		if brainfuck, ok := Brainfuck[string(v)]; ok {
			brainfuck()
		}
	}
}
