package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("$ go-doublet <dicionary_file.txt>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	dictFname := os.Args[1]
	dict := initDictionary(dictFname)
	graph := initWordgraph(dict)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nEnter two equal-length words separated by a space: ")
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		if len(words) < 2 {
			fmt.Printf("Exiting.\n")
			os.Exit(0)
		}

		result, dur := Doublet(words[0], words[1], dict, graph)
		if result != nil {
			fmt.Printf("Found a path from %s to %s in %d steps:\n", words[0], words[1], (len(*result) - 1))
			for i, step := range *result {
				fmt.Printf("%s ", step)
				if i < len(*result)-1 {
					fmt.Printf("--> ")
					if i > 0 && i%10 == 0 {
						fmt.Printf("\n\t")
					}
				}
			}
			fmt.Printf("\n")
		} else {
			fmt.Printf("No path could be found between %s and %s\n", words[0], words[1])
		}
		fmt.Printf("Algorithm ran in %d us\n", dur/1000)
		fmt.Printf("\nEnter two equal-length words separated by a space: ")
	}

	os.Exit(0)
}
