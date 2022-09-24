package database

import (
	"context"
	"time"

	model "github.com/Peyton232/todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// get
func (db *DB) GetTransactions(ID int64, name string) []model.Transaction {
	bucket := db.FindBucketByName(ID, name)
	return *bucket.History
}

// post
// TODO this will not work
func (db *DB) AddTransaction(ID int64, name string, transaction model.Transaction) error {
	history := db.GetTransactions(ID, name)
	newList := append(history, transaction)
	user := db.FindUserByID(ID)
	for _, buck := range *user.Buckets {
		if buck.Name == name {
			buck.History = &newList
		}
	}
	return db.UpdateUser(*user)
}

// put
func (db *DB) UpdateTransaction(ID int64, bucket model.Bucket) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndUpdate(ctx, makeFilterTransaction(ID, bucket.Name), bucket)
	return result.Err()
}

// get individualx
func (db *DB) FindTransactionByName(ID int64, name string) *model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, makeFilterTransaction(ID, name)).Decode(&user)
	return &user
}

// delete
func (db *DB) DeleteTransaction(ID int64, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndDelete(ctx, makeFilterTransaction(ID, name))
	return result.Err()
}

// TODO this does nto work
func makeFilterTransaction(ID int64, name string) bson.D {
	return bson.D{
		{"$and",
			bson.A{
				bson.M{"id": ID},
				bson.M{"buckets.$.name": name},
			},
		},
	}
}
