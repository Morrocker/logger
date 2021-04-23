package logger

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Bench(format string, a ...interface{}) {
	doLog(4, "magenta", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func BenchForce(format string, a ...interface{}) {
	doLog(1, "magenta", format, a...)
}
