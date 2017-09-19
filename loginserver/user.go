package loginserver

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id         bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	Login      string          `json:"login" bson:"login"`
	Pass       string          `json:"pass" bson:"pass"`
	Email      string          `json:"email" bson:"email"`
	Created    time.Time       `json:"created" bson:"created"`
	LastLogin  time.Time       `json:"last_login" bson:"last_login"`
	CitiesId   []bson.ObjectId `json:"cities_id" bson:"cities_id,omitempty"`
	SessionsId []bson.ObjectId `json:"sessions_id" bson:"sessions_id,omitempty"`
}
