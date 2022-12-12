package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me smething to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		buf = append(buf, text[i])

		if len(text[i:]) >= nlink && text[i:i+nlink] ==  link {
		
		}

	}
}
