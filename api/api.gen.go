// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/Peyton232/todo/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// returns all users
	// (GET /users)
	FindUsers(ctx echo.Context, params FindUsersParams) error
	// Creates a new User
	// (POST /users)
	AddUser(ctx echo.Context) error
	// Update an existing user
	// (PUT /users)
	UpdateUser(ctx echo.Context) error
	// Deletes a user by ID
	// (DELETE /users/{id})
	DeleteUser(ctx echo.Context, id int64) error
	// Returns a user by ID
	// (GET /users/{id})
	FindUserByID(ctx echo.Context, id int64) error
	// get all Achievements
	// (GET /users/{id}/Achievements/)
	GetAchievements(ctx echo.Context, id int64) error
	// add new Achievements
	// (POST /users/{id}/Achievements/)
	AddAchievements(ctx echo.Context, id int64) error
	// update Achievement
	// (PUT /users/{id}/Achievements/)
	UpdateAchievements(ctx echo.Context, id int64) error
	// delete an Achievement by name
	// (DELETE /users/{id}/Achievements/{name})
	DeleteAchievement(ctx echo.Context, id int64, name string) error
	// get an Achievement by name
	// (GET /users/{id}/Achievements/{name})
	GetAchievement(ctx echo.Context, id int64, name string) error
	// get alerts
	// (GET /users/{id}/alerts)
	GetAlerts(ctx echo.Context, id int64) error
	// insert new alert
	// (POST /users/{id}/alerts)
	AddAlert(ctx echo.Context, id int64) error
	// update an alert
	// (PUT /users/{id}/alerts)
	UpdateAlert(ctx echo.Context, id int64) error
	// delete an Achievement by name
	// (DELETE /users/{id}/alerts/{name})
	DeleteAlert(ctx echo.Context, id int64, name string) error
	// delete an alert by name
	// (GET /users/{id}/alerts/{name})
	GetAlert(ctx echo.Context, id int64, name string) error
	// returns all of this users buckets
	// (GET /users/{id}/buckets)
	FindBuckets(ctx echo.Context, id int32) error
	// Creates a new Bucket
	// (POST /users/{id}/buckets)
	AddBucket(ctx echo.Context, id int32) error
	// Update an existing bucket
	// (PUT /users/{id}/buckets)
	UpdateBucket(ctx echo.Context, id int32) error
	// Deletes a bucket by name
	// (DELETE /users/{id}/buckets/{name})
	DeleteBucket(ctx echo.Context, id int64, name string) error
	// Returns a bucket by name
	// (GET /users/{id}/buckets/{name})
	FindBucketByName(ctx echo.Context, id int32, name string) error
	// get transactions
	// (GET /users/{id}/buckets/{name}/transactions)
	GetTransactions(ctx echo.Context, id int64, name string) error
	// Inserts a new transaction
	// (POST /users/{id}/buckets/{name}/transactions)
	InsertTransaction(ctx echo.Context, id int64, name string) error
	// update a transaction
	// (PUT /users/{id}/buckets/{name}/transactions)
	UpdateTransaction(ctx echo.Context, id int64, name string) error
	// delete a transaction
	// (DELETE /users/{id}/buckets/{name}/transactions/{timestamp})
	DeeleteTransaction(ctx echo.Context, id int64, name string, timestamp string) error
	// get a transaction by timestamp
	// (GET /users/{id}/buckets/{name}/transactions/{timestamp})
	GetTransaction(ctx echo.Context, id int64, name string, timestamp string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindUsers converts echo context to params.
func (w *ServerInterfaceWrapper) FindUsers(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindUsersParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindUsers(ctx, params)
	return err
}

// AddUser converts echo context to params.
func (w *ServerInterfaceWrapper) AddUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddUser(ctx)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateUser(ctx)
	return err
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUser(ctx, id)
	return err
}

// FindUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindUserByID(ctx, id)
	return err
}

