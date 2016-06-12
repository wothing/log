package log

import (
	"bytes"
	"fmt"
	"reflect"
)

// PrintStruct tries walk struct return formatted string
func PrintStruct(s interface{}) string {
	b := &bytes.Buffer{}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		b.WriteString("&")
		v = reflect.Indirect(v)
	}

	t := v.Type()

	if t.Kind() != reflect.Struct {
		return fmt.Sprint(s)
	}

	b.WriteString(v.Type().String())

	b.WriteString("{")
	for i, l := 0, v.NumField(); i < l; i++ {
		if i > 0 {
			b.WriteString(" ")
		}

		// field name
		b.WriteString(t.Field(i).Name)
		b.WriteString(":")

		v2 := v.Field(i)

		switch v2.Kind() {
		case reflect.Ptr, reflect.Interface:
			if v2.Elem().IsValid() {
				b.WriteString(PrintStruct(v2.Interface()))
			} else {
				b.WriteString("nil")
			}
		case reflect.Struct:
			b.WriteString(PrintStruct(v2.Interface()))
		default:
			b.WriteString(fmt.Sprint(v2))
		}
	}
	b.WriteString("}")
	return b.String()
}
