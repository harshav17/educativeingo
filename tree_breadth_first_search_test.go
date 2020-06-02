package educativeingo

import (
	"reflect"
	"testing"
)

func Test_levelOrderTraversal(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, &treeNode{6, nil, nil}, &treeNode{7, nil, nil}}}
	output := [][]int{{1}, {2, 3}, {4, 5, 6, 7}}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sanity",
			args: args{&input},
			want: output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseLevelOrderTrav(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, &treeNode{6, nil, nil}, &treeNode{7, nil, nil}}}
	output := [][]int{{4, 5, 6, 7}, {2, 3}, {1}}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sanity",
			args: args{&input},
			want: output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLevelOrderTrav(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLevelOrderTrav() = %v, want %v", got, tt.want)
			}
		})
	}
}
