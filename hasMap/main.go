package main

import "fmt"

type hashmap struct {
	data map[string]string
}

func (h *hashmap) Set(key, val string) {
	h.data[key] = val
}
func (h *hashmap) Get(key string) (string, bool) {
	val, ok := h.data[key]
	return val, ok
}
func (h *hashmap) Delete(key string) {
	delete(h.data, key)
}

func main() {
	h := &hashmap{
		data: make(map[string]string),
	}

	h.Set("key1", "value1")
	h.Set("key2", "value2")

	value, ok := h.Get("key1")
	if ok {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Value not found")
	}

	h.Delete("key2")
	value, ok = h.Get("key2")
	if ok {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Value not found")
	}

	fmt.Println(h.data)
}
