package echo

import (
	"runtime"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

func Start(context echo.Context) *sentry.Span {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return sentry.StartSpan(context.Request().Context(), frame.Function, sentry.TransactionName(context.Request().URL.Path))
}
