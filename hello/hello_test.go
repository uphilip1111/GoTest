package hello

import "testing"

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}
