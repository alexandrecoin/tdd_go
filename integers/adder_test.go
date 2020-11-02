package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("it returns 4 when adding 2 and 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("Expected %d but got %d", expected, sum)
		}
	})

	t.Run("it returns 3 when adding 1 and 2", func(t *testing.T) {
		sum := Add(1, 2)
		expected := 3
		if sum != expected {
			t.Errorf("Expected %d but got %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
