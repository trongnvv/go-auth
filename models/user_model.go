package models

type UserSchema struct {
	*BaseSchema `bson:",inline"`
	Username    string `bson:"username"`
	Password    string `bson:"password"`
}
