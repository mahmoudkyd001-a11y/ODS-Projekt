// Edit this file, as it is a specific handler function for your service
package rest

import (
	"antragservice6/core/log"
	"antragservice6/core/tracing"
	"antragservice6/web/pages"

	"github.com/labstack/echo/v5"
)

// Antragsformular anzeigen
func GetAntragForm(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("GetAntragForm")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("GetAntragForm failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	lzr := i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	log.Debug().
		Msg("Received request for GetAntragForm, running business logic")

	/*
	 * Business Logic
	 */

	// 200 => Formular anzeigen
	return pages.AntragForm().Render((*c).Request().Context(), (*c).Response())
}
