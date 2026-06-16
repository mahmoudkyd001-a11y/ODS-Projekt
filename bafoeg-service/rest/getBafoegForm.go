// Edit this file, as it is a specific handler function for your service
package rest

import (
	"bafoegservice/core/log"
	"bafoegservice/core/tracing"
	"bafoegservice/web/pages"

	"github.com/labstack/echo/v5"
)

// BAföG Antrag anzeigen
func GetBafoegForm(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("GetBafoegForm")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("GetBafoegForm failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	lzr := i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	log.Debug().
		Msg("Received request for GetBafoegForm, running business logic")

	/*
	 * Business Logic
	 */

	// 200 => Formular anzeigen
	return pages.BafoegForm().Render((*c).Request().Context(), (*c).Response())
}
