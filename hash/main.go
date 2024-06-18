package main

import "fmt"

func hashint64(n int64) uint64 {
	hash := uint64(n % 1000)

	return hash
}

func hashstr(v string) uint64 {
	var hash uint64

	for _, char := range v {
		hash = (hash*31 + uint64(char)) % 1000
	}

	return hash
}

func main() {
	number := int64(123456789)
	hash := hashint64(number)
	fmt.Println(hash)

	str := "Hello world!"
	hash = hashstr(str)
	fmt.Println(hash)
}
