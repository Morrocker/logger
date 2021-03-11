package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
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
)

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Notice(format string, a ...interface{}) {
	noticeLog(1, format, a...)
}

// NoticeV same as Notice(), but will only print when verbose or debug options are set
func NoticeV(format string, a ...interface{}) {
	noticeLog(2, format, a...)
}

// NoticeD same as Notice(), but will only print when the debug options is set
func NoticeD(format string, a ...interface{}) {
	noticeLog(3, format, a...)
}

// NoticeB same as Notice(), but will only print when the benchmark options is set
func NoticeB(format string, a ...interface{}) {
	noticeLog(4, format, a...)
}

func noticeLog(t int, f string, a ...interface{}) {
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
		printLog(os.Stdout, "blue", f, a...)
	}
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Alert(format string, a ...interface{}) {
	alertLog(1, format, a...)
}

// AlertV same as Alert(), but will only print when verbose or debug options are set
func AlertV(format string, a ...interface{}) {
	alertLog(2, format, a...)
}

// AlertD same as Alert(), but will only print when the debug options is set
func AlertD(format string, a ...interface{}) {
	alertLog(3, format, a...)
}

// AlertB same as Alert(), but will only print when the benchmark options is set
func AlertB(format string, a ...interface{}) {
	alertLog(4, format, a...)
}

func alertLog(t int, f string, a ...interface{}) {
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
		printLog(os.Stdout, "yellow", f, a...)
	}
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Info(format string, a ...interface{}) {
	infoLog(1, format, a...)
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func InfoV(format string, a ...interface{}) {
	infoLog(2, format, a...)
}

// InfoD same as Info(), but will only print when the debug options is set
func InfoD(format string, a ...interface{}) {
	infoLog(3, format, a...)
}

// InfoB same as Info(), but will only print when the banchnark options is set
func InfoB(format string, a ...interface{}) {
	infoLog(4, format, a...)
}

func infoLog(t int, f string, a ...interface{}) {
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
		printLog(os.Stdout, "cyan", f, a...)
	}
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Task(format string, a ...interface{}) {
	taskLog(1, format, a...)
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func TaskV(format string, a ...interface{}) {
	taskLog(2, format, a...)
}

// TaskD same as Task(), but will only print when the debug options is set
func TaskD(format string, a ...interface{}) {
	taskLog(3, format, a...)
}

// TaskB same as Task(), but will only print when the benchmark options is set
func TaskB(format string, a ...interface{}) {
	taskLog(4, format, a...)
}

func taskLog(t int, f string, a ...interface{}) {
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
		printLog(os.Stdout, "green", f, a...)
	}
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Error(format string, a ...interface{}) {
	errorLog(1, format, a...)
}

// ErrorV same as Error(), but will only print when verbose or debug options are set
func ErrorV(format string, a ...interface{}) {
	errorLog(2, format, a...)
}

// ErrorD same as Error(), but will only print when the debug options is set
func ErrorD(format string, a ...interface{}) {
	errorLog(3, format, a...)
}

// ErrorB same as Error(), but will only print when the benchmark options is set
func ErrorB(format string, a ...interface{}) {
	errorLog(4, format, a...)
}

func errorLog(t int, f string, a ...interface{}) {
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
		printLog(os.Stderr, "red", f, a...)
	}
}

// Obj prints out the given object using the spew library
func Obj(obj interface{}, objName string) {
	objLog(1, obj, objName)
}

// ObjV same as Obj(), but will only print when verbose or debug options are set
func ObjV(obj interface{}, objName string) {
	objLog(2, obj, objName)
}

// ObjD same as Obj(), but will only print when the debug options is set
func ObjD(obj interface{}, objName string) {
	objLog(3, obj, objName)
}

func objLog(t int, obj interface{}, objName string) {
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
	}
	if print && !silent {
		printLog(os.Stderr, "cyan", "Object: %s", objName)
		spew.Dump(obj)
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
