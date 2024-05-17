package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"Negative", -1, 1},
		{"Zero", 0, 0},
		{"Positive", 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.arg); got != tt.want {
				t.Errorf("Abs(%d) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
