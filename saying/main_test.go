package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	returnedString := Greet("Juan")
	if returnedString != "Hello my dear loquita Juan" {
		t.Error("Got", returnedString, "Expected: Hello my dear loquita Juan")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Juan"))
	// Output:
	// Hello my dear loquita Juan
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Juan")
	}
}
