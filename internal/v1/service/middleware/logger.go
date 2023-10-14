package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func NewStructuredLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{Logger: logger})
}

type StructuredLogger struct {
	Logger zerolog.Logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: l.Logger}
	logFields := map[string]interface{}{
		"ts":         time.Now().UTC().Format(time.RFC1123),
		"httpScheme": "http",
		"httpProto":  r.Proto,
		"httpMethod": r.Method,
		"remoteAddr": r.RemoteAddr,
		"userAgent":  r.UserAgent(),
		"uri":        r.RequestURI,
	}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["requestId"] = reqID
	}

	entry.Logger = entry.Logger.With().Fields(logFields).Logger()
	entry.Logger.Info().Msg("Entering request")

	return entry
}

type StructuredLoggerEntry struct {
	Logger zerolog.Logger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger.Info().
		Int("resp_status", status).
		Int("resp_bytes_length", bytes).
		Float64("resp_elapsed_ms", float64(elapsed.Nanoseconds())/1000000.0).
		Msg("Exiting request")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger.Error().
		Str("stack", string(stack)).
		Interface("panic", v).
		Msg("Panic")
}
