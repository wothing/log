package log

import (
	"fmt"
	"os"
	"runtime"
)

// Helper functions, nothing more but adding Trace ID to log functions.
// No more comments in this files, please check comments in logext.go

func (l *Logger) PrintfT(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Linfo, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) PrintT(traceID string, v ...interface{}) { 
	l.Output(traceID, Linfo, 2, fmt.Sprint(v...)) 
}

func (l *Logger) PrintlnT(traceID string, v ...interface{}) { 
	l.Output(traceID, Linfo, 2, fmt.Sprintln(v...)) 
}

func (l *Logger) DebugfT(traceID string, format string, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(traceID, Ldebug, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) DebugT(traceID string, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(traceID, Ldebug, 2, fmt.Sprintln(v...))
}

func (l *Logger) InfofT(traceID string, format string, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(traceID, Linfo, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) InfoT(traceID string, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(traceID, Linfo, 2, fmt.Sprintln(v...))
}

func (l *Logger) WarnfT(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Lwarn, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) WarnT(traceID string, v ...interface{}) { 
	l.Output(traceID, Lwarn, 2, fmt.Sprintln(v...)) 
}

func (l *Logger) ErrorfT(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Lerror, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) ErrorT(traceID string, v ...interface{}) { 
	l.Output(traceID, Lerror, 2, fmt.Sprintln(v...)) 
}

func (l *Logger) FatalT(traceID string, v ...interface{}) {
	l.Output(traceID, Lfatal, 2, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) FatalfT(traceID string, format string, v ...interface{}) {
	l.Output(traceID, Lfatal, 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) FatallnT(traceID string, v ...interface{}) {
	l.Output(traceID, Lfatal, 2, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) PanicT(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(traceID, Lpanic, 2, s)
	panic(s)
}

func (l *Logger) PanicfT(traceID string, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(traceID, Lpanic, 2, s)
	panic(s)
}

func (l *Logger) PaniclnT(traceID string, v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(traceID, Lpanic, 2, s)
	panic(s)
}

func (l *Logger) StackT(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	l.Output(traceID, Lerror, 2, s)
}

func PrintT(traceID string, v ...interface{}) {
	Std.Output(traceID, Linfo, 2, fmt.Sprint(v...))
}

func PrintfT(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Linfo, 2, fmt.Sprintf(format, v...))
}

func PrintlnT(traceID string, v ...interface{}) {
	Std.Output(traceID, Linfo, 2, fmt.Sprintln(v...))
}

func DebugfT(traceID string, format string, v ...interface{}) {
	if Ldebug < Std.Level {
		return
	}
	Std.Output(traceID, Ldebug, 2, fmt.Sprintf(format, v...))
}

func DebugT(traceID string, v ...interface{}) {
	if Ldebug < Std.Level {
		return
	}
	Std.Output(traceID, Ldebug, 2, fmt.Sprintln(v...))
}

func InfofT(traceID string, format string, v ...interface{}) {
	if Linfo < Std.Level {
		return
	}
	Std.Output(traceID, Linfo, 2, fmt.Sprintf(format, v...))
}

func InfoT(traceID string, v ...interface{}) {
	if Linfo < Std.Level {
		return
	}
	Std.Output(traceID, Linfo, 2, fmt.Sprintln(v...))
}

func WarnfT(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Lwarn, 2, fmt.Sprintf(format, v...))
}

func WarnT(traceID string, v ...interface{}) { 
	Std.Output(traceID, Lwarn, 2, fmt.Sprintln(v...)) 
}

func ErrorfT(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Lerror, 2, fmt.Sprintf(format, v...))
}

func ErrorT(traceID string, v ...interface{}) { 
	Std.Output(traceID, Lerror, 2, fmt.Sprintln(v...)) 
}

func FatalT(traceID string, v ...interface{}) {
	Std.Output(traceID, Lfatal, 2, fmt.Sprint(v...))
	os.Exit(1)
}

func FatalfT(traceID string, format string, v ...interface{}) {
	Std.Output(traceID, Lfatal, 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func FatallnT(traceID string, v ...interface{}) {
	Std.Output(traceID, Lfatal, 2, fmt.Sprintln(v...))
	os.Exit(1)
}

func PanicT(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	Std.Output(traceID, Lpanic, 2, s)
	panic(s)
}

func PanicfT(traceID string, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	Std.Output(traceID, Lpanic, 2, s)
	panic(s)
}

func PaniclnT(traceID string, v ...interface{}) {
	s := fmt.Sprintln(v...)
	Std.Output(traceID, Lpanic, 2, s)
	panic(s)
}

func StackT(traceID string, v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	Std.Output(traceID, Lerror, 2, s)
}
