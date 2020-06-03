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

func Test_zigzagLevelOrderTrav(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, &treeNode{6, nil, nil}, &treeNode{7, nil, nil}}}
	output := [][]int{{1}, {3, 2}, {4, 5, 6, 7}}

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
			if got := zigzagLevelOrderTrav(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrderTrav() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLevelAverages(t *testing.T) {
	type args struct {
		root *treeNode
	}
	input := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, &treeNode{6, nil, nil}, &treeNode{7, nil, nil}}}
	output := []float32{1, 2.5, 5.5}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "sanity",
			args: args{&input},
			want: output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLevelAverages(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLevelAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinimumDepth(t *testing.T) {
	type args struct {
		root *treeNode
	}
	input := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, nil, nil}}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{&input},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumDepth(tt.args.root); got != tt.want {
				t.Errorf("findMinimumDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
