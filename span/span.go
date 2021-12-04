package sentry

import (
	"context"
	"runtime"

	"github.com/getsentry/sentry-go"
)

func Start(ctx context.Context) *sentry.Span {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	span := sentry.StartSpan(ctx, frame.Function)

	return span
}
