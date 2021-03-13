package logger

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

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
