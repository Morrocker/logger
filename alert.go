package logger

import "os"

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
