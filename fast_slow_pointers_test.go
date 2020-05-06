package educativeingo

import (
	"reflect"
	"testing"
)

func Test_linkedListCycle(t *testing.T) {
	type args struct {
		root *node
	}

	one := node{1, nil}
	two := node{2, &one}
	three := node{3, &two}
	four := node{4, &three}
	five := node{5, &four}
	six := node{6, &five}

	//loop back to three
	one.next = &three

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sanity",
			args: args{&six},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.root); got != tt.want {
				t.Errorf("linkedListCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startOfCycle(t *testing.T) {
	type args struct {
		root *node
	}

	one := node{1, nil}
	two := node{2, &one}
	three := node{3, &two}
	four := node{4, &three}
	five := node{5, &four}
	six := node{6, &five}

	//loop back to three
	one.next = &three

	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "sanity",
			args: args{&six},
			want: &three,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfCycle(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("startOfCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHappyNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sanity",
			args: args{23},
			want: true,
		},
		{
			name: "sanity",
			args: args{12},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappyNumber(tt.args.num); got != tt.want {
				t.Errorf("isHappyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}