// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
