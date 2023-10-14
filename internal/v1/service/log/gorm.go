package log

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogger struct {
	logger                *zerolog.Logger
	SlowThreshold         time.Duration
	SkipErrRecordNotFound bool
	Debug                 bool
	LogLevel              logger.LogLevel
}

func NewGormLogger(logger *zerolog.Logger) *GormLogger {
	return &GormLogger{
		logger:                logger,
		SkipErrRecordNotFound: false,
		Debug:                 true,
	}
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	g.LogLevel = level
	return g
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		g.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Str("message", fmt.Sprintf(s, i...)).Send()
	}
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		g.logger.Warn().Str("requestId", middleware.GetReqID(ctx)).Str("message", fmt.Sprintf(s, i...)).Send()
	}
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		g.logger.Error().Str("requestId", middleware.GetReqID(ctx)).Str("message", fmt.Sprintf(s, i...)).Send()
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, _ := fc()
	message := fmt.Sprintf("query: %s - elapsed %d - file %s", sql, elapsed, utils.FileWithLineNum())
	switch {
	case err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && g.SkipErrRecordNotFound):
		g.logger.Error().Str("requestId", middleware.GetReqID(ctx)).Str("error", err.Error()).Str("query", sql).Float64("elapsed", float64(elapsed.Nanoseconds())/1e6).Str("file", utils.FileWithLineNum()).Msg(message)
	case g.SlowThreshold != 0 && elapsed > g.SlowThreshold:
		g.logger.Warn().Str("requestId", middleware.GetReqID(ctx)).Str("query", sql).Float64("elapsed", float64(elapsed.Nanoseconds())/1e6).Str("file", utils.FileWithLineNum()).Msg(message)
	case g.LogLevel == logger.Info:
		g.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Str("query", sql).Float64("elapsed", float64(elapsed.Nanoseconds())/1e6).Str("file", utils.FileWithLineNum()).Msg(message)
	}
}
