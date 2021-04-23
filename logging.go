package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type Scope struct {
	Info  int
	Task  int
	Error int
	Alert int
	Note  int
	Bench int
}

const (
	None = iota
	Regular
	Verbose
	Debug
)

var (
	outputFile  string
	writeScope  Scope = Scope{}
	silent            = false
	debug             = false
	verbose           = false
	benchmark         = false
	timestamp         = true
	preNote           = true
	dualMode          = false
	colorOn           = true
	tsFormat          = "2006-01-02 15:04:05"
	red               = color.New(color.FgHiRed).SprintFunc()
	errorPreMsg       = "[ERROR]"
	cyan              = color.New(color.FgHiCyan).SprintFunc()
	infoPreMsg        = "[INFO]"
	green             = color.New(color.FgHiGreen).SprintFunc()
	taskPreMsg        = "[TASK]"
	yellow            = color.New(color.FgHiYellow).SprintFunc()
	alertPreMsg       = "[ALERT]"
	blue              = color.New(color.FgHiBlue).SprintFunc()
	notePreMsg        = "[NOTE]"
	magenta           = color.New(color.FgHiMagenta).SprintFunc()
	benchPreMsg       = "[BENCH]"
)

// SetModes sets the verbose and debug variables according to given parameters
func SetModes(v, d, b bool) {
	verbose = v
	debug = d
	benchmark = b
}

// ToggleSilent enables/disables silent mode. No logs will be shown if enabled. Note that this does not prevent file logging.
func ToggleSilent() {
	silent = !silent
}

// ToggleTimestamp enables/disables timestamp on log
func ToggleTimestamp() {
	timestamp = !timestamp
}

// ToggleTimestamp enables/disables timestamp on log
func ToggleColor() {
	colorOn = !colorOn
}

// ToggleBenchmark enables/disables benchmark logging
func ToggleBenchmark() {
	benchmark = !benchmark
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

func printLog(color, format string, a ...interface{}) {
	var writter io.Writer = os.Stdout
	format = logFormat(format)
	if colorOn {
		switch color {
		case "blue":
			a = coalesce(blue(notePreMsg), a...)
		case "red":
			writter = os.Stderr
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
	} else {
		switch color {
		case "blue":
			a = coalesce(notePreMsg, a...)
		case "red":
			a = coalesce(errorPreMsg, a...)
		case "yellow":
			a = coalesce(alertPreMsg, a...)
		case "cyan":
			a = coalesce(infoPreMsg, a...)
		case "green":
			a = coalesce(taskPreMsg, a...)
		case "magenta":
			a = coalesce(benchPreMsg, a...)
		}

	}
	// fmt.Println(format)
	// fmt.Println(a)
	fmt.Fprintf(writter, format+"\n", a...)
}

func writeLog(color string, format string, a ...interface{}) {
	format = logFormat(format)
	switch color {
	case "blue":
		a = coalesce(notePreMsg, a...)
	case "red":
		a = coalesce(errorPreMsg, a...)
	case "yellow":
		a = coalesce(alertPreMsg, a...)
	case "cyan":
		a = coalesce(infoPreMsg, a...)
	case "green":
		a = coalesce(taskPreMsg, a...)
	case "magenta":
		a = coalesce(benchPreMsg, a...)
	}

	f, err := os.OpenFile(outputFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	txt := fmt.Sprintf(format+"\n", a...)
	if _, err := f.WriteString(txt); err != nil {
		log.Println(err)
	}
}

func doLog(n int, color string, format string, a ...interface{}) {
	switch {
	case (outputFile == "" && silent):
	}

	print := false
	switch n {
	case Regular:
		print = true
	case Verbose:
		if verbose || debug {
			print = true
		}
	case Debug:
		if debug {
			print = true
		}
	case 4:
		if benchmark {
			print = true
		}
	}
	if !print {
		return
	}

	if outputFile != "" && silent {
		writeLog(color, format, a...)
	} else if outputFile == "" && !silent {
		printLog(color, format, a...)
	} else {
		writeLog(color, format, a...)
		if dualMode {
			printLog(color, format, a...)
		}
	}
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

func shouldPrint() {

}

func getDate() string {
	d := time.Now().Format(tsFormat)
	return d
}

// OutputToFile sets the log to be printed to a file instead of StdOut. What gets written can be tuned.
func OutputToFile(filename string, scope Scope, dual bool) {
	outputFile = filename
	writeScope = scope
	dualMode = dual
}

// StopLogToFile turns off the OutputToFile option and resets these preferences
func StopToFile() {

}
