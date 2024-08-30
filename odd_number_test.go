package main

import "testing"

// Write your test here

func TestFindOddNumber(t *testing.T) {
	type args struct {
		text []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1st case",
			args: args{[]int{7}},
			want: 7,
		},
		{
			name: "2nd case",
			args: args{[]int{0}},
			want: 0,
		},
		{
			name: "3rd case",
			args: args{[]int{1, 1, 2}},
			want: 2,
		},
		{
			name: "4th case",
			args: args{[]int{0, 1, 0, 1, 0}},
			want: 0,
		},
		{
			name: "5th case",
			args: args{[]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOddNumber(tt.args.text); got != tt.want {
				t.Errorf("FindOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
