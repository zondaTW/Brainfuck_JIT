package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func great(buf []byte, ptr *int) {
	*ptr++
}

func less(buf []byte, ptr *int) {
	*ptr--
}

func plus(buf []byte, ptr *int) {
	buf[*ptr]++
}

func hyphen(buf []byte, ptr *int) {
	buf[*ptr]--
}

func dot(buf []byte, ptr *int) {
	fmt.Printf("%c", buf[*ptr])
}

func comma(buf []byte, ptr *int) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	buf[*ptr] = []byte(input)[0]
}

var filename string

func init() {
	flag.StringVar(&filename, "i", "", "input brainfuck file")
}

func main() {
	flag.Parse()
	if filename == "" {
		fmt.Println("Filename empty")
		return
	}

	Brainfuck := map[string]func(buf []byte, ptr *int){
		">": great,
		"<": less,
		"+": plus,
		"-": hyphen,
		".": dot,
		",": comma,
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 256)
	ptr := 0
	whileEipList := list.New()

	var value string = ""
	for idx, tempIdx := 0, 0; ; {
		if idx >= len(content) {
			break
		}

		value = string(content[idx])
		if brainfuck, ok := Brainfuck[value]; ok {
			brainfuck(buf, &ptr)
		}

		switch value {
		case "[":
			if buf[ptr] != 0 {
				whileEipList.PushBack(idx)
			} else {
				idx = whileEipList.Back().Value.(int)
				whileEipList.Remove(whileEipList.Back())
			}
		case "]":
			tempIdx = idx
			idx = whileEipList.Back().Value.(int)
			whileEipList.Remove(whileEipList.Back())
			whileEipList.PushBack(tempIdx)
			continue
		}

		idx++
	}
}
