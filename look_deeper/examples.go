package lookdeeper

import (
	"fmt"
	"unsafe"
)

func Run() {
	fmt.Println("Looking deeper to:")
	lookToSlice()
	lookToMap()
	lookIntoPointers()
}

func lookIntoPointers() {
	return
	var ip *int
	var ipp **int = &ip
	fmt.Println("Look into pointers (var ip *int)", "pointer var address:", ipp, "pointer:", ip, "pointed value:", "")
	ip = new(int)
	fmt.Println("Look into pointers (var ip *int)", "pointer var address:", ipp, "pointer:", ip, "pointed value:", **ipp)
	*ip = 10
	fmt.Println("Look into pointers (var ip *int)", "pointer var address:", ipp, "pointer:", ip, "pointed value:", *ip)

	var xsp *[]int
	fmt.Println("Look into pointers (var xsp *[]int)", "pointer var address:", &xsp, "pointer:", xsp, "pointed value:", "")
	xsp = new([]int)
	fmt.Println("Look into pointers (var xsp *[]int)", "pointer var address:", &xsp, "pointer:", xsp, "pointed value:", *xsp)
	var xspp *[]int
	tmp := append(*xsp, 42)
	xspp = &tmp
	fmt.Println("Look into pointers (xspp := append(*xsp, 42)", "pointer var address:", &xspp, "pointer:", xspp, "pointed value:", *xspp)

}

func lookToMap() {
	return
	map_ := make(map[int]int)

	fmt.Println("Look to map (new map):", unsafe.Pointer(&map_), map_)
	lookToMapDownStream(map_)
	fmt.Println("Look to map (down stream return):", unsafe.Pointer(&map_), map_)
}

func lookToMapDownStream(in map[int]int) {
	in[42] = 24
	fmt.Println("Look to map (down stream):", unsafe.Pointer(&in), in)
}

func lookToSlice() {
	slice := []int{42}
	fmt.Println("Look to slice (new slice):", unsafe.Pointer(&slice), cap(slice), slice)
	lookToSliceDownStream(slice, 0)
	fmt.Println("Look to slice (down stream return):", unsafe.Pointer(&slice), cap(slice), slice)
}

func lookToSliceDownStream(in []int, i int) {
	in[i] += 1
	in = append(in, 10)
	in[i] = 124
	fmt.Println("Look to slice (down stream):", unsafe.Pointer(&in), cap(in), in)
}
