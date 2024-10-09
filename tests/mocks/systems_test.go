package mocks

import (
	"testing"
)

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
