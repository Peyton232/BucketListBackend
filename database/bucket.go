package database

import (
	"context"
	"time"

	model "github.com/Peyton232/todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// get
func (db *DB) GetBuckets(ID int64) []model.Bucket {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, bson.M{}).Decode(&user)
	return *user.Buckets
}

// post
func (db *DB) AddBucket(ID int64, bucket model.Bucket) error {
	user := *db.FindUserByID(ID)
	newList := append(*user.Buckets, bucket)
	user.Buckets = &newList
	return db.UpdateUser(user)
}

// put
func (db *DB) UpdateBucket(ID int64, bucket model.Bucket) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndUpdate(ctx, makeFilter(ID, bucket.Name), bucket)
	return result.Err()
}

// get individualx
func (db *DB) FindBucketByName(ID int64, name string) *model.Bucket {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	bucket := model.Bucket{}
	db.users.FindOne(ctx, makeFilter(ID, name)).Decode(&bucket)
	return &bucket
}

// delete
func (db *DB) DeleteBucket(ID int64, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndDelete(ctx, makeFilter(ID, name))
	return result.Err()
}

func makeFilter(ID int64, name string) bson.D {
	return bson.D{
		{"$and",
			bson.A{
				bson.M{"id": ID},
				bson.M{"buckets.$.name": name},
			},
		},
	}
}
