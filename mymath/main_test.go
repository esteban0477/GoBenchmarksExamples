package mymath

import "testing"

type test struct {
	xi  []int
	avg float64
}

func TestCenteredAvg(t *testing.T) {
	tableTest := []test{
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{1, 2, 3}, 2},
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{1, 2, 3, 4, 5}, 3},
	}

	for _, value := range tableTest {
		result := CenteredAvg(value.xi)
		if result != value.avg {
			t.Errorf("Got %v, expected %v", result, value.avg)
		}

	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5})
	}
}
