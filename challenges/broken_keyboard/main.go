package main

import "fmt"

func main() {

	cases := []struct {
		In string
		Ex string
	}{
		{"abcd", "cd"},
		{"ABCD", "CD"},
		{"abba", "a"},
		{"ABBA", "A"},
		{"aAbB", ""},
		{"xXyYzZbB", "xXyY"},
		{"xXyYBbzZ", "xXzZ"},
		{"abcdefgijklmnopqrstuvwxyz", "cdefgijklmnopqrstuvwxyz"},
		{"ABCDEFGIJKLMNOPQRSTUVWXYZ", "CDEFGIJKLMNOPQRSTUVWXYZ"},
	}

	for _, case_ := range cases {
		result := solution(case_.In)
		fmt.Println("Passed:", case_.Ex == result, "Expexted:", case_.Ex, "Actual:", result)
	}
}

/*
	  У Пети сломалась клавиатура.
	   Когда он вводит b, то вместо этого стирается последняя введенная строчная буква.
		 Когда вводит B (большая английская бэ), то стирается последняя введенная
		 заглавная буква. Остальные работают нормально.
		 Дана строка из больших и маленьких английских букв,
		 показывающая последовательность нажатия клавиш.
		 Нужно вывести, что будет введено по факту.
*/
func solution(in string) string {
	// Here is chars[i][0] is the original char
	//  chars[i][1] is the deletion flag
	chars := make([][2]byte, len(in))

	for i, v := range []byte(in) {
		chars[i] = [2]byte{v, 0}
	}
	var result []byte
	// Here is last[0] is index of the last capital char
	//  last[1] is index of the last non capital char
	last := [2]int{-1, -1}

	for i, char := range chars {
		if char[0] == byte('b') {
			if last[1] > -1 {
				chars[last[1]][1] = 1
				chars[i][1] = 1
			}
		} else if char[0] == byte('B') {
			if last[0] > -1 {
				chars[last[0]][1] = 1
				chars[i][1] = 1
			}
		} else {
			if isCapital(char[0]) {
				last[0] = i
			} else {
				last[1] = i
			}
		}
	}

	for _, v := range chars {
		if v[1] == 0 {
			result = append(result, v[0])
		}
	}

	return string(result)
}

func isCapital(ch byte) bool {
	return ch < byte('a')
}
