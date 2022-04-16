package main

import (
	"bufio"
	"datafile/datafile"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("please enter the file path: ")

	reader := bufio.NewReader(os.Stdin)

	bytes, err := reader.ReadBytes('\n')

	if err != nil {
		log.Fatal(err)
	}

	fileName := string(bytes)

	fileName = strings.Trim(fileName, "\n")

	lines, err := datafile.GetString(fileName)

	if err == nil {
		fmt.Println(lines)
	} else {
		log.Fatal(err)
	}
}
