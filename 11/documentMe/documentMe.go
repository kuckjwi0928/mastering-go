// Package documentMe 이것은 documentMe 패키지이다.
package documentMe

// Pie Pie는 전역변수 입니다.
const Pie = 3.1415912

// S1 이 함수는 스트링의 길이를 구해온다.
func S1(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}
