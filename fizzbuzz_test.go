package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	type checkFunc func([]int) error

	tests := [...]struct {
		in 	int
		check	checkFunc
	}{
		{1, isEmptyList}
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Factors of %d", tc.in), func(t *testing.T) {
			factors := factorsOf(tc.in)
			if err := tc.check(factors); err != nil {
				t.Error(err)
			}
		})
	}
}
