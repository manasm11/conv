package conv

import "testing"

func TestToInt(t *testing.T) {
	tests := []struct {
		str    string
		output int
	}{
		{"0", 0},
		{"1", 1},
		{"123", 123},
		{"893177376", 893177376},
		{"-314", -314},
	}
	for i, test := range tests {
		if o := ToInt(test.str); o != test.output {
			t.Errorf("FAILED (%v): Excpected %v, got %v", i, test.output, o)
		}
	}
}

func TestToFloat(t *testing.T) {
	tests := []struct {
		str    string
		output float64
	}{
		{"0.00", 0},
		{"1.00", 1},
		{"1.23", 1.23},
		{"8.93177376", 8.93177376},
		{"-314.43", -314.43},
	}
	for i, test := range tests {
		if o := ToFloat(test.str); o != test.output {
			t.Errorf("FAILED (%v): Excpected %v, got %v", i, test.output, o)
		}
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		str    string
		output bool
	}{
		{"true", true},
		{"false", false},
	}
	for i, test := range tests {
		if o := ToBool(test.str); o != test.output {
			t.Errorf("FAILED (%v): Excpected %v, got %v", i, test.output, o)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		v      any
		output string
	}{
		{true, "true"},
		{false, "false"},
		{3, "3"},
		{3.65, "3.65"},
		{-123, "-123"},
		{struct {
			a int
		}{1}, "struct { a int }{a:1}"},
	}
	for i, test := range tests {
		if o := ToString(test.v); o != test.output {
			t.Errorf("FAILED (%v): Excpected %v, got %v", i, test.output, o)
		}
	}
}
