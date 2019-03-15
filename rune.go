package faker

func RuneBetween(a rune, b rune) rune {
	return rune(IntBetween(int(a), int(b)))
}

func RunesBetween(n int, a rune, b rune) []rune {
	buffer := make([]rune, n)
	for i := 0; i < n; i++ {
		buffer[i] = RuneBetween(a, b)
	}
	return buffer
}

func RunesIn(n int, r []rune) []rune {
	b := make([]rune, n)
	for i := 0; i < n; i++ {
		b[i] = RuneIn(r)
	}
	return b
}

func RuneIn(r []rune) rune {
	return r[Intn(len(r))]
}
