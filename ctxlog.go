package log

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func getTracerIDFromCtx(ctx context.Context) string {
	nid := "00000000-0000-0000-0000-000000000000"

	if ctx == nil {
		return nid
	}

	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			return md["tid"][0]
		}
	}
	return nid
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(getTracerIDFromCtx(ctx), Ldebug, fmt.Sprintf(format, v...))
}

func (l *Logger) CtxDebug(ctx context.Context, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(getTracerIDFromCtx(ctx), Ldebug, fmt.Sprintln(v...))
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprintf(format, v...))
}

func (l *Logger) CtxInfo(ctx context.Context, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprintln(v...))
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lwarn, fmt.Sprintf(format, v...))
}

func (l *Logger) CtxWarn(ctx context.Context, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lwarn, fmt.Sprintln(v...))
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lerror, fmt.Sprintf(format, v...))
}

func (l *Logger) CtxError(ctx context.Context, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lerror, fmt.Sprintln(v...))
}

func (l *Logger) CtxFatal(ctx context.Context, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lfatal, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lfatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) CtxFatalln(ctx context.Context, v ...interface{}) {
	l.Output(getTracerIDFromCtx(ctx), Lfatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) CtxPanic(ctx context.Context, v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(getTracerIDFromCtx(ctx), Lpanic, s)
	panic(s)
}

func (l *Logger) CtxPanicf(ctx context.Context, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(getTracerIDFromCtx(ctx), Lpanic, s)
	panic(s)
}

func (l *Logger) CtxPanicln(ctx context.Context, v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(getTracerIDFromCtx(ctx), Lpanic, s)
	panic(s)
}

func (l *Logger) CtxStack(ctx context.Context, v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	l.Output(getTracerIDFromCtx(ctx), Lerror, s)
}

func CtxPrint(ctx context.Context, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprint(v...))
}

func CtxPrintf(ctx context.Context, format string, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprintf(format, v...))
}

func CtxPrintln(ctx context.Context, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprintln(v...))
}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	if Ldebug < Std.Level {
		return
	}
	Std.Output(getTracerIDFromCtx(ctx), Ldebug, fmt.Sprintf(format, v...))
}

func CtxDebug(ctx context.Context, v ...interface{}) {
	if Ldebug < Std.Level {
		return
	}
	Std.Output(getTracerIDFromCtx(ctx), Ldebug, fmt.Sprintln(v...))
}

func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	if Linfo < Std.Level {
		return
	}
	Std.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprintf(format, v...))
}

func CtxInfo(ctx context.Context, v ...interface{}) {
	if Linfo < Std.Level {
		return
	}
	Std.Output(getTracerIDFromCtx(ctx), Linfo, fmt.Sprintln(v...))
}

func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lwarn, fmt.Sprintf(format, v...))
}

func CtxWarn(ctx context.Context, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lwarn, fmt.Sprintln(v...))
}

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lerror, fmt.Sprintf(format, v...))
}

func CtxError(ctx context.Context, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lerror, fmt.Sprintln(v...))
}

func CtxFatal(ctx context.Context, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lfatal, fmt.Sprint(v...))
	os.Exit(1)
}

func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lfatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func CtxFatalln(ctx context.Context, v ...interface{}) {
	Std.Output(getTracerIDFromCtx(ctx), Lfatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func CtxPanic(ctx context.Context, v ...interface{}) {
	s := fmt.Sprint(v...)
	Std.Output(getTracerIDFromCtx(ctx), Lpanic, s)
	panic(s)
}

func CtxPanicf(ctx context.Context, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	Std.Output(getTracerIDFromCtx(ctx), Lpanic, s)
	panic(s)
}

func CtxPanicln(ctx context.Context, v ...interface{}) {
	s := fmt.Sprintln(v...)
	Std.Output(getTracerIDFromCtx(ctx), Lpanic, s)
	panic(s)
}

func CtxStack(ctx context.Context, v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	Std.Output(getTracerIDFromCtx(ctx), Lerror, s)
}
