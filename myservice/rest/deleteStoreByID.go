// Edit this file, as it is a specific handler function for your service
package rest

import (
	"myservice/core/log"
	"myservice/core/tracing"

	"net/http"

	"github.com/labstack/echo/v5"
)

// Deletes a store with the given store id.
func DeleteStoreByID(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("DeleteStoreByID")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("DeleteStoreByID failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	_ = i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	id := (*c).Param("id")

	log.Debug().
		Any("id", id).
		Msg("Received request for DeleteStoreByID, running business logic")

		/*
		 * Business Logic
		 */

		// 200 => A store was deleted successfully.
		// 400 => Invalid store id.
		// 404 => Store with given id wasn't found.
	return (*c).String(http.StatusNotImplemented, "Temporary handler stub.")
}
