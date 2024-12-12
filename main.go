package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countRunes := flag.Bool("r", false, "Count runes")
	flag.Parse()

	if !*countLines && !*countRunes {
		fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), "number of words.")
	} else if *countRunes {
		fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), "number of runes.")
	} else {
		fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), "number of lines.")
	}
}

func count(r io.Reader, countLines bool, countRunes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines && !countRunes {
		scanner.Split(bufio.ScanWords)
	} else if countRunes {
		scanner.Split(bufio.ScanRunes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
