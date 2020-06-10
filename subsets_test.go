package educativeingo

import (
	"reflect"
	"testing"
)

func TestFindSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sanity",
			args: args{nums: []int{1, 5, 3}},
			want: [][]int{{}, {1}, {5}, {1, 5}, {3}, {1, 3}, {5, 3}, {1, 5, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
