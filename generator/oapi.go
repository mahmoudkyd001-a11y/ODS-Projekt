package generator

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func updateOAPIOperation(op *openapi3.Operation, opID string, opSummary string, opDefault string) {
	op.OperationID = opID
	op.Summary = opSummary
	if resp := op.Responses.Map()[opDefault]; resp != nil {
		op.Responses.Set("default", resp)
	}
}

func createOAPIResponse(rDesc string) *openapi3.Response {
	r := openapi3.NewResponse()
	r.Description = &rDesc
	return r
}
