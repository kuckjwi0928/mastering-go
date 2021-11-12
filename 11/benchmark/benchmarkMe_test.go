package benchmark

import "testing"

var result int

func Benchmark_fibonacci1(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci1(30)
	}
	result = r
}

func Benchmark_fibonacci2(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci2(30)
	}
	result = r
}
