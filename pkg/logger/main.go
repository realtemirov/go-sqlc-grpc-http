package logger

import (
	"os"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggingConfig struct {
	Level             string        `json:"level" yaml:"level" envDefault:"info"`
	LogFirstN         int           `json:"log_first_n" yaml:"log_first_n" envDefault:"3"`
	LogThereAfter     int           `json:"log_there_after" yaml:"log_there_after" envDefault:"10"`
	LogInterval       time.Duration `json:"log_interval" yaml:"log_interval" envDefault:"1s"`
	ProjectName       string        `json:"project_name" yaml:"project_name"`
	ConsoleLogEnabled bool          `json:"console_log_enabled" yaml:"console_log_enabled" envDefault:"true"`
}

func lowerCaseLevelEncoder(
	level zapcore.Level,
	enc zapcore.PrimitiveArrayEncoder,
) {
	if level == zap.PanicLevel || level == zap.DPanicLevel {
		enc.AppendString("error")
		return
	}

	zapcore.LowercaseColorLevelEncoder(level, enc)
}

func New(config *LoggingConfig) *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	globalLogLevel, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		panic(err)
	}

	level := zap.NewAtomicLevelAt(globalLogLevel)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeLevel = lowerCaseLevelEncoder
	productionCfg.StacktraceKey = ""

	var (
		encoder zapcore.Encoder
	)

	if config.ConsoleLogEnabled {
		encoder = zapcore.NewConsoleEncoder(productionCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(productionCfg)
	}

	jsonOutCore := zapcore.NewCore(encoder, stdout, level)

	samplingCore := zapcore.NewSamplerWithOptions(
		jsonOutCore,
		config.LogInterval,
		config.LogFirstN,
		config.LogThereAfter,
	)

	logger := zap.New(samplingCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)).Named(config.ProjectName)

	// Return a logger wrapped with a conditional Sync for compatibility with environments like Docker.
	return logger.WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return &syncableCore{c, stdout}
	}))
}

// syncableCore wraps zapcore.Core and conditionally syncs the logger.
type syncableCore struct {
	zapcore.Core
	writer zapcore.WriteSyncer
}

// Sync handles syncing of the logger, only calling Sync() on stdout for non-Windows systems.
func (c *syncableCore) Sync() error {
	if runtime.GOOS != "windows" {
		// Sync only if not Windows to avoid stdout errors
		if err := c.writer.Sync(); err != nil && err.Error() != "sync /dev/stdout: invalid argument" {
			return err
		}
	}
	return nil
}
