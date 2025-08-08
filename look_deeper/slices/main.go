package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	initSlice := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cutSlice := initSlice[2:4]

	fmt.Printf("initSlice %v cap: %d, len: %d\n", initSlice, cap(initSlice), len(initSlice))
	fmt.Printf("cutSlice %v cap: %d, len: %d\n", cutSlice, cap(cutSlice), len(cutSlice))
	initSlice[5] = 55
	cutSlice = cutSlice[0:8]
	fmt.Printf("cutSlice %v cap: %d, len: %d\n", cutSlice, cap(cutSlice), len(cutSlice))

	fmt.Printf("|init: %d, cut: %d|\n"+
		"|init: %v, cut: %v|\n",
		reflect.ValueOf(initSlice).Pointer(),
		reflect.ValueOf(cutSlice).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&initSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cutSlice)),
	)

	cutSlice = append(cutSlice, 255)
	initSlice[5] = 155
	fmt.Printf("initSlice %v cap: %d, len: %d\n", initSlice, cap(initSlice), len(initSlice))
	fmt.Printf("cutSlice %v cap: %d, len: %d\n", cutSlice, cap(cutSlice), len(cutSlice))
	fmt.Printf("|init: %d, cut: %d|\n"+
		"|init: %v, cut: %v|\n",
		reflect.ValueOf(initSlice).Pointer(),
		reflect.ValueOf(cutSlice).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&initSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cutSlice)),
	)

	fmt.Printf("init size: %d\n", unsafe.Sizeof(initSlice))
	fmt.Printf("str size: %d, len: %d, bytes: %v\n", unsafe.Sizeof(""), len("Ё"), []byte("Ё"))
}
