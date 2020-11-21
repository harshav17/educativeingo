package educativeingo

import (
	"testing"
)

func Test_zeroOneKnapsack(t *testing.T) {
	type args struct {
		profits  []int
		weights  []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroOneKnapsack(tt.args.profits, tt.args.weights, tt.args.capacity); got != tt.want {
				t.Errorf("zeroOneKnapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{30},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFibTopDown(tt.args.n); got != tt.want {
				t.Errorf("findFib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFindFibTopDown(n int, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findFibTopDown(n)
	}
}

func BenchmarkFibTopDown2(b *testing.B) {
	benchmarkFindFibTopDown(5, b)
}

func BenchmarkFibTopDown3(b *testing.B) {
	benchmarkFindFibTopDown(10, b)
}

func BenchmarkFibTopDown4(b *testing.B) {
	benchmarkFindFibTopDown(20, b)
}

func BenchmarkFibTopDown5(b *testing.B) {
	benchmarkFindFibTopDown(30, b)
}

func BenchmarkFibTopDown6(b *testing.B) {
	benchmarkFindFibTopDown(50, b)
}

func BenchmarkFibTopDown7(b *testing.B) {
	benchmarkFindFibTopDown(100, b)
}

func BenchmarkFibTopDown8(b *testing.B) {
	benchmarkFindFibTopDown(1000, b)
}

func Test_fibNaiveRec(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{30},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibNaiveRec(tt.args.n); got != tt.want {
				t.Errorf("fibNaiveRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFindFibNaiveRec(n int, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fibNaiveRec(n)
	}
}

func BenchmarkFibNaiveRec2(b *testing.B) {
	benchmarkFindFibNaiveRec(5, b)
}

func BenchmarkFibNaiveRec3(b *testing.B) {
	benchmarkFindFibNaiveRec(10, b)
}

func BenchmarkFibNaiveRec4(b *testing.B) {
	benchmarkFindFibNaiveRec(20, b)
}

func BenchmarkFibNaiveRec5(b *testing.B) {
	benchmarkFindFibNaiveRec(30, b)
}

func BenchmarkFibNaiveRec6(b *testing.B) {
	benchmarkFindFibNaiveRec(50, b)
}

func Test_fibBottomUp(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sanity",
			args: args{30},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFibBottomUp(tt.args.n); got != tt.want {
				t.Errorf("fibBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFindFibBottomUp(n int, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findFibBottomUp(n)
	}
}

func BenchmarkFibBottomUp2(b *testing.B) {
	benchmarkFindFibBottomUp(5, b)
}

func BenchmarkFibBottomUp3(b *testing.B) {
	benchmarkFindFibBottomUp(10, b)
}

func BenchmarkFibBottomUp4(b *testing.B) {
	benchmarkFindFibBottomUp(20, b)
}

func BenchmarkFibBottomUp5(b *testing.B) {
	benchmarkFindFibBottomUp(30, b)
}

func BenchmarkFibBottomUp6(b *testing.B) {
	benchmarkFindFibBottomUp(50, b)
}

func BenchmarkFibBottomUp7(b *testing.B) {
	benchmarkFindFibBottomUp(100, b)
}

func BenchmarkFibBottomUp8(b *testing.B) {
	benchmarkFindFibBottomUp(1000, b)
}
