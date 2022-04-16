package datafile

import (
	"bufio"
	"os"
)

/*
read strings from file, return string slice or error.
*/

func GetString(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err1 := file.Close()

	if err1 != nil {
		return nil, err1
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
