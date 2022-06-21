package logging_zap

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogConfig struct {
	Level   zap.AtomicLevel    `json:"level" yaml:"level"`
	Logfile *lumberjack.Logger `json:"logfile" yaml:"logfile"`
}

func (c *LogConfig) NewLogger(ginDebugMode bool) LoggerImpl {
	logLevel := zapcore.DebugLevel

	if !ginDebugMode {
		logLevel = c.Level.Level()
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:      "file",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	})

	// 实现判断日志等级的interface
	targetLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel
	})

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(c.Logfile)), targetLevel),
	)

	logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return Wrap(logger)
}
