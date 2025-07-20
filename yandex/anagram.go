package yandex

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	mapS1 := CalcStat(s1)
	mapS2 := CalcStat(s2)
	if IsItAnagram(mapS1, mapS2) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func CalcStat(s string) map[rune]int {
	mapS := make(map[rune]int)
	for _, char := range s {
		mapS[char] += 1
	}
	return mapS
}

func IsItAnagram(mapS1 map[rune]int, mapS2 map[rune]int) bool {
	chars := make(map[rune]bool)
	for char, _ := range mapS1 {
		chars[char] = true
	}
	for char, _ := range mapS2 {
		chars[char] = true
	}

	for char, _ := range chars {
		count1, ok1 := mapS1[char]
		count2, ok2 := mapS2[char]
		if !(ok1 && ok2 && count1 == count2) {
			return false
		}
	}

	return true
}
