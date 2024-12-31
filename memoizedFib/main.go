package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Memoized func(int) int

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

var fibMem Memoized

func FibMemoized(n int) int {
	n += 1
	fibMem = Memoize(func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fibMem(n-2) + fibMem(n-1)
	})
	return fibMem(n)
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	t := time.Now()
	fmt.Printf("%d [%v]\n", FibMemoized(n), time.Since(t))
}
