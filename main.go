package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// parse args
	if len(os.Args) <= 1 || len(os.Args) >= 4 {
		fmt.Println("num of arguments is invalid")
		fmt.Println("arg1 : -n [optional]")
		fmt.Println("arg2 : file name [positional]")
		return
	}
	args := os.Args[1:]
	// - set n option flag
	print_line_numbers := false
	if len(args) == 2 && args[0] == "-n" {
		print_line_numbers = true
	}
	fmt.Println(print_line_numbers)

	// - set input file name
	filename := ""
	if len(args) == 2 {
		filename = args[1]
	} else {
		filename = args[0]
	}

	// open file
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	// read lines
	fr := bufio.NewScanner(f)
	if err := fr.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var i int = 0
	for fr.Scan() {
		if print_line_numbers {
			fmt.Print(i, " : ")
		}
		fmt.Println(fr.Text())
		i++
	}

}
