package log

import (
	"fmt"
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
	Std.Output(traceID, Linfo, fmt.Sprintf("calling "+service+", "+format, v...))
	return traceID, service, startTime
}

func TraceOut(traceID string, service string, startTime time.Time) {
	Std.Output(traceID, Linfo, fmt.Sprintf("finished "+service+", took=%v", time.Since(startTime)))
}

func TraceCtx(ctx context.Context, service string, format string, v ...interface{}) (string, string, time.Time) {
	traceID := ""
	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			traceID = md["tid"][0]
		}
	}

	startTime := time.Now()
	Std.Output(traceID, Linfo, fmt.Sprintf("calling "+service+", "+format, v...))
	return traceID, service, startTime
}
