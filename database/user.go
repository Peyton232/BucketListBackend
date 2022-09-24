package database

import (
	"context"
	"time"

	model "github.com/Peyton232/todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// get
func (db *DB) GetUsers() []model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	users := []model.User{}
	db.users.FindOne(ctx, bson.M{}).Decode(&users)
	return users
}

// post
func (db *DB) AddUser(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.users.InsertOne(ctx, user)
	return err
}

// put
func (db *DB) UpdateUser(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndUpdate(ctx, bson.M{"id": user.Id}, user)
	return result.Err()
}

// get individualx
func (db *DB) FindUserByID(ID int64) *model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, bson.M{"id": ID}).Decode(&user)
	return &user
}

// delete
func (db *DB) DeleteUser(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndDelete(ctx, bson.M{"id": ID})
	return result.Err()
}
