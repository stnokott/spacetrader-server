package log

import (
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestComponentFormatterFormat(t *testing.T) {
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name    string
		f       componentFormatter
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "expected behaviour",
			f:    componentFormatter{MaxComponentLength: 3},
			args: args{
				&logrus.Entry{
					Level: logrus.DebugLevel,
					Time:  time.Date(2024, 8, 31, 16, 40, 38, 123, time.UTC),
					Data: logrus.Fields{
						_componentField: "Foo",
					},
					Message: "encountered a wild bar",
				},
			},
			want:    []byte("[DEBU] 2024-08-31 16:40:38 [Foo] encountered a wild bar\n"),
			wantErr: false,
		},
		{
			name: "padded component",
			f:    componentFormatter{MaxComponentLength: 6},
			args: args{
				&logrus.Entry{
					Level: logrus.DebugLevel,
					Time:  time.Date(2024, 8, 31, 16, 40, 38, 123, time.UTC),
					Data: logrus.Fields{
						_componentField: "Foo",
					},
					Message: "encountered a wild bar",
				},
			},
			want:    []byte("[DEBU] 2024-08-31 16:40:38 [Foo]    encountered a wild bar\n"),
			wantErr: false,
		},
		{
			name: "component too long",
			f:    componentFormatter{MaxComponentLength: 3},
			args: args{
				&logrus.Entry{
					Level: logrus.DebugLevel,
					Time:  time.Date(2024, 8, 31, 16, 40, 38, 123, time.UTC),
					Data: logrus.Fields{
						_componentField: "Foobarfuzz",
					},
					Message: "encountered a wild bar",
				},
			},
			want:    []byte("[DEBU] 2024-08-31 16:40:38 [Foobarfuzz] encountered a wild bar\n"),
			wantErr: false,
		},
		{
			name: "invalid component type",
			f:    componentFormatter{MaxComponentLength: 3},
			args: args{
				&logrus.Entry{
					Level: logrus.DebugLevel,
					Time:  time.Date(2024, 8, 31, 16, 40, 38, 123, time.UTC),
					Data: logrus.Fields{
						_componentField: 420,
					},
					Message: "encountered a wild bar",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing component name",
			f:    componentFormatter{MaxComponentLength: 3},
			args: args{
				&logrus.Entry{
					Level:   logrus.DebugLevel,
					Time:    time.Date(2024, 8, 31, 16, 40, 38, 123, time.UTC),
					Data:    logrus.Fields{},
					Message: "encountered a wild bar",
				},
			},
			want:    []byte("[DEBU] 2024-08-31 16:40:38 [!COMP] encountered a wild bar\n"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Format(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentFormatter.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComponentFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
