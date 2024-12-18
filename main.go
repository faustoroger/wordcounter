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

	if *countLines && *countRunes {
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		if *countLines {
			fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), "number of lines.")
		} else if *countRunes {
			fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), "number of runes.")
		} else {
			fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), "number of words.")
		}
	}

	var (
		f   *os.File
		err error
	)

	if flag.NArg() == 1 {
		f, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file %s: %s\n", flag.Arg(0), err)
			os.Exit(1)
		}
	}

	defer f.Close()

	if f != nil {
		if *countLines && !*countRunes {
			fmt.Println("The content has", count(f, *countLines, *countRunes), "number of lines.")
		} else if *countRunes && !*countLines {
			fmt.Println("The content has", count(f, *countLines, *countRunes), "number of runes.")
		} else {
			fmt.Println("The content has", count(f, *countLines, *countRunes), "number of words.")
		}
	}
}

func count(r io.Reader, countLines bool, countRunes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines && !countRunes {
		scanner.Split(bufio.ScanWords)
	} else if !countLines && countRunes {
		scanner.Split(bufio.ScanRunes)
	} else if countLines && !countRunes {
		scanner.Split(bufio.ScanLines)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
