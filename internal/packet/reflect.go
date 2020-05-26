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
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func Detail(i interface{}) (sl []string) {
	v := reflect.ValueOf(i)
	sl = detail(v, 1, 2)
	return sl
}

func detail(v reflect.Value, indentLevel int, indentWidth int) (sl []string) {
	indent := ""
	for i := 0; i < indentLevel*indentWidth; i++ {
		indent = indent + " "
	}

	t := v.Type()

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	switch v.Kind() {
	case reflect.Struct:
		align := 0
		for i := 0; i < t.NumField(); i++ {
			if align < len(t.Field(i).Name) {
				align = len(t.Field(i).Name)
			}
		}
		for i := 0; i < t.NumField(); i++ {
			t2 := t.Field(i)
			if t2.Anonymous {
				continue
			}

			// Do not print private variables.
			if r := rune(t2.Name[0]); !unicode.IsUpper(r) {
				continue
			}

			// Check child type is struct or not
			if t2.Type.Kind() == reflect.Struct {
				sl = append(sl, fmt.Sprintf("%s%-*v : ", indent, align, t2.Name))
				sl = append(sl, detail(v.Field(i), indentLevel+1, indentWidth)...)
			} else {
				sl = append(sl, fmt.Sprintf("%s%-*v : %v", indent, align, t2.Name, oneline(v.Field(i))))
			}
		}
	default:
		sl = append(sl, fmt.Sprintf("%s%v", indent, oneline(v)))
	}
	return sl
}
