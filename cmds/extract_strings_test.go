package cmds

import (
	"reflect"
	"testing"
)

func TestMergeLocales(t *testing.T) {
	tests := []struct {
		a   []string
		b   []string
		exp []string
	}{
		{
			// result should be the union of locales (a disjoint from b)
			a:   []string{"en"},
			b:   []string{"fr"},
			exp: []string{"en", "fr"},
		},
		{
			// result should be the union of locales (a subset of b)
			a:   []string{"fr"},
			b:   []string{"en", "fr"},
			exp: []string{"en", "fr"},
		},
		{
			// if one of the entries has no locales, result should have no locales
			a:   nil,
			b:   []string{"en", "fr"},
			exp: nil,
		},
		{
			// if one of the entries has no locales, result should have no locales
			a:   []string{"en"},
			b:   nil,
			exp: nil,
		},
	}
	for _, test := range tests {
		got := mergeLocales(test.a, test.b)
		if !reflect.DeepEqual(got, test.exp) {
			t.Fatalf("got %v, expected %v", got, test.exp)
		}
	}
}
