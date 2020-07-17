package prison

func prisonAfterNDays(cells []int, N int) []int {
	if N < 1 {
		return cells
	}

	var b byte = 0
	var w byte = 128
	for _, c := range cells {
		b += w * byte(c)
		w = w >> 1
	}

	dump := func(b byte) []int {
		w = 128
		tmp := make([]int, len(cells))
		for i := 0; i < len(cells); i++ {
			if b&w == w {
				tmp[i] = 1
			} else {
				tmp[i] = 0
			}
			w = w >> 1
		}
		return tmp
	}

	var i byte = 0
	j := N
	d := b
	s := b

	for {
		if j < 1 {
			return dump(^b)
		}

		b = ((b >> 1) ^ (b << 1)) | 0x81

		i++
		j--

		if i == 1 {
			s = b
		}

		if i > 1 && b == s {
			break
		}
	}

	k := ((N - 1) % (int(i) - 1)) + 1
	for k > 0 {
		d = ((d >> 1) ^ (d << 1)) | 0x81
		k--
	}

	return dump(^d)
}

func prisonAfterNDaysSlow(cells []int, N int) []int {
	work := make([]int, len(cells))
	temp := make([]int, len(cells))
	copy(temp, cells)

	k := 1
	for j := 0; j < len(temp)-2; j++ {
		if temp[j] == temp[j+2] {
			work[k] = 1
		} else {
			work[k] = 0
		}
		k++
	}
	work[0], work[len(work)-1] = 0, 0
	first := make([]int, len(cells))
	copy(temp, work)
	copy(first, work)

	cycles := make([][]int, 0, 16)
	cycles = append(cycles, first)
	repeatedPeriod := 0
	for i := 1; i < N; i++ {
		k := 1
		for j := 0; j < len(temp)-2; j++ {
			if temp[j] == temp[j+2] {
				work[k] = 1
			} else {
				work[k] = 0
			}
			k++
		}
		work, temp = temp, work
		repeated := true
		for j, t := range temp {
			if t != first[j] {
				repeated = false
				break
			}
		}
		if repeated {
			repeatedPeriod = i
			break
		}
		copied := make([]int, len(temp))
		copy(copied, temp)
		cycles = append(cycles, copied)
	}

	if repeatedPeriod == 0 {
		return temp
	}
	return cycles[N%repeatedPeriod-1]
}
