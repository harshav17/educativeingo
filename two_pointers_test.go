package educativeingo

import (
	"reflect"
	"testing"
)

func Test_pairWithTargerSum(t *testing.T) {
	type args struct {
		arr       []int
		targerSum int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sanity2",
			args: args{[]int{1, 2, 3, 4, 6}, 6},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairWithTargerSum(tt.args.arr, tt.args.targerSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairWithTargerSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity2",
			args: args{[]int{2, 3, 3, 3, 6, 9, 9}},
			want: 4,
		},
		{
			name: "sanity2",
			args: args{[]int{2, 2, 2, 11}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.arr); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sorterdArraySquares(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sanity2",
			args: args{[]int{-2, -1, 0, 2, 3}},
			want: []int{0, 1, 4, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sorterdArraySquares(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sorterdArraySquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tripletSumToZero(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
	}{
		{
			name: "sanity2",
			args: args{[]int{-3, 0, 1, 2, -1, 1, -2}},
			want: [][3]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tripletSumToZero(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tripletSumToZero() = %v, want %v", got, tt.want)
			}
		})
	}
}