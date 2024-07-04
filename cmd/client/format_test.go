package main

import (
	"testing"
	"time"
)

func TestFmtTimeUntil(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zero",
			args: args{
				from: time.Date(2024, 1, 1, 1, 1, 1, 1, time.UTC),
				to:   time.Date(2024, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			want: "00d 00h:00m:00s",
		},
		{
			name: "full",
			args: args{
				from: time.Date(2024, 1, 1, 1, 1, 1, 1, time.UTC),
				to:   time.Date(2024, 1, 12, 22, 57, 45, 334, time.UTC),
			},
			want: "11d 21h:56m:44s",
		},
		{
			name: "no days",
			args: args{
				from: time.Date(2024, 1, 1, 1, 1, 1, 1, time.UTC),
				to:   time.Date(2024, 1, 1, 22, 57, 45, 334, time.UTC),
			},
			want: "00d 21h:56m:44s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmtDuration(tt.args.to.Sub(tt.args.from)); got != tt.want {
				t.Errorf("fmtTimeUntil() = %v, want %v", got, tt.want)
			}
		})
	}
}
