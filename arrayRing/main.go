package main

import "fmt"

// IntRing - кольцевой массив
type IntRing struct {
	size  int
	data  []*int
	Start int
	End   int
}

// NewIntRing - конструктор кольца
func NewIntRing(size, start int) (*IntRing, error) {
	if start >= size {
		return nil, fmt.Errorf("start %d is greater than size %d", start, size)
	}
	return &IntRing{
		size,
		make([]*int, size, size),
		start,
		start,
	}, nil
}

// Size - получение размера кольца
func (r IntRing) Size() int {
	return r.size
}

// IsEmpty - пустое ли кольцо
func (r *IntRing) IsEmpty() bool {
	return r.Start == r.End
}

// IsFull - достигнут ли конец
func (r *IntRing) IsFull() bool {
	return (r.End < r.Start && r.Start-r.End == 1) || (r.Start == 0 && r.End == r.size-1)
}

// Read - чтение элемента
func (r *IntRing) Read() (int, error) {
	if !r.IsEmpty() {
		el := r.data[r.Start]
		r.Start++
		for el == nil && r.Start < r.End {
			el = r.data[r.Start]
			r.Start++
		}
		if el == nil {
			return 0, fmt.Errorf("cannot read value from %d at %d", r.Start, r.End)
		}

		return *el, nil
	}
	return 0, fmt.Errorf("cannot read value from %d at %d", r.Start, r.End)
}

// Writ - добавление одного элемента
func (r *IntRing) Write(v int) error {
	if r.IsEmpty() || !r.IsFull() {
		r.data[r.End] = &v
		r.End++

		return nil
	}
	return fmt.Errorf("buffer full")
}

// Remove - удаление элемента
func (r *IntRing) Remove(index int) error {
	if index < 0 || index > r.size {
		return fmt.Errorf("invalid list index")
	}
	r.data[index] = nil

	return nil
}

// Print - печать содержимого кольца
func (r *IntRing) Pritn() {
	if r.Start == r.End {
		for _, el := range r.data[r.Start:r.End] {
			fmt.Printf("%d\t", *el)
		}
		fmt.Printf("\n")
	} else if r.Start > r.End {
		tempData := append(r.data[r.Start:], r.data[:r.End]...)
		for _, el := range tempData {
			fmt.Printf("%d\t", *el)
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("list empty")
	}
}

func main() {
	fmt.Println("Создания кольцевого массива")
	ring, err := NewIntRing(5, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Pritn()

	fmt.Println("Записываем один элемент '1'")
	err = ring.Write(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Записываем еще один элемент '2'")
	err = ring.Write(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Удалим 0 елемент")
	err = ring.Remove(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Прочитаем что внутри")
	el, err := ring.Read()
	fmt.Println(el)
	fmt.Println("Еще разок прочитаем что внутри")
	el, err = ring.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("И еще разок прочитаем что внутри")
	fmt.Println(el)
	el, err = ring.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
}
