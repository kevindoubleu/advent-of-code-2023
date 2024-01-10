package day01

import (
	"testing"
)

func Test_getFirstDigitWithSpelled(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy spelled",
			args: args{"two1nine"},
			want: 2,
		}, {
			name: "prefixed",
			args: args{"abcone2threexyz"},
			want: 1,
		}, {
			name: "not spelled",
			args: args{"4nineeightseven2"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstDigitWithSpelled(tt.args.line); got != tt.want {
				t.Errorf("getFirstDigitWithSpelled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastDigitWithSpelled(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy spelled",
			args: args{"two1nine"},
			want: 9,
		}, {
			name: "suffixed",
			args: args{"abcone2threexyz"},
			want: 3,
		}, {
			name: "not spelled",
			args: args{"4nineeightseven2"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastDigitWithSpelled(tt.args.line); got != tt.want {
				t.Errorf("getLastDigitWithSpelled() = %v, want %v", got, tt.want)
			}
		})
	}
}
