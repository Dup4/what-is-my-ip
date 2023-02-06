package version

import "testing"

func TestToString(t *testing.T) {
	if len(ToString()) == 0 {
		t.Errorf("ToString() should not be empty")
	}
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToString()
	}
}
