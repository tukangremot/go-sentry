package sentry

import (
	"context"
	"runtime"

	"github.com/getsentry/sentry-go"
)

func Start(ctx context.Context, options ...sentry.SpanOption) *sentry.Span {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return sentry.StartSpan(ctx, frame.Function, options...)
}
