package educativeingo

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "sanity",
			args: args{[][2]int{{1, 4}, {2, 6}, {3, 5}}},
			want: [][2]int{{1, 6}},
		},
		{
			name: "sanity2",
			args: args{[][2]int{{6, 7}, {2, 4}, {5, 9}}},
			want: [][2]int{{2, 4}, {5, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertIntervals(t *testing.T) {
	type args struct {
		intervals   [][2]int
		newInterval [2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "sanity",
			args: args{
				[][2]int{{1, 3}, {5, 7}, {8, 12}},
				[2]int{4, 6},
			},
			want: [][2]int{{1, 3}, {4, 7}, {8, 12}},
		},
		{
			name: "sanity",
			args: args{
				[][2]int{{1, 3}, {5, 7}, {8, 12}},
				[2]int{4, 10},
			},
			want: [][2]int{{1, 3}, {4, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntervals(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
