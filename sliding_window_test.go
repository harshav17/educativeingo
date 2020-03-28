package educativeingo

import (
	"testing"
)

func Test_maxSubArrayOfSizeK(t *testing.T) {
	type args struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "sanity",
			args: args{3, []int{2, 1, 5, 1, 3, 2}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayOfSizeK(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("maxSubArrayOfSizeK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestSubarrayWithGivenSum(t *testing.T) {
	type args struct {
		s   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{7, []int{2, 1, 5, 2, 3, 2}},
			want: 2,
		},
		{
			name: "sanity",
			args: args{7, []int{2, 1, 5, 2, 8}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubarrayWithGivenSum(tt.args.s, tt.args.arr); got != tt.want {
				t.Errorf("smallestSubarrayWithGivenSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubsWithKDistChars(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{"araaci", 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsWithKDistChars(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("longestSubsWithKDistChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noRepeatSubstring(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{"aabccbb"},
			want: 3,
		},
		{
			name: "abbbb",
			args: args{"abbbb"},
			want: 2,
		},
		{
			name: "abccde",
			args: args{"abccde"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noRepeatSubstring(tt.args.str); got != tt.want {
				t.Errorf("noRepeatSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_characterReplacement(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{"aabccbb", 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}