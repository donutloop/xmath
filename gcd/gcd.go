package gcd

func Eulidean(a, b int) int {

	var min, max int

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	r := 0
	d := 0
	for {
		r = max % min
		if r == 0 {
			return min
		}

		d = (max - r) / min
		if max != (d*min)+r {
			panic("Eulidean is broken")
		}

		max = min
		min = r
	}

	return -1
}
