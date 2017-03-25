package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: wc FILE [FILE...]")
		os.Exit(-1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input file:", err)
		os.Exit(-1)
	}

	lineCounter, wordCounter, characterCounter := 0, 0, 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		characterCounter++
		if scanner.Text() == "\n" {
			lineCounter++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(-1)
	}

	report(lineCounter, wordCounter, characterCounter, os.Args[1])
}

func report(lcount, wcount, ccount int, filename string) {
	fmt.Printf("%6d %6d %6d %s\n", lcount, wcount, ccount, filename)
}
