package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// CreateProductZapLogger 创建一个生产级别的 zap 日志记录器。
//
// @Summary 创建一个生产级别的 zap 日志记录器
// @Description 根据给定的日志级别，大小和数量限制，创建一个生产级别的 zap 日志记录器。
// @Tags Logging
func CreateProductZapLogger(level zapcore.Level, maxSize, maxBackups, maxAge int, compress bool) (*zap.Logger, error) {
	// 配置 lumberjack 日志轮转，用来控制日志记录的大小
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./app.log", // 日志文件路径
		MaxSize:    maxSize,     // 每个日志文件的最大大小（单位：MB）
		MaxBackups: maxBackups,  // 保留的旧日志文件数量
		MaxAge:     maxAge,      // 保留旧日志文件的最长天数
		Compress:   compress,    // 是否压缩旧日志文件
	}

	// 创建 zap 的核心配置
	writeSyncer := zapcore.AddSync(lumberjackLogger)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 使用 JSON 格式编码日志
		writeSyncer,
		level, // 设置日志级别
	)

	// 创建 zap logger
	logger := zap.New(core, zap.AddCaller()) // 添加调用者信息
	return logger, nil
	//defer logger.Sync() // 确保日志缓冲区被刷新
}
