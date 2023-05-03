package log

// Iterface for logging.
type Logger interface {
	// Infof may be used to write messages to the log Printf style
	Infof(format string, args ...interface{})
	// Info may be used to write messages to the log.
	Info(format string)
	// Warn may be used to write warning messages to the log Printf style
	Warn(format string)
	// Warnf may be used to write warning messages to the log Printf style
	Warnf(format string, args ...interface{})
	// Errorf may be used to write error messages to the log Printf style
	Errorf(format string, args ...interface{})
	// Error may be used to write error messages to the log.
	Error(format string)
	// Printf may be used to write messages to the log Printf style without timestamp
	Printf(format string, args ...interface{})
	// Print may be used to write messages to the log without timestamp
	Print(format string)
}
