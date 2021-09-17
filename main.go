package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open file
	f, err := os.Open("test.txt")
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

	for fr.Scan() {
		fmt.Println(fr.Text())
	}

}
