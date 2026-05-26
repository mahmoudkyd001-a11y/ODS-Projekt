// Edit this file, as it is a specific handler function for your service
package rest

import (
	"myservice/core/log"
	"myservice/core/tracing"

	"net/http"

	"github.com/labstack/echo/v5"
)

// Returns a list of all the stores.
func GetAllStores(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("GetAllStores")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("GetAllStores failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	_ = i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	log.Debug().
		Msg("Received request for GetAllStores, running business logic")

		/*
		 * Business Logic
		 */

		// 200 => An array of Store objects.
	return (*c).String(http.StatusNotImplemented, "Temporary handler stub.")
}
