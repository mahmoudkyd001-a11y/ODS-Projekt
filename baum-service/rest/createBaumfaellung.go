// Edit this file, as it is a specific handler function for your service
package rest

import (
	"baumservice/core/log"
	"baumservice/core/tracing"
	"baumservice/entities"
	"baumservice/rest/middleware"
	"baumservice/usecases"
	"baumservice/web/pages"

	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/yaml.v2"
)

// Baumfällantrag einreichen
func CreateBaumfaellung(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("CreateBaumfaellung")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("CreateBaumfaellung failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	lzr := i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	bodyContent := new(entities.Baumfaellung)
	if err := middleware.Bind(c, bodyContent); err != nil {
		log.Error().Err(err).Str("traceId", traceId).Str("spanId", spanId).Msg("Failed to bind bodyContent (typeBaumfaellung) to reqest at CreateBaumfaellung")
		return (*c).String(http.StatusBadRequest, "Invalid request body")
	}
	// TODO improve body validation
	if err := bodyContent.Validate(); err != nil {
		return (*c).String(http.StatusUnprocessableEntity, err.Error())
	}

	log.Debug().
		Msg("Received request for CreateBaumfaellung, running business logic")

		/*
		 * Business Logic
		 * bodyContent: contains parsed request body
		 */

		// 201 => Antrag erstellt
	return (*c).String(http.StatusNotImplemented, "Temporary handler stub.")
}
