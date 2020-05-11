package main

import (
	"fmt"
	"testing"
)

type test struct {
	numbers []int
	result  int
}

func TestMySum(t *testing.T) {

	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{-123, 123, 0}, 0},
		test{[]int{-1, -1, -1}, -3},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
	}

	for _, value := range tests {
		result := mySum(value.numbers...)
		if result != value.result {
			t.Errorf("Expected result %v, Got: %v", value.result, result)
		}
	}
}

func ExamplemySum() {
	fmt.Println(mySum(2, 3))
	// Output:
	// 7
}
