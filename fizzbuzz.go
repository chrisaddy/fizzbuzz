// Package fizzbuzz buzzes fizzes.
package main

import (
	"fmt"
	"go/doc"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type Maybe interface {
	Return(value interface{}) Maybe
	Bind(func(interface{}) Maybe) Maybe
}

// Given an array of strings and a search string,
// contains checks for the inclusion of the search string in the array.
func contains(s []string, term string) bool {
	i := sort.SearchStrings(s, term)
	return i < len(s) && s[i] == term
}

func fizzbuzz(count int) string {
	if count%15 == 0 {
		return "FIZZBUZZ"
	}
	if count%3 == 0 {
		return "FIZZ"
	}

	if count%5 == 0 {
		return "BUZZ"
	}

	return strconv.Itoa(count)
}

func runner(mood []string) {
	sort.Strings(mood)
	i := 1
	for {
		fmt.Println(fizzbuzz(i))
		if contains(mood, "gentle") {
			time.Sleep(time.Second)

		}
		if contains(mood, "scatter-brained") {
			i = rand.Intn(10000)
		}
		if contains(mood, "impatient") {
			if i == 100 {
				return
			}
		}
		i += 1
	}

	return
}

func main() {
	mood := os.Args[1:]
	runner(mood)
}
