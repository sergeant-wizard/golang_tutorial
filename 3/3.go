package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	ret := make(map[string]int)
	for _, field := range fields {
		_, ok := ret[field]
		if ok {
			ret[field] += 1
		} else {
			ret[field] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
