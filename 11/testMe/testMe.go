package testMe

func f1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return f1(n-1) + f1(n-2)
}

func s1(s string) int {
	if s == "" {
		return 0
	}
	n := 1
	for range s {
		n++
	}
	return n
}
