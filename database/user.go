package database

import (
	"context"
	"log"
	"time"

	model "github.com/Peyton232/todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// get
func (db *DB) GetUsers() []model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	users := []model.User{}
	res, _ := db.users.Find(ctx, bson.M{})
	for res.Next(ctx) {
		var user model.User
		err := res.Decode(&user)
		if err != nil {
			log.Print(err)
			log.Print("\nunable to read user model in database package\n")
			return nil
		}
		users = append(users, user)
	}
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
