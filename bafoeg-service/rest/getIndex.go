// Edit this file, as it is a specific handler function for your service
package rest

import (
	"bafoegservice/core"
	"bafoegservice/core/log"
	"bafoegservice/core/tracing"
	"bafoegservice/web/pages"

	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// successfully deliver index page
func GetIndex(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", "/index.html").Msg("GetIndex")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
	 	log.Error().Err(err).Msg("GetIndex failed")
		return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	lzr := i18n.NewLocalizer(core.Bundle, core.Language(c))
	return Render(c, http.StatusOK, pages.Index(lzr, core.AppConfig.Title))
}
