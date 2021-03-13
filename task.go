package logger

import "os"

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
