package log

import (
	"reflect"
)

func PkgPath(obj interface{}) string {
	ret := reflect.TypeOf(obj).PkgPath()
	if "" == ret {
		ret = "unknown"
	}
	return ret
}
