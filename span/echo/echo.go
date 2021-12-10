package echo

import (
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	span "github.com/tukangremot/go-sentry/span"
)

func Start(context echo.Context) *sentry.Span {
	return span.Start(context.Request().Context(), sentry.TransactionName(context.Request().URL.Path))
}
