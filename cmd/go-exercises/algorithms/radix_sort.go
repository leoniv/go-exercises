package algorithms

func RadixSort(in []uint32) []uint32 {
	const radix = 256
	n := len(in)
	out := make([]uint32, n)

	for k := range 4 {
		stat := new([radix]int)
		// Calculate stat
		for _, word := range in {
			stat[getByte(word, k)] += 1
		}
		// Calculate prefix sum
		for i := 1; i < radix; i++ {
			stat[i] += stat[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			b := getByte(in[i], k)
			stat[b]--
			out[stat[b]] = in[i]
		}

		copy(in, out)
	}

	return in
}

func getByte(in uint32, n int) byte {
	ret := (in >> (n * 8)) & 0xFF
	return byte(ret)
}
