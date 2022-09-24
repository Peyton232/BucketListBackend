package api

import (
	"net/http"

	. "github.com/Peyton232/todo/database"
	. "github.com/Peyton232/todo/models"
)

type HandlerChi struct {
	ServerInterfaceWrapper
	DB *DB
}

func NewHandler() (*HandlerChi, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	return &HandlerChi{
		DB: db,
	}, nil
}

// Here, we implement all of the handlers in the ServerInterface
func (p HandlerChi) FindUsers(w http.ResponseWriter, r *http.Request, params FindUsersParams) {
	// BUSINESS LOGIC
}

func (p HandlerChi) AddUser(w http.ResponseWriter, r *http.Request) {
	// BUSINESS LOGIC
}

func (p HandlerChi) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// BUSINESS LOGIC
}

func (p HandlerChi) DeleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC
}

func (p HandlerChi) FindUserByID(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) GetAchievements(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) AddAchievements(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) UpdateAchievements(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) DeleteAchievement(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}
func (p HandlerChi) GetAchievement(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) GetAlerts(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) AddAlert(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) UpdateAlert(w http.ResponseWriter, r *http.Request, id int64) {
	// BUSINESS LOGIC

}

func (p HandlerChi) DeleteAlert(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) GetAlert(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) FindBuckets(w http.ResponseWriter, r *http.Request, id int32) {
	// BUSINESS LOGIC

}

func (p HandlerChi) AddBucket(w http.ResponseWriter, r *http.Request, id int32) {
	// BUSINESS LOGIC

}

func (p HandlerChi) UpdateBucket(w http.ResponseWriter, r *http.Request, id int32) {
	// BUSINESS LOGIC

}

func (p HandlerChi) DeleteBucket(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) FindBucketByName(w http.ResponseWriter, r *http.Request, id int32, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) GetTransactions(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) InsertTransaction(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) UpdateTransaction(w http.ResponseWriter, r *http.Request, id int64, name string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) DeeleteTransaction(w http.ResponseWriter, r *http.Request, id int64, name string, timestamp string) {
	// BUSINESS LOGIC

}

func (p HandlerChi) GetTransaction(w http.ResponseWriter, r *http.Request, id int64, name string, timestamp string) {
	// BUSINESS LOGIC

}
