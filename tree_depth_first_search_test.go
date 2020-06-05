package educativeingo

import "testing"

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
