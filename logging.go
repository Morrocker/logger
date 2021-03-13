package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	silent      = false
	debug       = false
	verbose     = false
	benchmark   = false
	timestamp   = true
	preNote     = true
	tsFormat    = "2006-01-02 15:04:05"
	red         = color.New(color.FgHiRed).SprintFunc()
	errorPreMsg = "[ERROR]"
	cyan        = color.New(color.FgHiCyan).SprintFunc()
	infoPreMsg  = "[INFO]"
	green       = color.New(color.FgHiGreen).SprintFunc()
	taskPreMsg  = "[TASK]"
	yellow      = color.New(color.FgHiYellow).SprintFunc()
	alertPreMsg = "[ALERT]"
	blue        = color.New(color.FgHiBlue).SprintFunc()
	notePreMsg  = "[NOTE]"
	magenta     = color.New(color.FgHiMagenta).SprintFunc()
	benchPreMsg = "[BENCH]"
)

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Bench(format string, a ...interface{}) {
	benchmarkLog(1, format, a...)
}
func benchmarkLog(t int, f string, a ...interface{}) {
	print := false
	switch t {
	case 1:
		print = true
	case 2:
		if verbose || debug {
			print = true
		}
	case 3:
		if debug {
			print = true
		}
	case 4:
		if benchmark {
			print = true
		}
	}

	if print && !silent {
		printLog(os.Stdout, "magenta", f, a...)
	}
}

// SetModes sets the verbose and debug variables according to given parameters
func SetModes(v, d, b bool) {
	verbose = v
	debug = d
	benchmark = b
}

// ToggleSilent enables/disables silent mode. No logs will be shown if enabled.
func ToggleSilent() {
	silent = !silent
}

// ToggleTimestamp enables/disables timestamp on log
func ToggleTimestamp() {
	timestamp = !timestamp
}

// TogglePreNote enables/disables timestamp on log
func TogglePreNote() {
	preNote = !preNote
}

func logFormat(format string) string {
	var ret string
	if timestamp && preNote {
		ret = "%s\t%s: "
	} else if timestamp {
		ret = "%s: "
	} else if preNote {
		ret = "%s\t"
	}
	return ret + format
}

func printLog(writer io.Writer, color, format string, a ...interface{}) {
	format = logFormat(format)
	switch color {
	case "blue":
		a = coalesce(blue(notePreMsg), a...)
	case "red":
		a = coalesce(red(errorPreMsg), a...)
	case "yellow":
		a = coalesce(yellow(alertPreMsg), a...)
	case "cyan":
		a = coalesce(cyan(infoPreMsg), a...)
	case "green":
		a = coalesce(green(taskPreMsg), a...)
	case "magenta":
		a = coalesce(magenta(benchPreMsg), a...)
	}
	fmt.Fprintf(writer, format+"\n", a...)
}

func coalesce(header string, a ...interface{}) []interface{} {
	d := getDate()
	var ret []interface{}
	if timestamp && preNote {
		ret = []interface{}{header, d}
	} else if timestamp {
		ret = []interface{}{d}
	} else if preNote {
		ret = []interface{}{header}
	}

	ret = append(ret, a...)
	return ret
}

func getDate() string {
	d := time.Now().Format(tsFormat)
	return d
}
