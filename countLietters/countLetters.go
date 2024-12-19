package countLetters

// CountLetters подсчитывает количество символов в строке
func CountLetters(s string, l byte) (c int) {
	c = 0
	for _, b := range []byte(s) {
		if b == l {
			c++
		}
	}
	return
}
