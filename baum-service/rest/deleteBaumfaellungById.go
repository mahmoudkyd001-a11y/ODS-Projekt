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

// Antrag löschen
func DeleteBaumfaellungById(c *echo.Context) error {
	// trace span
	ctx := (*c).Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", (*c).Request().URL.Path).Msg("DeleteBaumfaellungById")

	/* Remove this comment to use session
	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("DeleteBaumfaellungById failed")
	 	return (*c).NoContent(http.StatusInternalServerError)
	}
	*/

	/* Remove this comment to use localization
	lzr := i18n.NewLocalizer(core.Bundle, core.Language(c))
	*/

	id := (*c).Param("id")

	log.Debug().
		Any("id", id).
		Msg("Received request for DeleteBaumfaellungById, running business logic")

		/*
		 * Business Logic
		 */

		// 200 => Antrag gelöscht
	return (*c).String(http.StatusNotImplemented, "Temporary handler stub.")
}