// GetAchievements converts echo context to params.
func (w *ServerInterfaceWrapper) GetAchievements(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAchievements(ctx, id)
	return err
}

// AddAchievements converts echo context to params.
func (w *ServerInterfaceWrapper) AddAchievements(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddAchievements(ctx, id)
	return err
}

// UpdateAchievements converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateAchievements(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateAchievements(ctx, id)
	return err
}

// DeleteAchievement converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAchievement(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteAchievement(ctx, id, name)
	return err
}

// GetAchievement converts echo context to params.
func (w *ServerInterfaceWrapper) GetAchievement(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAchievement(ctx, id, name)
	return err
}

// GetAlerts converts echo context to params.
func (w *ServerInterfaceWrapper) GetAlerts(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAlerts(ctx, id)
	return err
}

// AddAlert converts echo context to params.
func (w *ServerInterfaceWrapper) AddAlert(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddAlert(ctx, id)
	return err
}

// UpdateAlert converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateAlert(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateAlert(ctx, id)
	return err
}

// DeleteAlert converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAlert(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteAlert(ctx, id, name)
	return err
}

// GetAlert converts echo context to params.
func (w *ServerInterfaceWrapper) GetAlert(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAlert(ctx, id, name)
	return err
}

// FindBuckets converts echo context to params.
func (w *ServerInterfaceWrapper) FindBuckets(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindBuckets(ctx, id)
	return err
}

// AddBucket converts echo context to params.
func (w *ServerInterfaceWrapper) AddBucket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddBucket(ctx, id)
	return err
}

// UpdateBucket converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateBucket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateBucket(ctx, id)
	return err
}

// DeleteBucket converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteBucket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteBucket(ctx, id, name)
	return err
}

// FindBucketByName converts echo context to params.
func (w *ServerInterfaceWrapper) FindBucketByName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindBucketByName(ctx, id, name)
	return err
}

// GetTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTransactions(ctx, id, name)
	return err
}

// InsertTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) InsertTransaction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.InsertTransaction(ctx, id, name)
	return err
}

// UpdateTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTransaction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTransaction(ctx, id, name)
	return err
}

// DeeleteTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) DeeleteTransaction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// ------------- Path parameter "timestamp" -------------
	var timestamp string

	err = runtime.BindStyledParameterWithLocation("simple", false, "timestamp", runtime.ParamLocationPath, ctx.Param("timestamp"), &timestamp)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timestamp: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeeleteTransaction(ctx, id, name, timestamp)
	return err
}

// GetTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransaction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// ------------- Path parameter "timestamp" -------------
	var timestamp string

	err = runtime.BindStyledParameterWithLocation("simple", false, "timestamp", runtime.ParamLocationPath, ctx.Param("timestamp"), &timestamp)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timestamp: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTransaction(ctx, id, name, timestamp)
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
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.FindUsers)
	router.POST(baseURL+"/users", wrapper.AddUser)
	router.PUT(baseURL+"/users", wrapper.UpdateUser)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUser)
	router.GET(baseURL+"/users/:id", wrapper.FindUserByID)
	router.GET(baseURL+"/users/:id/Achievements/", wrapper.GetAchievements)
	router.POST(baseURL+"/users/:id/Achievements/", wrapper.AddAchievements)
	router.PUT(baseURL+"/users/:id/Achievements/", wrapper.UpdateAchievements)
	router.DELETE(baseURL+"/users/:id/Achievements/:name", wrapper.DeleteAchievement)
	router.GET(baseURL+"/users/:id/Achievements/:name", wrapper.GetAchievement)
	router.GET(baseURL+"/users/:id/alerts", wrapper.GetAlerts)
	router.POST(baseURL+"/users/:id/alerts", wrapper.AddAlert)
	router.PUT(baseURL+"/users/:id/alerts", wrapper.UpdateAlert)
	router.DELETE(baseURL+"/users/:id/alerts/:name", wrapper.DeleteAlert)
	router.GET(baseURL+"/users/:id/alerts/:name", wrapper.GetAlert)
	router.GET(baseURL+"/users/:id/buckets", wrapper.FindBuckets)
	router.POST(baseURL+"/users/:id/buckets", wrapper.AddBucket)
	router.PUT(baseURL+"/users/:id/buckets", wrapper.UpdateBucket)
	router.DELETE(baseURL+"/users/:id/buckets/:name", wrapper.DeleteBucket)
	router.GET(baseURL+"/users/:id/buckets/:name", wrapper.FindBucketByName)
	router.GET(baseURL+"/users/:id/buckets/:name/transactions", wrapper.GetTransactions)
	router.POST(baseURL+"/users/:id/buckets/:name/transactions", wrapper.InsertTransaction)
	router.PUT(baseURL+"/users/:id/buckets/:name/transactions", wrapper.UpdateTransaction)
	router.DELETE(baseURL+"/users/:id/buckets/:name/transactions/:timestamp", wrapper.DeeleteTransaction)
	router.GET(baseURL+"/users/:id/buckets/:name/transactions/:timestamp", wrapper.GetTransaction)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+ybX2/juBHAvwrB9qEFjDj3B33wU+O6VwRt7w7d5GmRB1oa2WwpUktS8RqBv3sxJGVJ",
	"FiUrt8bae/bTek16OJz5zXAYjd5oovJCSZDW0NkbNckacuY+PiRrDq+Qg7T430KrArTl4AaZH0zxcwom",
	"0bywXEk6o2tmiF1zQ5YAkuznTajdFkBndKmUACbpbkIlywEFhBFjNZcrHNCwYTp9yFXp185LYXkh4JeM",
	"zu7v7r/bC5NlvgRd/+TfYAxbxYS6KZ9KrlHnj37pl92EPgjQkf2lYEEzv3pHvXzsKtVEXGheJv+DmCX3",
	"u2zbMSm1BmmJHydcErsGsnRSiFEkY5pOxlhmzY1VeosrcAu5W/SPGjI6o3+Y1t6fBtdPnzSThiVOjd1e",
	"HtOabZtOa2v7M8uBqKyhY+3x2nDqFXQm1CYfwsbvsJoqo+ToUkSUEBydJsSW5GzFk5gGVlkmur90X1em",
	"dtvYazLGyDG4qrUmlYcRgr9rrXSXgUSlke24ycSNTWimdM4snVEu7Q/f11vj0sLKO7qBZUxQNTw5Am1Y",
	"sMnur1pl3Fu8rXjGJZMJZ+KfUm0EpCv4F7xCxMBrtSEmZ9ryRIBxmJQGNOEyupU1KwouwZgBcSul0oYg",
	"Q1KF24mJ49KUwvbIEvg1Ot3PMnHbKmnXYvuBCeZDqS2jJidXEpC/FEgBmrjfRSUejyPc11Fv7RNZM2xH",
	"J5km8UBsQ8So1NIS1vFRmTNJNLCULQUQ+FwIJhkOuvVaa3XiFD4XIE3EQDwEJzeESRQKRgJRmkhlo6nC",
	"8hyMZXkRifpq6KhCB0avZaLln9FPaGQh0FYfe87KvDpjR2Xh5ukbycIMj613SHOnXESOz3HjBYVTLCKJ",
	"R2qBZ8k/lUB4eoD0iGRW1ClnSJ8qMwVH/6SjUfW0LXxU8RxIhnPIn1LIWCksCZH956NO51jEVGq97F52",
	"O5dZMuUzuLQscREGOeMCLVNwCyz/q9mw1Qr0HVe0inr6wX9HHn59JE/A8k4wURzJlCZPKlVkXqYrsIQV",
	"BWoJOje/ZB9Av/IEha2tLWbTab3O1E2Z4lxu0YYUpXghdEJfQRu/yHd393f37mAuQLKC0xn9wX01oQWz",
	"a0fDFH3mPq0gkkH+A7bU0hAmhPOuIZlWuXO22RoL+JHZOk/jOc+SBIwhFg2CgeKSwmNKZ/QnLtNntx5q",
	"gI6ybvGPh8vm7DPPy5z4XIS+1eCyN7GKaKcTRe/QGf1Ugt7Wphc852gFT5A7xI7RuHtBEkyhpPER/f39",
	"feX0UB2zohA8cRuZ/tf4dFivMCq0XBrpBNauQ4YzY6WOB8eR/C6NhhTxRUpsZYn5NrGQEghzJtSUee7O",
	"RKoPUXBxrEyEmr9pYBYMYUTCpqoEHCQps0tm4I4sSq8/ztKAQtXG3SPaxDykDhjqYxWMnat0ezJTeJ/0",
	"+MAqwtIU/3G0W6WBNlOG1SXsvpCc36zcJQLS9nrFe1FGAMFEQJgng8mUlEXKLKChuUZGWAeEZzfjDCw8",
	"Bxa8ijcChgjwTvKVGzeWy5UvCXCWP2imbzzdeR4E2MhZ7r9HiAyXKxGOFcwZKVE+hzwuiClxb5F8sXA/",
	"D5gMHjGPCzxWqkgP6oQzBU/H+khxlUHb633ny19+HHe+/NjduNPEq5Fekk8Xe4d4T2zJ4wL1Gy4XDty2",
	"9+bjorcomG/d4DucloFN1l/NZ1cc14du9Qy0g7p5pTHT3oLSFbpCkObsDhL/AHswPpKKi4JhXxaOc07r",
	"StgtFA9d1rx2XiQ0UU/3loz4rasbBsF4SNPLA+P0tUgbhe7lsen6dpla1gXS16lRjqh68ZSi7TrY9ZWt",
	"5b6+aW47XqreOO1wGswXyrgbqu9ENZjv8JzoPYPfkJ4RxfYxnn0F2J5xXpwnhyvK8Hf1tpKRpcOzm/7F",
	"D/9I+HIDshfIKD5YHTor990RXF0wTFy7/rvhdsNtGwWnZu0gD9bPT6IImrXakDAnxl41cg23Dv/oaMR9",
	"Ayde8EXDuaz3esGlQfWx0nNTozeLMPD7LNWCn7tFmvPruW8RfdpdLnUdoo5fGuLkhevCdcN37qvBt8ff",
	"IVbxI/DUl4CLoLS3HqsC7OIrsW8Pt99W69e/iue+qti6IXXVSHkd+4r5RhPT0R6VMNd3qTBiCkh4xpMK",
	"ne6Tp3kQfgRA38OQfjF9Z2xA6evt+kZbUKo2Xt+VVDEysiUldCFnSodnWrH7wLxqED43GacvyCoWuh4I",
	"ljl388tRBS+/AaaOt5O2wFwllfM9lWdow+lX61tqxFlWu4gerqPvCY3OnBCJrd4cVzYd6c4ZR/BX7M/p",
	"Lf7qZDjQanKCMjDSDhSWvuiGoAqAI5eAunvkkBhGgvH66rL59mc/4YzJ7ux0XPmp2sVn+J4QUtm08caH",
	"GWxHak2M3FCf2uMXelHdv852Pgjf+eih9Sbi8QcQDTdd7GOIFkpjHka0X0tqo/fopj21ZlwNfKevLtu4",
	"DeJ15qch79D0EgPBc1tdg2z7fePBRySD4eBr2isNhxtpQ89A2pCNLwumb/tXPUc9IhnkcwFu1vUBOnnH",
	"i7eR9eq3bW9RcdK/qx+m3v5urOZMLK+bLhkqh2+A3wA/V8NPP7NOoAH9WkFZalG/R21VqpbuXem7xivV",
	"rOB097L7fwAAAP//Bdkn4jBHAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}