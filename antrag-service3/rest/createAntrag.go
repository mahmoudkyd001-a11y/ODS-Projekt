// Edit this file, as it is a specific handler function for your service
package rest

import (
	"antragservice3/core/log"
	"antragservice3/core/tracing"
	"antragservice3/entities"
	"antragservice3/rest/middleware"

	"net/http"

	"github.com/labstack/echo/v5"
)

// Antrag einreichen
func CreateAntrag(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("CreateAntrag")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("CreateAntrag failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	lzr := i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	bodyContent := new(entities.Antrag)
	if err := middleware.Bind(c, bodyContent); err != nil {
		log.Error().Err(err).Str("traceId", traceId).Str("spanId", spanId).Msg("Failed to bind bodyContent (typeAntrag) to reqest at CreateAntrag")
		return (*c).String(http.StatusBadRequest, "Invalid request body")
	}
	// TODO improve body validation
	if err := bodyContent.Validate(); err != nil {
		return (*c).String(http.StatusUnprocessableEntity, err.Error())
	}

	log.Debug().
		Msg("Received request for CreateAntrag, running business logic")

		/*
		 * Business Logic
		 * bodyContent: contains parsed request body
		 */

		// 201 => Antrag erstellt
	return (*c).String(http.StatusNotImplemented, "Temporary handler stub.")
}
