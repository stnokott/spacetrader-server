// Package log contains the custom logger for this application.
package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

const environment = "debug" // TODO: configure from environment

const _componentField = "comp"

// ForComponent creates a new logger for the specified component.
//
// Any custom fields - using WithField(s) - are discarded.
func ForComponent(name string) *logrus.Entry {
	l := logrus.New()
	// set level
	if environment == "prod" {
		l.SetLevel(logrus.InfoLevel)
	} else {
		l.SetLevel(logrus.DebugLevel)
	}
	// use custom formatter
	l.SetFormatter(defaultFormatter)
	l.SetOutput(colorable.NewColorableStdout())

	// set component name via field
	e := l.WithField(_componentField, name)
	// set formatter's max component length
	if len(name) > defaultFormatter.MaxComponentLength {
		defaultFormatter.MaxComponentLength = len(name)
	}
	addComponentColor(name)
	return e
}

type componentFormatter struct {
	MaxComponentLength int
}

var _ logrus.Formatter = (*componentFormatter)(nil)

var defaultFormatter = &componentFormatter{}

var levelStrings = map[logrus.Level]string{
	logrus.PanicLevel: "PAN",
	logrus.FatalLevel: "FTL",
	logrus.ErrorLevel: "ERR",
	logrus.WarnLevel:  "WRN",
	logrus.InfoLevel:  "INF",
	logrus.DebugLevel: "DBG",
	logrus.TraceLevel: "TRC",
}
var levelColors = map[logrus.Level]aurora.Color{
	logrus.PanicLevel: aurora.RedBg | aurora.WhiteFg,
	logrus.FatalLevel: aurora.RedBg | aurora.WhiteFg,
	logrus.ErrorLevel: aurora.RedBg | aurora.WhiteFg,
	logrus.WarnLevel:  aurora.MagentaBg | aurora.BrightBg | aurora.WhiteFg,
	logrus.InfoLevel:  aurora.BlueBg | aurora.WhiteFg,
	logrus.DebugLevel: aurora.WhiteFg,
	logrus.TraceLevel: aurora.WhiteFg,
}

// maps each component name to a unique color which is passed to aurora.Index().
// this is used to give each component a unique color to better distinguish them in the log.
var componentColors = map[string]uint8{}

// we predefine a biased (ordered) list of colors, beginning with ones that are very clear.
// this prevents using faint colors which can happen when randomly choosing available colors.
var componentColorPool = []uint8{
	31,  // turqoise
	41,  // light yellow
	55,  // dark blue
	89,  // pink
	84,  // dark green
	197, // light red
	202, // orange
	204, // light pink
	225, // light blue
}

func addComponentColor(name string) {
	// noop if already exists
	if _, exists := componentColors[name]; exists {
		return
	}

	if len(componentColors) >= len(componentColorPool) {
		panic(fmt.Sprintf("%d logger components exceeded", len(componentColorPool)))
	}

	componentColors[name] = componentColorPool[len(componentColors)]
}

// Format implements logrus.Formatter
func (f *componentFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	comp, exists := entry.Data[_componentField]
	if !exists {
		comp = "!COMP"
	}
	componentName, ok := comp.(string)
	if !ok {
		return nil, fmt.Errorf("field '%s' invalid type, expected string, got %T", _componentField, comp)
	}

	levelStr := "[" + levelStrings[entry.Level] + "]"
	timeStr := entry.Time.Format(time.DateTime)
	compStr := "[" + componentName + "]"
	// space padding
	if len(componentName) < f.MaxComponentLength {
		compStr += strings.Repeat(" ", f.MaxComponentLength-len(componentName))
	}

	var out string
	if entry.Level == logrus.DebugLevel || entry.Level == logrus.TraceLevel {
		// faint style for the whole message
		out = aurora.Faint(fmt.Sprintln(levelStr, timeStr, compStr, entry.Message)).String()
	} else {
		// regular, colored style
		out = fmt.Sprintln(
			aurora.Colorize(levelStr, levelColors[entry.Level]),
			timeStr,
			aurora.Index(componentColors[componentName], compStr),
			entry.Message,
		)
	}

	return []byte(out), nil
}

func padSpaces(s string, pad int) string {
	if len(s) >= pad {
		return s
	}
	return s + strings.Repeat(" ", pad-len(s))
}
