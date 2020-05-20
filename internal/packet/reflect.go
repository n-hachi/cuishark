package packet

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// Oneline returns json like string which represent input variable.
func Oneline(i interface{}) (s string) {
	v := reflect.ValueOf(i)
	s = oneline(v)
	return s
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
		for i := 0; i < t.NumField(); i++ {
			t2 := t.Field(i)
			if t2.Anonymous {
				continue
			}

			// Do not print private variables.
			if r := rune(t2.Name[0]); !unicode.IsUpper(r) {
				continue
			}
			s += fmt.Sprintf("%s:%v, ", t.Field(i).Name, oneline(v.Field(i)))
		}
		s = strings.TrimRight(s, ", ")
		s += "}"
	case reflect.Array, reflect.Slice:
		s = "["
		for i := 0; i < v.Len(); i++ {
			s += fmt.Sprintf("%v, ", oneline(v.Index(i)))
		}
		s = strings.TrimRight(s, ", ")
		s += "]"
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}
