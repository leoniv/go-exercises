package main

func main() {
	solution([]int{2, 5, 6, 7}, []int{1, 3, 4, 6})
}

/*
Два массива интов, могут быть разной длины.

	  Отсортированы по возрастанию.
		Вывести объединение без дублей,
		сохранив сортировку и использовав
		константу дополнительной памяти.
*/
func solution(in1 []int, in2 []int) {
	if len(in1) == 0 || len(in2) == 0 {
		for _, v := range append(in1, in2...) {
			print(v, " ")
		}
		return
	}
	for len(in1) > 0 || len(in2) > 0 {
		if len(in1) == 0 {
			print(in2[0], " ")
			in2 = in2[1:]
		} else if len(in2) == 0 {
			print(in1[0], " ")
			in1 = in1[1:]
		} else {
			if in1[0] == in2[0] {
				print(in1[0], " ")
				in1 = in1[1:]
				in2 = in2[1:]
			} else if in1[0] < in2[0] {
				print(in1[0], " ")
				in1 = in1[1:]
			} else { // in1[0] > in2[0]
				print(in2[0], " ")
				in2 = in2[1:]
			}
		}
	}
}
