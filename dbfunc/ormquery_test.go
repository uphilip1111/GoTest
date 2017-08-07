package dbfunc

import "testing"

func BenchmarkQueryByORM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryByORM()
	}
}
