// Package log contains the custom logger for this application.
package log

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const environment = "debug" // TODO: configure from environment

const _componentField = "comp"

// ForComponent creates a new logger for the specified component.
//
// Any custom fields - using WithField(s) - are discarded.
func ForComponent(name string) *logrus.Entry {
	l := logrus.New()
	if environment == "prod" {
		l.SetLevel(logrus.InfoLevel)
	} else {
		l.SetLevel(logrus.DebugLevel)
	}
	l.SetFormatter(defaultFormatter)
	l.SetOutput(os.Stdout)

	e := l.WithField(_componentField, name)
	if len(name) > defaultFormatter.MaxComponentLength {
		defaultFormatter.MaxComponentLength = len(name)
	}
	return e
}

type componentFormatter struct {
	MaxComponentLength int
}

var _ logrus.Formatter = (*componentFormatter)(nil)

var defaultFormatter = &componentFormatter{}

var levelStrings = map[logrus.Level]string{
	logrus.PanicLevel: "PANI",
	logrus.FatalLevel: "FATA",
	logrus.ErrorLevel: "ERRO",
	logrus.WarnLevel:  "WARN",
	logrus.InfoLevel:  "INFO",
	logrus.DebugLevel: "DEBU",
	logrus.TraceLevel: "TRAC",
}

// Format implements logrus.Formatter
func (f *componentFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buf bytes.Buffer

	comp, exists := entry.Data[_componentField]
	if !exists {
		comp = "!COMP"
	}
	componentName, ok := comp.(string)
	if !ok {
		return nil, fmt.Errorf("field '%s' invalid type, expected string, got %T", _componentField, comp)
	}

	buf.WriteByte('[')
	buf.WriteString(levelStrings[entry.Level])
	buf.WriteString("] ")
	buf.WriteString(entry.Time.Format(time.DateTime))
	buf.WriteString(" [")
	buf.WriteString(componentName)
	buf.WriteByte(']')
	if len(componentName) < f.MaxComponentLength {
		buf.WriteString(strings.Repeat(" ", f.MaxComponentLength-len(componentName)))
	}
	buf.WriteByte(' ')
	buf.WriteString(entry.Message)
	buf.WriteByte('\n')

	return buf.Bytes(), nil
}

func padSpaces(s string, pad int) string {
	if len(s) >= pad {
		return s
	}
	return s + strings.Repeat(" ", pad-len(s))
}
