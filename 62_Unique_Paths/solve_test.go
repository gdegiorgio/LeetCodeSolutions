package main

import (
	"fmt"
	"testing"
)

func Test_UniquePaths(t *testing.T) {
	var tests = []struct {
		m, n int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.m, tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := UniquePaths(tt.m, tt.n)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
