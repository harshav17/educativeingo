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

func Test_findSuccessor(t *testing.T) {
	type args struct {
		root *treeNode
		key  int
	}

	input := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, nil, nil}}

	tests := []struct {
		name string
		args args
		want *treeNode
	}{
		{
			name: "sanity",
			args: args{root: &input, key: 3},
			want: &treeNode{4, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSuccessor(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectLevelOrderSiblings(t *testing.T) {
	type args struct {
		root *treeNodeWNext
	}

	input := treeNodeWNext{1, &treeNodeWNext{2, &treeNodeWNext{4, nil, nil, nil}, &treeNodeWNext{5, nil, nil, nil}, nil}, &treeNodeWNext{3, &treeNodeWNext{6, nil, nil, nil}, &treeNodeWNext{7, nil, nil, nil}, nil}, nil}

	tests := []struct {
		name string
		args args
		want *treeNodeWNext
	}{
		{
			name: "sanity",
			args: args{root: &input},
			want: &input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectLevelOrderSiblings(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectLevelOrderSiblings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectAllSiblings(t *testing.T) {
	type args struct {
		root *treeNodeWNext
	}

	input := treeNodeWNext{1, &treeNodeWNext{2, &treeNodeWNext{4, nil, nil, nil}, &treeNodeWNext{5, nil, nil, nil}, nil}, &treeNodeWNext{3, &treeNodeWNext{6, nil, nil, nil}, &treeNodeWNext{7, nil, nil, nil}, nil}, nil}

	tests := []struct {
		name string
		args args
		want *treeNodeWNext
	}{
		{
			name: "sanity",
			args: args{root: &input},
			want: &input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectAllSiblings(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectAllSiblings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeRightView(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input1 := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, nil, nil}}
	input2 := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{3, &treeNode{6, nil, nil}, &treeNode{7, nil, nil}}}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sanity",
			args: args{root: &input1},
			want: []int{1, 3, 5},
		},
		{
			name: "sanity",
			args: args{root: &input2},
			want: []int{1, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeRightView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeRightView() = %v, want %v", got, tt.want)
			}
		})
	}
}
