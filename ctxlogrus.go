package ctxlogrus

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var getDefaultLogger = func() logrus.Ext1FieldLogger { return logrus.StandardLogger() }

// DisableDefaultLogger disables the default logger.
// After calling this function, calls to FromContext will return nil
// if no logger is set in the context.
func DisableDefaultLogger() {
	getDefaultLogger = func() logrus.Ext1FieldLogger {
		return nil
	}
}

// EnableDefaultLogger enables the default logger.
// After calling this function, calls to FromContext will return the default logger
// if no logger is set in the context.
func EnableDefaultLogger() {
	getDefaultLogger = func() logrus.Ext1FieldLogger { return logrus.StandardLogger() }
}

type LoggerKeyType string

const LoggerKey = LoggerKeyType("logger")

// NewContext creates a new context with the provided FieldLoggerlogger.
// If the context is nil, it creates a new background context.
// The logger is stored in the context under the key "logger".
func NewContext(ctx context.Context, logger logrus.Ext1FieldLogger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, LoggerKey, logger)
}

// FromContext retrieves the logger from the context.
// If no logger is found, it returns the default logger if enabled,
// or nil if the default logger is disabled.
func FromContext(ctx context.Context) logrus.Ext1FieldLogger {
	logger, ok := ctx.Value(LoggerKey).(logrus.Ext1FieldLogger)
	if !ok {
		return getDefaultLogger()
	}
	return logger
}

// WithError creates an entry from the logger in context and adds an error to it, using the value defined in ErrorKey as key.
func WithError(ctx context.Context, err error) logrus.Ext1FieldLogger {
	l := FromContext(ctx)
	if l == nil {
		return nil
	}
	return l.WithError(err)
}

// WithField creates an entry from the logger in context and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(ctx context.Context, key string, value interface{}) logrus.Ext1FieldLogger {
	l := FromContext(ctx)
	if l == nil {
		return nil
	}
	return l.WithField(key, value)
}

// WithFields creates an entry from the logger in context and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(ctx context.Context, fields logrus.Fields) logrus.Ext1FieldLogger {
	l := FromContext(ctx)
	if l == nil {
		return nil
	}
	return l.WithFields(fields)
}

// Trace logs a message at level Trace on the logger from the context.
func Trace(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Trace(args...)
}

// Debug logs a message at level Debug on the logger from the context.
func Debug(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Debug(args...)
}

// Print logs a message at level Info on the logger from the context.
func Print(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Print(args...)
}

// Info logs a message at level Info on the logger from the context.
func Info(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Info(args...)
}

// Warn logs a message at level Warn on the logger from the context.
func Warn(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warn(args...)
}

// Warning logs a message at level Warn on the logger from the context.
func Warning(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warning(args...)
}

// Error logs a message at level Error on the logger from the context.
func Error(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Error(args...)
}

// Panic logs a message at level Panic on the logger from the context.
func Panic(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Panic(args...)
}

// Fatal logs a message at level Fatal on the logger from the context then
// the process will exit with status set to 1.
func Fatal(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		os.Exit(1)
	}
	l.Fatal(args...)
}

// Debugf logs a message at level Debug on the logger from the context.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Debugf(format, args...)
}

// Printf logs a message at level Info on the logger from the context.
func Printf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Printf(format, args...)
}

// Infof logs a message at level Info on the logger from the context.
func Infof(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Infof(format, args...)
}

// Warnf logs a message at level Warn on the logger from the context.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the logger from the context.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warningf(format, args...)
}

// Errorf logs a message at level Error on the logger from the context.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the logger from the context.
func Panicf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the logger from the context then
// the process will exit with status set to 1.
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		os.Exit(1)
	}
	l.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the logger from the context.
func Traceln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Traceln(args...)
}

// Debugln logs a message at level Debug on the logger from the context.
func Debugln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Debugln(args...)
}

// Println logs a message at level Info on the logger from the context.
func Println(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Println(args...)
}

// Infoln logs a message at level Info on the logger from the context.
func Infoln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Infoln(args...)
}

// Warnln logs a message at level Warn on the logger from the context.
func Warnln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warnln(args...)
}

// Warningln logs a message at level Warn on the logger from the context.
func Warningln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warningln(args...)
}

// Errorln logs a message at level Error on the logger from the context.
func Errorln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Errorln(args...)
}

// Panicln logs a message at level Panic on the logger from the context.
func Panicln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the logger from the context then
// the process will exit with status set to 1.
func Fatalln(ctx context.Context, args ...interface{}) {
	l := FromContext(ctx)
	if l == nil {
		os.Exit(1)
	}
	l.Fatalln(args...)
}
