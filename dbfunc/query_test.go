package dbfunc

import (
	"testing"
)

func BenchmarkQueryFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryFunc()
	}
}
