package educativeingo

import (
	"reflect"
	"testing"
)

func Test_cyclicSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sanity",
			args: args{[]int{3, 1, 5, 4, 2}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sanity",
			args: args{[]int{2, 6, 4, 3, 1, 5}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "sanity",
			args: args{[]int{1, 5, 6, 4, 3, 2}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cyclicSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cyclicSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{[]int{4, 0, 3, 1}},
			want: 2,
		},
		{
			name: "sanity",
			args: args{[]int{8, 3, 5, 2, 4, 6, 0, 1}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("findMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAllMissingNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sanity",
			args: args{[]int{2, 3, 1, 8, 2, 3, 5, 1}},
			want: []int{4, 6, 7},
		},
		{
			name: "sanity",
			args: args{[]int{2, 4, 1, 2}},
			want: []int{3},
		},
		{
			name: "sanity",
			args: args{[]int{2, 3, 2, 1}},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllMissingNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllMissingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{[]int{1, 4, 4, 3, 2}},
			want: 4,
		},
		{
			name: "sanity",
			args: args{[]int{2, 1, 3, 3, 5, 4}},
			want: 3,
		},
		{
			name: "sanity",
			args: args{[]int{2, 4, 1, 4, 4}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAllDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sanity",
			args: args{[]int{3, 4, 4, 5, 5}},
			want: []int{5, 4},
		},
		{
			name: "sanity",
			args: args{[]int{5, 4, 7, 2, 3, 5, 3}},
			want: []int{3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
