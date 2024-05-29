package main_test

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	testcases := map[string]func(*testing.B){
		"append without zero cap and zero len alloc": func(b *testing.B) {
			s := make([]int, 0)
			for n := range 1_000 {
				s = append(s, n)
			}
		},
		"append with cap but zero len alloc": func(b *testing.B) {
			s := make([]int, 1_000)
			for n := range 1_000 {
				s = append(s, n)
			}
		},
		"append with cap and len alloc": func(b *testing.B) {
			s := make([]int, 1_000, 1_000)
			for n := range 1_000 {
				s = append(s, n)
			}
		},
		"idx append with cap and len alloc": func(b *testing.B) {
			s := make([]int, 1_000, 1_000)
			for n := range 1_000 {
				s[n] = n
			}
		},
	}

	for range b.N {
		for name, test := range testcases {
			b.Run(name, test)
		}
	}
}
