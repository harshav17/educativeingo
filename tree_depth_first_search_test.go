package educativeingo

import (
	"reflect"
	"testing"
)

func Test_hasPath(t *testing.T) {
	type args struct {
		root *treeNode
		sum  int
	}

	input1 := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, &treeNode{6, nil, nil}, &treeNode{7, nil, nil}}}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sanity",
			args: args{root: &input1, sum: 10},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPath(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllPaths(t *testing.T) {
	type args struct {
		root *treeNode
		sum  int
	}

	input1 := treeNode{1, &treeNode{7, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{9, &treeNode{2, nil, nil}, &treeNode{7, nil, nil}}}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sanity",
			args: args{root: &input1, sum: 12},
			want: [][]int{{1, 7, 4}, {1, 9, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllPaths(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
