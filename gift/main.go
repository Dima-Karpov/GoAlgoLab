package main

import (
	"fmt"
	"sort"
)

type Gift struct {
	Name  string
	Price int
}

func main() {
	gifts := []Gift{
		{Name: "Кукла", Price: 100},
		{Name: "Машинка", Price: 200},
		{Name: "Коструктор", Price: 150},
		{Name: "Книга", Price: 350},
	}

	sort.Slice(gifts, func(i, j int) bool {
		return gifts[i].Price < gifts[j].Price
	})

	fmt.Println("Подарки отстортированы по убыванию стоимости: ")
	for _, gift := range gifts {
		fmt.Printf("%s - %d\n", gift.Name, gift.Price)
	}
}
