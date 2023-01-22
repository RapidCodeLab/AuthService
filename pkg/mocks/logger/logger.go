package logger

type Logger struct{}

func (l *Logger) Infof(format string, args ...any) {}

func (l *Logger) Debugf(format string, args ...any) {}

func (l *Logger) Warnf(format string, args ...any) {}

func (l *Logger) Errorf(format string, args ...any) {}
