package main

import (
	"reflect"
	"testing"
)

func Test_reverseSlice(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Reversing a list of negative numbers",
			args: args{input: []int{-1, -2, -3, -4, -5}},
			want: []int{-5, -4, -3, -2, -1},
		},
		{
			name: "Reversing a list of mixed numbers",
			args: args{input: []int{3, -1, 2, -4, 5}},
			want: []int{5, -4, 2, -1, 6},
		},
		{
			name: "Reversing a single element list",
			args: args{input: []int{1}},
			want: []int{1, 2},
		},
		{
			name: "Reversing an empty list",
			args: args{input: []int{}},
			want: []int{},
		},
	}

	failureOccurred := false

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseSlice(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				failureOccurred = true
				t.Error("reverseSlice() =", got, ", want", tt.want)
				// t.Errorf("reverseSlice() = %v, want %v", got, tt.want)
			}
		})

		// If a failure occurred, stop further test execution
		if failureOccurred {
			t.Fatal("Stopping further test execution due to failure")
		}
	}
}
