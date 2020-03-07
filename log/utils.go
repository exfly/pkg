package log

import (
	"reflect"
)

func PkgPath(obj interface{}) string {
	ret := reflect.TypeOf(obj).PkgPath()
	switch ret {
	case "main":
		ret = "root"
	case "":
		ret = "unknown"
	}
	return ret
}
