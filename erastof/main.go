package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	type data_t int64
	var crez []bool
	debug := false
	eratos := func() {
		count := func() data_t {
			var j data_t
			for i := range crez {
				if crez[i] {
					j++
				}
			}
			return j
		}
		for i := range crez {
			if crez[i] {
				for j := i + i; j < len(crez); j += i {
					// вычисление всех чисел, кратных i
					crez[j] = false
				}
			}
		}
		fmt.Printf("count: %d\n", count())
	}
	// диагностика среза
	show := func(s []bool) {
		const inlin = 10
		var j data_t
		for i := range s {
			if s[i] {
				j++
				fmt.Printf("%4d", i)
				if 0 == (j % inlin) {
					fmt.Printf("\n")
				}
				if j%inlin != 0 {
					fmt.Printf("\n")
				}
			}
		}
	}
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s [-]<number>\n", os.Args[0])
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	var long data_t = data_t(n)
	if long < 0 {
		long, debug = -long, true
	}
	crez = make([]bool, long)
	for i := range crez {
		if i > 1 {
			crez[i] = true
		}
	}
	eratos()
	if debug {
		show(crez)
	}
}
