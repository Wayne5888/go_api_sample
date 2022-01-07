package test

func Add(x int, y int) int {
	return x + y
}

func IsFindName(findName string, names []string) bool {
	for _, n := range names {
		if n == findName {
			return true
		}
	}
	return false
}
