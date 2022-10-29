package pystrings

import (
	"testing"
)

func TestRsplit(t *testing.T) {
	var tests = []struct {
		S   string
		Sep string
		Sub []string
	}{
		{
			S:   "a-b-c",
			Sep: "-",
			Sub: []string{"a", "b", "c"},
		},
		{
			S:   "a",
			Sep: "-",
			Sub: []string{"a"},
		},
		{
			S:   "a--b",
			Sep: "-",
			Sub: []string{"a", "", "b"},
		},
	}
	for i := 0; i < len(tests); i++ {
		s := tests[i].S
		subs := tests[i].Sub
		sep := tests[i].Sep
		rsubs := RSplit(s, sep)

		for j := 0; j < len(rsubs); j++ {
			if rsubs[j] != subs[j] {
				t.Fatalf("%s expect: %v but get: %v", s, subs, rsubs)
			}
		}
	}
}

func TestRsplitN(t *testing.T) {
	var tests = []struct {
		S   string
		Sep string
		N   int
		Sub []string
	}{
		{
			S: "a-b-c",
			Sep: "+	",
			N:   0,
			Sub: []string{"a-b-c"},
		},
		{
			S:   "a-b-c",
			Sep: "-",
			N:   0,
			Sub: []string{"a-b-c"},
		},
		{
			S:   "a-b-c",
			Sep: "-",
			N:   1,
			Sub: []string{"a-b", "c"},
		},
		{
			S:   "a-b-c",
			N:   2,
			Sep: "-",
			Sub: []string{"a", "b", "c"},
		},
		{
			S:   "a-b-c",
			N:   3,
			Sep: "-",
			Sub: []string{"a", "b", "c"},
		},
		{
			S:   "a--b",
			Sep: "-",
			N:   0,
			Sub: []string{"a--b"},
		},
		{
			S:   "a--b",
			Sep: "-",
			N:   1,
			Sub: []string{"a-", "b"},
		},
		{
			S:   "a--b",
			Sep: "-",
			N:   2,
			Sub: []string{"a", "", "b"},
		},
		{
			S:   "a--b",
			Sep: "-",
			N:   3,
			Sub: []string{"a", "", "b"},
		},
	}
	for i := 0; i < len(tests); i++ {
		s := tests[i].S
		subs := tests[i].Sub
		sep := tests[i].Sep
		rsubs := RSplitN(s, sep, tests[i].N)

		for j := 0; j < len(rsubs); j++ {
			if rsubs[j] != subs[j] {
				t.Fatalf("%s expect: %v but get: %v", s, subs, rsubs)
			}
		}
	}
}
