package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	if result := Years(5); result != 35 {
		t.Errorf("Got: %v, expected 35", result)
	}
}

func TestYearsTwo(t *testing.T) {
	if result := YearsTwo(5); result != 35 {
		t.Errorf("Got: %v, expected 35", result)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(Years(5))
	// Output:
	// 35
}
