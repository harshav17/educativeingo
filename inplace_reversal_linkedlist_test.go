package educativeingo

import (
	"reflect"
	"testing"
)

func Test_reverseASubList(t *testing.T) {
	type args struct {
		root *node
		p    int
		q    int
	}

	input := node{7, &node{6, &node{5, &node{4, &node{3, &node{2, &node{1, nil}}}}}}}
	output := node{7, &node{4, &node{5, &node{6, &node{3, &node{2, &node{1, nil}}}}}}}

	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "sanity",
			args: args{root: &input, p: 2, q: 4},
			want: &output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseASubList(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseASubList() = %v, want %v", got, tt.want)
			}
		})
	}
}
