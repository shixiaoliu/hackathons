package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	// DEBUG level for detailed information
	DEBUG LogLevel = iota
	// INFO level for general operational information
	INFO
	// WARNING level for potential issues
	WARNING
	// ERROR level for error conditions
	ERROR
	// FATAL level for critical errors that require termination
	FATAL
)

var levelNames = map[LogLevel]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	FATAL:   "FATAL",
}

// Logger is a custom logger with support for different log levels
type Logger struct {
	level  LogLevel
	prefix string
	logger *log.Logger
}

// New creates a new Logger with the specified level and output
func New(level LogLevel, out io.Writer, prefix string) *Logger {
	return &Logger{
		level:  level,
		prefix: prefix,
		logger: log.New(out, "", log.LstdFlags),
	}
}

// NewDefaultLogger creates a new Logger with default settings
func NewDefaultLogger(level LogLevel) *Logger {
	return New(level, os.Stdout, "")
}

// NewFileLogger creates a new Logger that writes to a file
func NewFileLogger(level LogLevel, filename string, prefix string) (*Logger, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	// Open log file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	// Create logger that writes to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, file)
	return New(level, multiWriter, prefix), nil
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// GetLevel returns the current log level
func (l *Logger) GetLevel() LogLevel {
	return l.level
}

// log logs a message at the specified level
func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	// Get caller information
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	// Extract just the filename
	file = filepath.Base(file)

	// Format the message
	levelStr := levelNames[level]
	prefix := l.prefix
	if prefix != "" {
		prefix = "[" + prefix + "] "
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf(format, v...)
	logMessage := fmt.Sprintf("%s %s%s [%s:%d] %s",
		timestamp, prefix, levelStr, file, line, message)

	l.logger.Println(logMessage)

	// If this is a fatal message, exit the program
	if level == FATAL {
		os.Exit(1)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

// Info logs an informational message
func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

// Warning logs a warning message
func (l *Logger) Warning(format string, v ...interface{}) {
	l.log(WARNING, format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
}

// ParseLogLevel converts a string to a LogLevel
func ParseLogLevel(level string) (LogLevel, error) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING", "WARN":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return INFO, fmt.Errorf("invalid log level: %s", level)
	}
}

// Global logger instance
var defaultLogger = NewDefaultLogger(INFO)

// SetDefaultLogger sets the global default logger
func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

// Global logger functions
func Debug(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

func Warning(format string, v ...interface{}) {
	defaultLogger.Warning(format, v...)
}

func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	defaultLogger.Fatal(format, v...)
}
