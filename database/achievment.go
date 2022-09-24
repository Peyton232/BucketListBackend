package database

import (
	"context"
	"time"

	model "github.com/Peyton232/todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// get
func (db *DB) GetAchievment(ID int64) []model.Achievement {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, bson.M{}).Decode(&user)
	return *user.Achievements
}

// post
func (db *DB) AddAchievment(ID int64, achievment model.Achievement) error {
	user := *db.FindUserByID(ID)
	newList := append(*user.Achievements, achievment)
	user.Achievements = &newList
	return db.UpdateUser(user)
}

// put
func (db *DB) UpdateAchievment(ID int64, achievment model.Achievement) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndUpdate(ctx, makeFilterAchievment(ID, achievment.Name), achievment)
	return result.Err()
}

// get individual
func (db *DB) FindAchievmentByName(ID int64, name string) *model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	db.users.FindOne(ctx, makeFilterAchievment(ID, name)).Decode(&user)
	return &user
}

// delete
func (db *DB) DeleteAchievment(ID int64, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := db.users.FindOneAndDelete(ctx, makeFilterAchievment(ID, name))
	return result.Err()
}

func makeFilterAchievment(ID int64, name string) bson.D {
	return bson.D{
		{"$and",
			bson.A{
				bson.M{"id": ID},
				bson.M{"achievments.$.name": name},
			},
		},
	}
}
