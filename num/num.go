package num

// Num вовзращает число Фибоначи
// по  индексу
func Num(n int) int {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
