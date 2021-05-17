package main

import "testing"

func BenchmarkHexdiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hexdiff("3024a7c5ecb289adcaa00a06082a8648", "3002e3646fc42be9f24c6dcaa1b43b29")
	}
}
