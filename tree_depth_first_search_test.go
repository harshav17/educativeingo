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

func TestFindSumOfPathNumbers(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input1 := treeNode{1, &treeNode{7, nil, nil}, &treeNode{9, &treeNode{2, nil, nil}, &treeNode{9, nil, nil}}}
	input2 := treeNode{1, &treeNode{0, nil, &treeNode{1, nil, nil}}, &treeNode{1, &treeNode{6, nil, nil}, &treeNode{5, nil, nil}}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{root: &input1},
			want: 408,
		},
		{
			name: "sanity",
			args: args{root: &input2},
			want: 332,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSumOfPathNumbers(tt.args.root); got != tt.want {
				t.Errorf("FindSumOfPathNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPathWGivenSequence(t *testing.T) {
	type args struct {
		root *treeNode
		seq  []int
	}

	input1 := treeNode{1, &treeNode{7, nil, nil}, &treeNode{9, &treeNode{2, nil, nil}, &treeNode{9, nil, nil}}}
	input2 := treeNode{1, &treeNode{0, nil, &treeNode{1, nil, nil}}, &treeNode{1, &treeNode{6, nil, nil}, &treeNode{5, nil, nil}}}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sanity",
			args: args{root: &input1, seq: []int{1, 9, 9}},
			want: true,
		},
		{
			name: "sanity",
			args: args{root: &input2, seq: []int{1, 0, 7}},
			want: false,
		},
		{
			name: "sanity",
			args: args{root: &input2, seq: []int{1, 1, 6}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPathWGivenSequence(tt.args.root, tt.args.seq); got != tt.want {
				t.Errorf("FindPathWGivenSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPathsOfSum(t *testing.T) {
	type args struct {
		root *treeNode
		sum  int
	}

	input1 := treeNode{1, &treeNode{7, &treeNode{6, nil, nil}, &treeNode{5, nil, nil}}, &treeNode{9, &treeNode{2, nil, nil}, &treeNode{3, nil, nil}}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{root: &input1, sum: 12},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPathsOfSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("CountPathsOfSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDiameter(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input1 := treeNode{1, &treeNode{2, &treeNode{4, nil, nil}, nil}, &treeNode{3, &treeNode{5, nil, nil}, &treeNode{6, nil, nil}}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{root: &input1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDiameter(tt.args.root); got != tt.want {
				t.Errorf("FindDiameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfPathWMaxSum(t *testing.T) {
	type args struct {
		root *treeNode
	}

	input1 := treeNode{1, &treeNode{2, nil, nil}, &treeNode{3, nil, nil}}
	input2 := treeNode{1, &treeNode{2, &treeNode{1, nil, nil}, &treeNode{3, nil, nil}}, &treeNode{3, &treeNode{5, &treeNode{7, nil, nil}, &treeNode{8, nil, nil}}, &treeNode{6, &treeNode{9, nil, nil}, nil}}}
	input3 := treeNode{-1, &treeNode{-3, nil, nil}, nil}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{root: &input1},
			want: 6,
		},
		{
			name: "sanity",
			args: args{root: &input2},
			want: 31,
		},
		{
			name: "sanity",
			args: args{root: &input3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfPathWMaxSum(tt.args.root); got != tt.want {
				t.Errorf("SumOfPathWMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
