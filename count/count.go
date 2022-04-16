package count

func Count(names []string) (map[string]int, error) {
	voteCount := make(map[string]int, 0)

	for _, name := range names {
		voteCount[name] = voteCount[name] + 1
	}

	return voteCount, nil
}
