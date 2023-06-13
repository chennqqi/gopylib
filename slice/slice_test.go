package slice

import (
	"testing"
)

func TestContains(t *testing.T) {
	{
		teststrings := []string{"a", "b", "sdfsdfsf"}
		ss := Slice[string](teststrings)
		for i := 0; i < len(ss); i++ {
			if !ss.Contains(ss[i]) {
				t.Fatal("expect true")
			}
		}
		if ss.Contains("NOT EXISTS") {
			t.Fatal("expect false")
		}

	}

	{
		testint := []int{1, 2, 3, 10000}
		ss := Slice[int](testint)
		for i := 0; i < len(ss); i++ {
			if !ss.Contains(ss[i]) {
				t.Fatal("expect true")
			}
		}
		if ss.Contains(1211111) {
			t.Fatal("expect false")
		}
	}
}
