package dbfunc

import (
	"testing"
)

func BenchmarkQueryBySQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryBySQL()
	}
}
