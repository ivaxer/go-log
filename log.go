package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level int32

const (
	FATAL Level = iota - 4
	ERROR
	WARNING
	INFO
	DEBUG
	VDEBUG
	VVDEBUG
	VVVDEBUG
)

var levelNames = map[Level]string{
	FATAL:    "fatal",
	ERROR:    "error",
	WARNING:  "warning",
	INFO:     "info",
	DEBUG:    "debug0",
	VDEBUG:   "debug1",
	VVDEBUG:  "debug2",
	VVVDEBUG: "debug3",
}

type Logger struct {
	mu  sync.Mutex
	lvl Level
	w   io.Writer
}

func New(lvl Level) *Logger {
	return &Logger{lvl: lvl, w: os.Stderr}
}

func (l *Logger) print(lvl Level, args ...interface{}) {
	buf := &bytes.Buffer{}
	fmt.Fprintln(buf, args...)
	l.output(lvl, buf.Bytes())
}

func (l *Logger) printf(lvl Level, format string, args ...interface{}) {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, format, args...)

	if buf.Len() == 0 {
		return
	}

	if buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}
	l.output(lvl, buf.Bytes())
}

func (l *Logger) output(lvl Level, buf []byte) {
	if lvl > l.lvl {
		return
	}

	h := l.header(lvl)

	l.mu.Lock()
	defer l.mu.Unlock()

	l.w.Write(h)
	l.w.Write(buf)
}

func (l *Logger) header(lvl Level) []byte {
	// header format:
	//  YY/MM/DD HH:MM:SS.UUUUUU file:line] LEVEL:

	t := time.Now()
	year, month, day := t.Date()
	year %= 1000
	hour, minute, second := t.Clock()
	usecond := t.Nanosecond() / 1e3

	_, file, line, ok := runtime.Caller(4)
	if !ok {
		file, line = "???", 0
	} else {
		index := strings.LastIndex(file, "/")
		if index != -1 && index != len(file) {
			file = file[index+1:]
		}
	}

	level := levelNames[lvl]

	// TODO: don't use Sprintf because it's slow.
	format := "%.2d/%.2d/%.2d %.2d:%.2d:%.2d.%.6d %s:%d] %s: "
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, format, day, month, year, hour, minute, second,
		usecond, file, line, level)

	return buf.Bytes()
}

// VVVDebug logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) VVVDebug(args ...interface{}) {
	l.print(VVVDEBUG, args...)
}

// VVVDebugf logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) VVVDebugf(format string, args ...interface{}) {
	l.printf(VVVDEBUG, format, args...)
}

// VVDebug logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) VVDebug(args ...interface{}) {
	l.print(VVDEBUG, args...)
}

// VVDebugf logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) VVDebugf(format string, args ...interface{}) {
	l.printf(VVDEBUG, format, args...)
}

// VDebug logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) VDebug(args ...interface{}) {
	l.print(VDEBUG, args...)
}

// VDebugf logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) VDebugf(format string, args ...interface{}) {
	l.printf(VDEBUG, format, args...)
}

// Debug logs with DEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(args ...interface{}) {
	l.print(DEBUG, args...)
}

// Debugf logs with DEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.printf(DEBUG, format, args...)
}

// Info logs with INFO level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(args ...interface{}) {
	l.print(INFO, args...)
}

// Infof logs with INFO level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.printf(INFO, format, args...)
}

// Warning logs with WARNING level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warning(args ...interface{}) {
	l.print(WARNING, args...)
}

// Warningf logs with WARNING level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.printf(WARNING, format, args...)
}

// Error logs with ERROR level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(args ...interface{}) {
	l.print(ERROR, args...)
}

// Errorf logs with ERROR level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.printf(ERROR, format, args...)
}

// Fatal logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Fatal(args ...interface{}) {
	l.print(FATAL, args...)
	// TODO: print stack
	os.Exit(1)
}

// Fatalf logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.printf(FATAL, format, args...)
	// TODO: print stack
	os.Exit(1)
}

// SetLevel sets log level.
// Logs at or above this level go to log writer.
func (l *Logger) SetLevel(lvl Level) {
	l.lvl = lvl
}

var l = New(INFO)

// VVVDebug logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func VVVDebug(args ...interface{}) {
	l.print(VVVDEBUG, args...)
}

// VVVDebugf logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func VVVDebugf(format string, args ...interface{}) {
	l.printf(VVVDEBUG, format, args...)
}

// VVDebug logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func VVDebug(args ...interface{}) {
	l.print(VVDEBUG, args...)
}

// VVDebugf logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func VVDebugf(format string, args ...interface{}) {
	l.printf(VVDEBUG, format, args...)
}

// VDebug logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func VDebug(args ...interface{}) {
	l.print(VDEBUG, args...)
}

// VDebugf logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func VDebugf(format string, args ...interface{}) {
	l.printf(VDEBUG, format, args...)
}

// Debug logs with DEBUG level.
// Arguments are handled in the manner of fmt.Print.
func Debug(args ...interface{}) {
	l.print(DEBUG, args...)
}

// Debugf logs with DEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	l.printf(DEBUG, format, args...)
}

// Info logs with INFO level.
// Arguments are handled in the manner of fmt.Print.
func Info(args ...interface{}) {
	l.print(INFO, args...)
}

// Infof logs with INFO level.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	l.printf(INFO, format, args...)
}

// Warning logs with WARNING level.
// Arguments are handled in the manner of fmt.Print.
func Warning(args ...interface{}) {
	l.print(WARNING, args...)
}

// Warningf logs with WARNING level.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
	l.printf(WARNING, format, args...)
}

// Error logs with ERROR level.
// Arguments are handled in the manner of fmt.Print.
func Error(args ...interface{}) {
	l.print(ERROR, args...)
}

// Errorf logs with ERROR level.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	l.printf(ERROR, format, args...)
}

// Fatal logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(args ...interface{}) {
	l.print(FATAL, args...)
	// TODO: print stack
	os.Exit(1)
}

// Fatalf logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, args ...interface{}) {
	l.printf(FATAL, format, args...)
	// TODO: print stack
	os.Exit(1)
}

// SetLevel sets log level.
// Logs at or above this level go to log writer.
func SetLevel(lvl Level) {
	l.SetLevel(lvl)
}
