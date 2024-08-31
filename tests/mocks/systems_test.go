package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	type args struct {
		src        []int
		chunkIndex int
		chunkSize  int
	}
	tests := []struct {
		name      string
		args      args
		want      []int
		wantPanic bool
	}{
		{
			name: "exact fit",
			args: args{
				src:        []int{1, 2, 3, 4, 5},
				chunkIndex: 0,
				chunkSize:  5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "middle chunk",
			args: args{
				src:        []int{1, 2, 3, 4, 5, 6},
				chunkIndex: 1,
				chunkSize:  2,
			},
			want: []int{3, 4},
		},
		{
			name: "end chunk",
			args: args{
				src:        []int{1, 2, 3, 4, 5, 6},
				chunkIndex: 2,
				chunkSize:  2,
			},
			want: []int{5, 6},
		},
		{
			name: "end chunk cropped",
			args: args{
				src:        []int{1, 2, 3, 4, 5},
				chunkIndex: 2,
				chunkSize:  2,
			},
			want: []int{5},
		},
		{
			name: "chunk out of bounds exact",
			args: args{
				src:        []int{1, 2, 3, 4},
				chunkIndex: 2,
				chunkSize:  2,
			},
			wantPanic: true,
		},
		{
			name: "chunk out of bounds additional",
			args: args{
				src:        []int{1, 2, 3, 4, 5},
				chunkIndex: 3,
				chunkSize:  2,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					Chunk(tt.args.src, tt.args.chunkIndex, tt.args.chunkSize)
				})
				return
			}
			got := Chunk(tt.args.src, tt.args.chunkIndex, tt.args.chunkSize)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGenerateSystemName(t *testing.T) {
	type args struct {
		total int
		i     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single",
			args: args{
				total: 10,
				i:     5,
			},
			want: "SYSTEM-05",
		},
		{
			name: "double start",
			args: args{
				total: 50,
				i:     0,
			},
			want: "SYSTEM-00",
		},
		{
			name: "double middle",
			args: args{
				total: 50,
				i:     1,
			},
			want: "SYSTEM-01",
		},
		{
			name: "triple",
			args: args{
				total: 100,
				i:     29,
			},
			want: "SYSTEM-029",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSystemName(tt.args.total, tt.args.i); got != tt.want {
				t.Errorf("GenerateSystemName() = %v, want %v", got, tt.want)
			}
		})
	}
}
