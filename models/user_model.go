package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var UserModel *mongo.Collection

type UserSchema struct {
	*BaseSchema `bson:",inline"`
	Username    string `bson:"username"`
	Password    string `bson:"password"`
}
