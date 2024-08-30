package main

import "testing"

func TestCountSmilyFace(t *testing.T) {
	type args struct {
		text []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1st Case",
			args: args{[]string{":)", ";(", ";}", ":-D"}},
			want: 2,
		},
		{
			name: "2nd Case",
			args: args{[]string{";D", ":-(", ":-)", ";~)"}},
			want: 3,
		},
		{
			name: "3rd Case",
			args: args{[]string{";]", ":[", ";*", ":$", ";-D"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSmilyFace(tt.args.text); got != tt.want {
				t.Errorf("CountSmilyFace() = %v, want %v", got, tt.want)
			}
		})
	}
}
