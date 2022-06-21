package logging_zap

import (
	"go.uber.org/zap"
)

var _ LoggerImpl = (*Logger)(nil)

func Wrap(l *zap.Logger) LoggerImpl {
	l = l.WithOptions(zap.AddCallerSkip(1))
	return &Logger{l: l.Sugar()}
}

func MustNewDevelopment() LoggerImpl {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic("slog_zap: zap.NewDevelopment(): " + err.Error())
	}
	return Wrap(l)
}

type Logger struct {
	l *zap.SugaredLogger
}

func (s *Logger) WithField(key string, value interface{}) LoggerImpl {
	return &Logger{l: s.l.With(key, value)}
}

func (s *Logger) Warning(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s *Logger) Debugf(format string, args ...interface{}) {
	s.l.Debugf(format, args...)
}

func (s *Logger) Infof(format string, args ...interface{}) {
	s.l.Infof(format, args...)
}

func (s *Logger) Warnf(format string, args ...interface{}) {
	s.l.Warnf(format, args...)
}

func (s *Logger) Errorf(format string, args ...interface{}) {
	s.l.Errorf(format, args...)
}

func (s *Logger) Fatalf(format string, args ...interface{}) {
	s.l.Fatalf(format, args...)
}

func (s *Logger) Panicf(format string, args ...interface{}) {
	s.l.Panicf(format, args...)
}

func (s *Logger) Debug(args ...interface{}) {
	s.l.Debug(args...)
}

func (s *Logger) Info(args ...interface{}) {
	s.l.Info(args...)
}

func (s *Logger) Warn(args ...interface{}) {
	s.l.Warn(args...)
}

func (s *Logger) Error(args ...interface{}) {
	s.l.Error(args...)
}

func (s *Logger) Fatal(args ...interface{}) {
	s.l.Fatal(args...)
}

func (s *Logger) Panic(args ...interface{}) {
	s.l.Panic(args...)
}

func (s *Logger) Flush() error {
	return s.l.Sync()
}
