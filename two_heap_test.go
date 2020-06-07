package educativeingo

import (
	"testing"
)

func TestMedianOfStream_FindMedian(t *testing.T) {
	m := Constructor()
	m.InsertNum(3)
	m.InsertNum(1)

	m1 := Constructor()
	m1.InsertNum(3)
	m1.InsertNum(1)
	m1.InsertNum(5)

	m2 := Constructor()
	m2.InsertNum(3)
	m2.InsertNum(1)
	m2.InsertNum(5)
	m2.InsertNum(4)

	tests := []struct {
		name  string
		input *MedianOfStream
		want  float32
	}{
		{
			name:  "sanity",
			input: &m,
			want:  2.0,
		},
		{
			name:  "sanity",
			input: &m1,
			want:  3.0,
		},
		{
			name:  "sanity",
			input: &m2,
			want:  3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.FindMedian(); got != tt.want {
				t.Errorf("MedianOfStream.FindMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
