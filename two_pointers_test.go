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
		// TODO: Add test cases.2, 2, 2, 11
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
