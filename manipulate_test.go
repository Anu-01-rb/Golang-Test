package main

import (
	"reflect"
	"testing"
)

func TestManipulate(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "1st case",
			args: args{"a"},
			want: []string{"a"},
		},
		{
			name: "2nd case",
			args: args{"ab"},
			want: []string{"ab", "ba"},
		},
		{
			name: "3rd case",
			args: args{"abc"},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "4th case",
			args: args{"aabb"},
			want: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Manipulate(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manipulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
