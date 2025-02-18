func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	// cp := 0
	// fp := 1
	num := m[s[len(s)-1]]
	prev := s[len(s)-1]
	for pos := len(s) - 2; pos >= 0; pos-- {
		x := s[pos]
		if (prev == 'V' || prev == 'X') && x == 'I' {
			num -= m[x]
		} else if (prev == 'L' || prev == 'C') && x == 'X' {
			num -= m[x]
		} else if (prev == 'D' || prev == 'M') && x == 'C' {
			num -= m[x]
		} else {
			num += m[x]
			prev = x
		}

	}
	return num
}