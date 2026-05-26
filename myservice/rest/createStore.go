// Edit this file, as it is a specific handler function for your service
package rest

import (
	"myservice/core/log"
	"myservice/core/tracing"
	"myservice/entities"
	"myservice/rest/middleware"

	"net/http"

	"github.com/labstack/echo/v5"
)

// Creates a new store.
func CreateStore(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("CreateStore")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("CreateStore failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	_ = i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	bodyContent := new(entities.Store)
	if err := middleware.Bind(c, bodyContent); err != nil {
		log.Error().Err(err).Str("traceId", traceId).Str("spanId", spanId).Msg("Failed to bind bodyContent (typeStore) to reqest at CreateStore")
		return (*c).String(http.StatusBadRequest, "Invalid request body")
	}
	// TODO improve body validation
	if err := bodyContent.Validate(); err != nil {
		return (*c).String(http.StatusUnprocessableEntity, err.Error())
	}

	log.Debug().
		Msg("Received request for CreateStore, running business logic")

		/*
		 * Business Logic
		 * bodyContent: contains parsed request body
		 */

		// 200 => A store was created successfully.
		// 400 => Invalid store properties.
	return (*c).String(http.StatusNotImplemented, "Temporary handler stub.")
}
