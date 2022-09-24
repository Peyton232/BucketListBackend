package database

import (
	"context"
	"time"

	model "github.com/Peyton232/todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// get
func (db *DB) GetAlerts(ID int64) []model.Alert {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, bson.M{}).Decode(&user)
	return *user.Alerts
}

// post
func (db *DB) AddAlert(ID int64, alert model.Alert) error {
	user := *db.FindUserByID(ID)
	newList := append(*user.Alerts, alert)
	user.Alerts = &newList
	return db.UpdateUser(user)
}

// put
func (db *DB) UpdateAlert(ID int64, alert model.Alert) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndUpdate(ctx, makeFilterAlert(ID, alert.Message), alert)
	return result.Err()
}

// get individualx
func (db *DB) FindAlertByName(ID int64, name string) *model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, makeFilterAlert(ID, name)).Decode(&user)
	return &user
}

// delete
func (db *DB) DeleteAlert(ID int64, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndDelete(ctx, makeFilterAlert(ID, name))
	return result.Err()
}

// TODO fix model and add name field
func makeFilterAlert(ID int64, name string) bson.D {
	return bson.D{
		{"$and",
			bson.A{
				bson.M{"id": ID},
				bson.M{"alerts.$.message": name},
			},
		},
	}
}
