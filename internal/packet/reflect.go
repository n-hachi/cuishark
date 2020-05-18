package packet

import (
	"fmt"
	"reflect"
	"strings"
)

func Oneline(i interface{}) (s string) {
	v := reflect.ValueOf(i)
	return oneline(v)
}

func oneline(v reflect.Value) (s string) {
	t := v.Type()

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	switch v.Kind() {
	case reflect.Struct:
		s = "{"
		for idx := 0; idx < t.NumField(); idx++ {
			t2 := t.Field(idx)
			if t2.Anonymous {
				continue
			}
			s += fmt.Sprintf("%s=%v, ", t.Field(idx).Name, oneline(v.Field(idx)))
		}
		s = strings.TrimRight(s, ", ")
		s += "}"
	case reflect.Array:
		s = "["
		for idx := 0; idx < t.Len(); idx++ {
			s += fmt.Sprintf("%d:%v, ", idx, oneline(v.Index(idx)))
		}
		s = strings.TrimRight(s, ", ")
		s += "]"
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}
