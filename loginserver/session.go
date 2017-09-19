package loginserver

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	LoggedIn   time.Time
	LastAction time.Time
	Login      string
	Ip         string
	Key        string
}
