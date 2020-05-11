package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	returnedString := Greet("Juan")
	if returnedString != "Hello my dear Juan" {
		t.Error("Got", returnedString, "Expected: Hello my dear Juan")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Juan"))
	// Output:
	// Hello my dear Juan
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Juan")
	}
}
