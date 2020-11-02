package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("it return 'aaaaa' when given a as param", func (t *testing.T) {
		repeated := Repeat("a", 5)
		expected := ("aaaaa")
		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("it returns 'bbb' when given b as param", func (t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"
		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
