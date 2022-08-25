package sample

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	for _, tc := range []struct {
		name string
		i    int
		e    string
	}{
		{
			"test 0",
			0,
			"0",
		},
		{
			"test negative",
			-12,
			"-12",
		},
		{
			"test positive",
			+12,
			"12",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := IntToString(tc.i)
			if r != tc.e {
				t.Fatalf("Unexpected int2string conversion result: %v != %v", r, tc.e)
			}
		})
	}
}