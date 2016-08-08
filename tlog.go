package log

import (
	"fmt"
	"os"
	"runtime"
)

// Helper functions, nothing more but adding Trace ID to log functions.
// No more comments in this files, please check comments in logext.go

func (l *Logger) Tprintf(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Linfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Tprint(traceID string, v ...interface{}) {
	l.Output(traceID, Linfo, fmt.Sprint(v...))
}

func (l *Logger) Tprintln(traceID string, v ...interface{}) {
	l.Output(traceID, Linfo, fmt.Sprintln(v...))
}

func (l *Logger) Tdebugf(traceID string, format string, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(traceID, Ldebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Tdebug(traceID string, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(traceID, Ldebug, fmt.Sprintln(v...))
}

func (l *Logger) Tinfof(traceID string, format string, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(traceID, Linfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Tinfo(traceID string, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(traceID, Linfo, fmt.Sprintln(v...))
}

func (l *Logger) Twarnf(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Lwarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Twarn(traceID string, v ...interface{}) {
	l.Output(traceID, Lwarn, fmt.Sprintln(v...))
}

func (l *Logger) Terrorf(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Lerror, fmt.Sprintf(format, v...))
}

func (l *Logger) Terror(traceID string, v ...interface{}) {
	l.Output(traceID, Lerror, fmt.Sprintln(v...))
}

func (l *Logger) Tfatal(traceID string, v ...interface{}) {
	l.Output(traceID, Lfatal, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Tfatalf(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Lfatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Tfatalln(traceID string, v ...interface{}) {
	l.Output(traceID, Lfatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Tpanic(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(traceID, Lpanic, s)
	panic(s)
}

func (l *Logger) Tpanicf(traceID string, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(traceID, Lpanic, s)
	panic(s)
}

func (l *Logger) Tpanicln(traceID string, v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(traceID, Lpanic, s)
	panic(s)
}

func (l *Logger) Tstack(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	l.Output(traceID, Lerror, s)
}

func Tprint(traceID string, v ...interface{}) {
	Std.Output(traceID, Linfo, fmt.Sprint(v...))
}

func Tprintf(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Linfo, fmt.Sprintf(format, v...))
}

func Tprintln(traceID string, v ...interface{}) {
	Std.Output(traceID, Linfo, fmt.Sprintln(v...))
}

func Tdebugf(traceID string, format string, v ...interface{}) {
	if Ldebug < Std.Level {
		return
	}
	Std.Output(traceID, Ldebug, fmt.Sprintf(format, v...))
}

func Tdebug(traceID string, v ...interface{}) {
	if Ldebug < Std.Level {
		return
	}
	Std.Output(traceID, Ldebug, fmt.Sprintln(v...))
}

func Tinfof(traceID string, format string, v ...interface{}) {
	if Linfo < Std.Level {
		return
	}
	Std.Output(traceID, Linfo, fmt.Sprintf(format, v...))
}

func Tinfo(traceID string, v ...interface{}) {
	if Linfo < Std.Level {
		return
	}
	Std.Output(traceID, Linfo, fmt.Sprintln(v...))
}

func Twarnf(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Lwarn, fmt.Sprintf(format, v...))
}

func Twarn(traceID string, v ...interface{}) {
	Std.Output(traceID, Lwarn, fmt.Sprintln(v...))
}

func Terrorf(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Lerror, fmt.Sprintf(format, v...))
}

func Terror(traceID string, v ...interface{}) {
	Std.Output(traceID, Lerror, fmt.Sprintln(v...))
}

func Tfatal(traceID string, v ...interface{}) {
	Std.Output(traceID, Lfatal, fmt.Sprint(v...))
	os.Exit(1)
}

func Tfatalf(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Lfatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Tfatalln(traceID string, v ...interface{}) {
	Std.Output(traceID, Lfatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func Tpanic(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	Std.Output(traceID, Lpanic, s)
	panic(s)
}

func Tpanicf(traceID string, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	Std.Output(traceID, Lpanic, s)
	panic(s)
}

func Tpanicln(traceID string, v ...interface{}) {
	s := fmt.Sprintln(v...)
	Std.Output(traceID, Lpanic, s)
	panic(s)
}

func Tstack(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	Std.Output(traceID, Lerror, s)
}
