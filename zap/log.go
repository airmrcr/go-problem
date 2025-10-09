// Copyright (C) 2025 neocotic
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package problemzap

import (
	"context"
	"log/slog"

	"github.com/neocotic/go-problem"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// defaultZapLevel is the zapcore.Level used when one could not be derived.
	defaultZapLevel = zapcore.ErrorLevel

	// badZapKey is used when a zap.Logger is passed an unexpected arg key.
	badZapKey = "!BADKEY"
)

// Field returns a zapcore.Field populated with fields of the given problem.Problem.
//
// The name of the field is taken from problem.DefaultGenerator.LogArgKey and defaults to problem.DefaultLogArgKey if
// empty. NamedField should be used if a specific name is needed for the field.
func Field(prob *problem.Problem) zapcore.Field {
	return FieldUsing(problem.DefaultGenerator, prob)
}

// FieldUsing returns a zapcore.Field populated with fields of the given problem.Problem.
//
// The name of the field is taken from problem.Generator.LogArgKey and defaults to problem.DefaultLogArgKey if empty.
// NamedField should be used if a specific name is needed for the field.
func FieldUsing(gen *problem.Generator, prob *problem.Problem) zapcore.Field {
	key := gen.LogArgKey
	if key == "" {
		key = problem.DefaultLogArgKey
	}

	return NamedField(key, prob)
}

// GlobalLogger returns a problem.Logger that uses zap.L.
func GlobalLogger() problem.Logger {
	return LoggerFromContext(zap.L(), func(_ context.Context, logger *zap.Logger) *zap.Logger {
		return logger
	})
}

// GlobalLoggerContext returns a problem.Logger that uses zap.L while passing the context to the function provided to
// return the most appropriate zap.Logger.
//
// This can be useful for cases where the context is used to further enrich logs.
func GlobalLoggerContext(handleCtx func(ctx context.Context, logger *zap.Logger) *zap.Logger) problem.Logger {
	return LoggerFromContext(zap.L(), handleCtx)
}

// LoggerFrom returns a problem.Logger that uses the given zap.Logger.
func LoggerFrom(logger *zap.Logger) problem.Logger {
	return LoggerFromContext(logger, func(_ context.Context, _ *zap.Logger) *zap.Logger {
		return logger
	})
}

// LoggerFromContext returns a problem.Logger that uses the given zap.Logger while passing the context to the function
// provided to return the most appropriate zap.Logger.
//
// This can be useful for cases where the context is used to further enrich logs.
func LoggerFromContext(logger *zap.Logger, handleCtx func(ctx context.Context, logger *zap.Logger) *zap.Logger) problem.Logger {
	return func(ctx context.Context, level problem.LogLevel, msg string, args ...any) {
		handleCtx(ctx, logger).Log(convertLevel(level), msg, extractFields(args)...)
	}
}

// NamedField returns a zapcore.Field with the given key, populated with fields of the given problem.Problem.
func NamedField(key string, prob *problem.Problem) zapcore.Field {
	if prob == nil {
		return zap.Skip()
	}

	var fields []zapcore.Field
	logInfo := prob.LogInfo()

	if prob.Code != "" {
		fields = append(fields, zap.String("code", string(prob.Code)))
	}
	if prob.Detail != "" {
		fields = append(fields, zap.String("detail", prob.Detail))
	}
	if err := prob.Unwrap(); err != nil {
		fields = append(fields, zap.String("error", err.Error()))
	}
	if len(prob.Extensions) > 0 {
		fields = append(fields, mapField("extensions", prob.Extensions))
	}
	if prob.Instance != "" {
		fields = append(fields, zap.String("instance", prob.Instance))
	}
	if logInfo.Stack != "" {
		fields = append(fields, zap.String("stack", logInfo.Stack))
	}
	if prob.Status != 0 {
		fields = append(fields, zap.Int("status", prob.Status))
	}
	if prob.Title != "" {
		fields = append(fields, zap.String("title", prob.Title))
	}
	if prob.Type != "" {
		fields = append(fields, zap.String("type", prob.Type))
	}
	if logInfo.UUID != "" {
		fields = append(fields, zap.String("uuid", logInfo.UUID))
	}

	return zap.Dict(key, fields...)
}

// NoopLogger returns a problem.Logger that does nothing.
func NoopLogger() problem.Logger {
	return LoggerFromContext(zap.NewNop(), func(_ context.Context, logger *zap.Logger) *zap.Logger {
		return logger
	})
}

// argsToField turns a prefix of the non-empty args slice into a zapcore.Field and returns the unconsumed portion of the
// slice.
//
// If args[0] is a zapcore.Field, it returns it.
// If args[0] is a slog.Attr, it returns it and zapcore.Field contains the slog.Attr key-value pair.
// If args[0] is a string, it treats the first two elements as a key-value pair.
// Otherwise, it treats args[0] as a value with a missing key.
func argsToField(args []any) (zapcore.Field, []any) {
	switch k := args[0].(type) {
	case string:
		if len(args) == 1 {
			return zap.String(badZapKey, k), nil
		}
		switch v := args[1].(type) {
		case *problem.Problem:
			return NamedField(k, v), args[2:]
		case problem.Problem:
			return NamedField(k, &v), args[2:]
		default:
			return zap.Any(k, v), args[2:]
		}
	case slog.Attr:
		switch v := k.Value.Any().(type) {
		case *problem.Problem:
			return NamedField(k.Key, v), args[1:]
		case problem.Problem:
			return NamedField(k.Key, &v), args[1:]
		default:
			return zap.Any(k.Key, v), args[1:]
		}
	case zapcore.Field:
		return k, args[1:]
	case *problem.Problem:
		return NamedField(badZapKey, k), args[1:]
	case problem.Problem:
		return NamedField(badZapKey, &k), args[1:]
	default:
		return zap.Any(badZapKey, k), args[1:]
	}
}

// convertLevel returns the zapcore.Level representation of the problem.LogLevel, where possible, otherwise
// defaultZapLevel.
func convertLevel(level problem.LogLevel) zapcore.Level {
	switch level {
	case problem.LogLevelDebug:
		return zapcore.DebugLevel
	case problem.LogLevelInfo:
		return zapcore.InfoLevel
	case problem.LogLevelWarn:
		return zapcore.WarnLevel
	case problem.LogLevelError:
		return zapcore.ErrorLevel
	default:
		return defaultZapLevel
	}
}

// extractFields consumes all args into a slice of zapcore.Field.
func extractFields(args []any) (fields []zapcore.Field) {
	var f zapcore.Field
	for len(args) > 0 {
		f, args = argsToField(args)
		fields = append(fields, f)
	}
	return
}

// mapField returns a zapcore.Field containing all entries within the given map.
func mapField(key string, m map[string]any) zapcore.Field {
	var fields []zapcore.Field
	for k, v := range m {
		fields = append(fields, zap.Any(k, v))
	}
	return zap.Dict(key, fields...)
}
