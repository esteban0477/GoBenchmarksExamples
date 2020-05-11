package word

import (
	"fmt"
	"quote"
	"testing"
)

func TestUseCount(t *testing.T) {
	returnedMap := UseCount(quote.SunAlso)
	if returnedMap["was"] != 32 {
		t.Error("Something didn't work as expected!")
	}
}

func TestCount(t *testing.T) {
	returnedCounter := Count(quote.SunAlso)
	if returnedCounter != 1349 {
		t.Errorf("Expected 1349, Got %v", returnedCounter)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// Output:
	// 1349
}
