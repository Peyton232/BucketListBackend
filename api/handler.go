package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	. "github.com/Peyton232/todo/database"
	. "github.com/Peyton232/todo/models"
)

type Handler struct {
	// add any needed fields here
	DB *DB
}

func NewHandler() (*Handler, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	return &Handler{
		DB: db,
	}, nil
}

// Here, we implement all of the handlers in the ServerInterface
func (p *Handler) FindUsers(ctx echo.Context, params FindUsersParams) error {
	// BUSINESS LOGIC
	return ctx.HTML(404, "FUCK")
	// return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) AddUser(ctx echo.Context) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) UpdateUser(ctx echo.Context) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) DeleteUser(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) FindUserByID(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) GetAchievements(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) AddAchievements(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) UpdateAchievements(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) DeleteAchievement(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}
func (p *Handler) GetAchievement(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) GetAlerts(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) AddAlert(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) UpdateAlert(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) DeleteAlert(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) GetAlert(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) FindBuckets(ctx echo.Context, id int32) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) AddBucket(ctx echo.Context, id int32) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) UpdateBucket(ctx echo.Context, id int32) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) DeleteBucket(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) FindBucketByName(ctx echo.Context, id int32, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) GetTransactions(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) InsertTransaction(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) UpdateTransaction(ctx echo.Context, id int64, name string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) DeeleteTransaction(ctx echo.Context, id int64, name string, timestamp string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) GetTransaction(ctx echo.Context, id int64, name string, timestamp string) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}
