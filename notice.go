package logger

import "os"

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
