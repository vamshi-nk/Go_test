package iteration

import "testing"

func BenchmarkRepeated(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
