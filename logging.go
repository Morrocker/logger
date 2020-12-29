package logger

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

var debug, verbose bool

//var yellow := color.New(color.FgHiYellow).SprintFunc()
var (
	red    = color.New(color.FgHiRed).SprintFunc()
	cyan   = color.New(color.FgHiCyan).SprintFunc()
	green  = color.New(color.FgHiGreen).SprintFunc()
	yellow = color.New(color.FgHiYellow).SprintFunc()
	blue   = color.New(color.FgHiBlue).SprintFunc()
)

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Notice(f string, a ...interface{}) {
	a = coalesce(blue("[NOTE]"), a...)
	fmt.Printf("%s %s: "+f+"\n", a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Alert(f string, a ...interface{}) {
	alert(0, f, a...)
}

// AlertV works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end. Only outputs if verbose is enabled.
func AlertV(f string, a ...interface{}) {
	alert(1, f, a...)
}

// AlertD works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end. Only outputs if verbose or debug is enabled.
func AlertD(f string, a ...interface{}) {
	alert(2, f, a...)
}

// LogAlert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func alert(t int, f string, a ...interface{}) {
	switch t {
	case 1:
		if verbose || debug {
			a = coalesce(yellow("[ALERT]"), a...)
			fmt.Printf("%s %s: "+f+"\n", a...)
		}
	case 2:
		if debug {
			a = coalesce(yellow("[ALERT]"), a...)
			fmt.Printf("%s %s: "+f+"\n", a...)
		}
	default:
		a = coalesce(yellow("[ALERT]"), a...)
		fmt.Printf("%s %s: "+f+"\n", a...)
	}
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Info(f string, a ...interface{}) {
	info(0, f, a...)
}

// InfoV works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end. Only outputs if verbose is enabled.
func InfoV(f string, a ...interface{}) {
	info(1, f, a...)
}

// InfoD works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end. Only outputs if verbose or debug is enabled.
func InfoD(f string, a ...interface{}) {
	info(2, f, a...)
}

// LogInfo works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func info(t int, f string, a ...interface{}) {
	switch t {
	case 1:
		if verbose || debug {
			a = coalesce(cyan("[INFO] "), a...)
			fmt.Printf("%s %s: "+f+"\n", a...)
		}
	case 2:
		if debug {
			a = coalesce(cyan("[INFO] "), a...)
			fmt.Printf("%s %s: "+f+"\n", a...)
		}
	default:
		a = coalesce(cyan("[INFO] "), a...)
		fmt.Printf("%s %s: "+f+"\n", a...)
	}
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Task(f string, a ...interface{}) {
	task(0, f, a...)
}

// TaskV works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end. Only outputs if verbose is enabled.
func TaskV(f string, a ...interface{}) {
	task(1, f, a...)
}

// TaskD works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end. Only outputs if verbose or debug is enabled.
func TaskD(f string, a ...interface{}) {
	task(2, f, a...)
}

// LogTask works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func task(t int, f string, a ...interface{}) {
	switch t {
	case 1:
		if verbose || debug {
			a = coalesce(green("[TASK] "), a...)
			fmt.Printf("%s %s: "+f+"\n", a...)
			return
		}
		fallthrough
	case 2:
		if debug {
			a = coalesce(green("[TASK] "), a...)
			fmt.Printf("%s %s: "+f+"\n", a...)
		}
	default:
		a = coalesce(green("[TASK] "), a...)
		fmt.Printf("%s %s: "+f+"\n", a...)
	}
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Error(f string, a ...interface{}) {
	a = coalesce(red("[ERROR]"), a...)
	fmt.Printf("%s %s: "+f+"\n", a...)
}

// Obj prints out the given object using the spew library
func Obj(obj interface{}) {
	if debug {
		d := getDate()
		fmt.Printf("%s %s.\n", green("[OBJECT START]"), d)
		spew.Dump(obj)
		fmt.Printf("%s %s.\n", green("[OBJECT END]"), d)
	}

}

func coalesce(header string, a ...interface{}) []interface{} {
	d := getDate()
	ret := []interface{}{header, d}
	ret = append(ret, a...)
	return ret
}

func getDate() string {
	d := time.Now().Format("2006-01-02 15:04:05")
	return d
}

// SetModes sets the verbose and debug variables according to given parameters
func SetModes(v, d bool) {
	verbose = v
	debug = d
}
