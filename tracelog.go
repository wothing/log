package log

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// TraceIn and TraceOut use in function in and out,reduce code line
// Example:
//	func test() {
//		user := User{Name: "zhangsan", Age: 21, School: "xayddx"}
//		service := "verification.GetVerifiCode"
//		defer log.TraceOut(log.TraceIn("12345", service, "user:%v", user))
//		....
//	}

func TraceIn(traceID string, service string, format string, v ...interface{}) (string, string, time.Time) {
	startTime := time.Now()
	Std.Output(traceID, Linfo, 2, fmt.Sprintf("calling "+service+", "+format, printStructs(v)...))
	return traceID, service, startTime
}

func TraceOut(traceID string, service string, startTime time.Time) {
	Std.Output(traceID, Linfo, 2, fmt.Sprintf("finished "+service+", took=%v", time.Since(startTime)))
}

func TraceCtx(ctx context.Context, service string, format string, v ...interface{}) (string, string, time.Time) {
	traceID := ""
	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			traceID = md["tid"][0]
		}
	}

	startTime := time.Now()
	Std.Output(traceID, Linfo, 2, fmt.Sprintf("calling "+service+", "+format, v...))
	return traceID, service, startTime
}

func printStructs(ss []interface{}) []interface{} {
	for i, _ := range ss {
		ss[i] = printStruct(ss[i], true)
	}
	return ss
}

func printStruct(s interface{}, names bool) string {
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

		if names {
			b.WriteString(t.Field(i).Name)
			b.WriteString(":")
		}

		v2 := v.Field(i)

		switch v2.Kind() {
		case reflect.Ptr, reflect.Interface:
			if v2.Elem().IsValid() {
				b.WriteString(printStruct(v2.Interface(), names))
			} else {
				b.WriteString("nil")
			}
		case reflect.Struct:
			b.WriteString(printStruct(v2.Interface(), names))
		default:
			b.WriteString(fmt.Sprint(v2))
		}
	}
	b.WriteString("}")
	return b.String()
}
