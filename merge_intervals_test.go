package educativeingo

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "sanity",
			args: args{[][]int{{1, 4}, {2, 6}, {3, 5}}},
			want: [][2]int{{1, 6}},
		},
		{
			name: "sanity2",
			args: args{[][]int{{6, 7}, {2, 4}, {5, 9}}},
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

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		interval1 [][2]int
		interval2 [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "sanity",
			args: args{
				[][2]int{{1, 3}, {5, 6}, {7, 9}},
				[][2]int{{2, 3}, {5, 7}},
			},
			want: [][2]int{{2, 3}, {5, 6}, {7, 7}},
		},
		{
			name: "sanity",
			args: args{
				[][2]int{{1, 3}, {5, 7}, {9, 12}},
				[][2]int{{5, 10}},
			},
			want: [][2]int{{5, 7}, {9, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.interval1, tt.args.interval2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_conflictingAppointments(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sanity",
			args: args{[][]int{{1, 4}, {2, 6}, {3, 5}}},
			want: false,
		},
		{
			name: "sanity",
			args: args{[][]int{{3, 4}, {1, 2}, {5, 6}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conflictingAppointments(tt.args.intervals); got != tt.want {
				t.Errorf("conflictingAppointments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{[][]int{{1, 4}, {2, 6}, {3, 5}}},
			want: 3,
		},
		{
			name: "sanity",
			args: args{[][]int{{1, 4}, {2, 5}, {7, 9}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minimumMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxCPULoad(t *testing.T) {
	type args struct {
		jobs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{[][]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}},
			want: 7,
		},
		{
			name: "sanity",
			args: args{[][]int{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}}},
			want: 15,
		},
		{
			name: "sanity",
			args: args{[][]int{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxCPULoad(tt.args.jobs); got != tt.want {
				t.Errorf("findMaxCPULoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
