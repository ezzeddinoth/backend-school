package utils

func UniqueIntSlice(input []int) []int {
	if len(input) < 2 {
		return input
	}
	allKeys := make(map[int]bool)
	s := []int{}
	for _, item := range input {
		if _, exist := allKeys[item]; !exist {
			allKeys[item] = true
			s = append(s, item)
		}
	}
	return s
}
