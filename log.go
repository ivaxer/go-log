package log

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

type Logger struct {
	lvl Level
}

func New(lvl Level) *Logger {
	return &Logger{lvl}
}

// VVVDebug logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) VVVDebug(args ...interface{}) {}

// VVVDebugf logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) VVVDebugf(format string, args ...interface{}) {}

// VVDebug logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) VVDebug(args ...interface{}) {}

// VVDebugf logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) VVDebugf(format string, args ...interface{}) {}

// VDebug logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) VDebug(args ...interface{}) {}

// VDebugf logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) VDebugf(foramt string, args ...interface{}) {}

// Debug logs with DEBUG level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(args ...interface{}) {}

// Debugf logs with DEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, args ...interface{}) {}

// Info logs with INFO level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(args ...interface{}) {}

// Infof logs with INFO level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {}

// Warning logs with WARNING level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warning(args ...interface{}) {}

// Warningf logs with WARNING level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, args ...interface{}) {}

// Error logs with ERROR level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(args ...interface{}) {}

// Errorf logs with ERROR level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(foramt string, args ...interface{}) {}

// Fatal logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Fatal(args ...interface{}) {}

// Fatalf logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Fatalf(format string, args ...interface{}) {}

// SetLevel sets log level.
// Logs at or above this level go to log writer.
func (l *Logger) SetLevel(lvl Level) {
	l.lvl = lvl
}

// VVVDebug logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func VVVDebug(args ...interface{}) {}

// VVVDebugf logs with VVVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func VVVDebugf(format string, args ...interface{}) {}

// VVDebug logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func VVDebug(args ...interface{}) {}

// VVDebugf logs with VVDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func VVDebugf(format string, args ...interface{}) {}

// VDebug logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Print.
func VDebug(args ...interface{}) {}

// VDebugf logs with VDEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func VDebugf(format string, args ...interface{}) {}

// Debug logs with DEBUG level.
// Arguments are handled in the manner of fmt.Print.
func Debug(args ...interface{}) {}

// Debugf logs with DEBUG level.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {}

// Info logs with INFO level.
// Arguments are handled in the manner of fmt.Print.
func Info(args ...interface{}) {}

// Infof logs with INFO level.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {}

// Warning logs with WARNING level.
// Arguments are handled in the manner of fmt.Print.
func Warning(args ...interface{}) {}

// Warningf logs with WARNING level.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {}

// Error logs with ERROR level.
// Arguments are handled in the manner of fmt.Print.
func Error(args ...interface{}) {}

// Errorf logs with ERROR level.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(foramt string, args ...interface{}) {}

// Fatal logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(args ...interface{}) {}

// Fatalf logs with FATAL level, including a stack trace
// of all goroutins, than calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, args ...interface{}) {}

// SetLevel sets log level.
// Logs at or above this level go to log writer.
func SetLevel(lvl Level) {
	l.SetLevel(lvl)
}
