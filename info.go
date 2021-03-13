package logger

import "os"

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
