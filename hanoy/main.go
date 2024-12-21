package main

import (
	"os"
	"strconv"
)

var debug bool = false
var nopr uint64 = 0

func move(from, to, count int) {
	put := func(from, to int) {
		nopr++
		if !debug {
			return
		}
		print(from, " => ", to, ", ")
		if 0 == (nopr % 5) {
			print("\n")
		}
	}
	temp := func(from, to int) int {
		for i := 1; i <= 3; i++ {
			if i != from && i != to {
				return i
			}
		}
		panic(0)
	}

	if count > 1 {
		move(from, temp(from, to), count-1)
	}
	put(from, to)
	if count > 1 {
		move(temp(from, to), to, count-1)
	}
}

func main() {
	if len(os.Args) != 2 {
		print("usage: ", os.Args[0], " [-]<number>\n")
		return
	}

	debug = false
	size, _ := strconv.Atoi(os.Args[1])
	if size < 0 {
		size, debug = -size, true
	}
	move(1, 3, size)
	if debug && (nopr%5) != 0 {
		print("\n")
	}
	print("число перемещений ", nopr, "\n")
}
