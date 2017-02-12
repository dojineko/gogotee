package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if (len(os.Args) <= 1) || os.Args[1] == "" {
		log.Fatalln("filename is empty")
	}
	filename := os.Args[1]

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("file open failed")
	}
	defer file.Close()

	tee := io.TeeReader(os.Stdin, file)

	s := bufio.NewScanner(tee)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
