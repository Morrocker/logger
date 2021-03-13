package logger

import "os"

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
