package loginserver

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const MinPassSize = 7

type user struct {
	Id         bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	Login      string          `json:"login" bson:"login"`
	Pass       []byte          `json:"pass" bson:"pass"`
	Email      string          `json:"email" bson:"email"`
	Created    time.Time       `json:"created" bson:"created"`
	LastLogin  time.Time       `json:"last_login" bson:"last_login"`
	CitiesId   []bson.ObjectId `json:"cities_id" bson:"cities_id,omitempty"`
	SessionsId []bson.ObjectId `json:"sessions_id" bson:"sessions_id,omitempty"`
}

func (l *loginServer) NewUser(login, pass, email string) (*user, error) {

	if len(login) == 0 || len(email) == 0 {
		return nil, errors.New(EmptyFields)
	}
	if len(pass) < MinPassSize {
		return nil, errors.New(ShortPassword)

	}

	hashpass, err := bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {
		return nil, errors.New("There was a problem with the hash function")
	}
	user := user{Login: login, Pass: hashpass, Email: email, Created: time.Now()}

	db := l.Database.Clone()
	defer db.Close()
	err = storeUser(user, db)
	if err != nil {
		return nil, err
	}
	err = db.DB(DBName).C(DBUsers).Find(bson.M{"login": login}).One(&user)

	return nil, nil
}

func (l *loginServer) LoadUserByLogin(login string) (*user, error) {
	if len(login) == 0 {
		return nil, errors.New(EmptyFields)
	}
	db := l.Database.Clone()
	defer db.Close()
	users := db.DB(DBName).C(DBUsers)
	u := user{}
	err := users.Find(bson.M{"login": login}).One(&u)
	return &u, err
}

func (l *loginServer) SaveUserChanges(u *user) error {
	db := l.Database.Clone()
	defer db.Close()
	users := db.DB(DBName).C(DBUsers)
	return users.UpdateId(u.Id, u)
}

func (l *loginServer) UserExists(login, email string) error {
	db := l.Database.Clone()
	defer db.Close()
	users := db.DB(DBName).C(DBUsers)
	account, err := users.Find(bson.M{"$or": []bson.M{bson.M{"login": login}, bson.M{"email": email}}}).Count()
	if err != nil {
		return errors.New(DBError)
	}
	if account == 0 {
		return errors.New(AccountInexistent)
	}
	return nil

}
func (l *loginServer) DeleteUserByLogin(login string) error {
	db := l.Database.Clone()
	defer db.Close()
	users := db.DB(DBName).C(DBUsers)
	db.DB(DBName).C(DBSessions).RemoveAll(bson.M{"login": login})
	return users.Remove(bson.M{"login": login})
}

func storeUser(u user, db *mgo.Session) error {
	users := db.DB(DBName).C(DBUsers)

	account, err := users.Find(bson.M{"$or": []bson.M{bson.M{"login": u.Login}, bson.M{"email": u.Email}}}).Count()

	if err != nil {
		return errors.New(DBError)
	}
	if account > 0 {
		return errors.New(AccountExists)
	}
	users.Insert(u)
	return nil
}
