package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	if *lines {
		fmt.Println("The statement has", count(os.Stdin, *lines), "number of lines.")
	} else {
		fmt.Println("The statement has", count(os.Stdin, *lines), "number of words.")
	}
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
