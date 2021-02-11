// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture   string          `json:"architecture"`
	ImageType      string          `json:"image_type"`
	UploadRequests []UploadRequest `json:"upload_requests"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Version string `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Status string      `json:"status"`
	Type   UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"xHAeGa/a7f4bAAD//7APRnV6IAAA",
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
