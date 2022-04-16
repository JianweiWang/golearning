package count

import (
	"bufio"
	"datafile/golearning/datafile"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/JianweiWang/golearning/count"
)

func Count(names []string) (map[string]int, error) {
	voteCount := make(map[string]int, 0)

	for _, name := range names {
		voteCount[name] = voteCount[name] + 1
	}

	return voteCount, nil
}

func GetVotes() {
	fmt.Println("please enter the file path: ")

	reader := bufio.NewReader(os.Stdin)

	bytes, err := reader.ReadBytes('\n')

	if err != nil {
		log.Fatal(err)
	}

	fileName := string(bytes)

	fileName = strings.Trim(fileName, "\n")

	lines, err := datafile.GetString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	votes, err := count.Count(lines)

	if err != nil {
		log.Fatal(err)
	}

	var keys []string

	for k := range votes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %d \n", k, votes[k])
	}
}
