func shortestPalindrome(s string) string {
    if s==""{
        return s
    }
	l := 0
	r := len(s) - 1
	for r >= 0 {
		flag := false
        l = 0
        temp:=r
		for l < r {
			if s[l] != s[r] {
				flag = true
				break
			}
			l++
            r--
		}
		if flag == false {
            r= temp +1
			break
		}
		temp--
        r=temp
	}
    a:= s
	for r < len(s) {
		a = string(s[r]) + a
		r++
	}
	return a
} 