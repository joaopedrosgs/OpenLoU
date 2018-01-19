package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/entities"
	log "github.com/sirupsen/logrus"
	"time"
)

var db *gorm.DB
var context = log.WithField("Entity", "Database")

func InitDB() {
	err := errors.New("")
	db, err = gorm.Open("postgres", configuration.GetConnectionString())
	if err != nil {
		context.Error(err.Error())
	}
	db.DropTableIfExists(&entities.User{}, &entities.City{}, &entities.Construction{}, &entities.Upgrade{}, &entities.Dungeon{}, &entities.Continent{}, &entities.WorldResource{})
	db.AutoMigrate(&entities.User{}, &entities.City{}, &entities.Construction{}, &entities.Upgrade{}, &entities.Dungeon{}, &entities.Continent{}, &entities.WorldResource{})

}

func CreateUser(login, password_hash, email string) (*entities.User, error) {
	if db == nil {
		InitDB()
	}
	user := &entities.User{Name: login, Email: email, PasswordHash: password_hash, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.NewRecord(user)
	db.Create(user)
	if user.ID > 0 {
		return user, nil
	}
	CreateCity(user.ID)
	return nil, errors.New("Failed to create account!")
}

func GetUser(email string) (*entities.User, error) {
	if db == nil {
		InitDB()
	}
	user := &entities.User{}
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		return user, nil
	}
	return nil, errors.New("Account not found account!")
}

func GetCitiesInRange(x, y, radius, continent int) *[]entities.City {
	cities := []entities.City{}
	db.Where("x BETWEEN ? AND ? AND y BETWEEN ? AND ? AND continent_id = ?", x-radius, x+radius, y-radius, y+radius, continent).Find(&cities)
	return &cities
}

func CreateCity(userID uint) {
	continent := entities.Continent{IsActive: true, Size: configuration.GetSingleton().Parameters.General.ContinentSize, CitiesLimit: 250, X: 0, Y: 0}
	db.Where("Is_Active = ? AND Cities_Limit <= ?", true, 250).FirstOrCreate(&continent)
	city := entities.City{}
	city.ContinentID = continent.ID
	city.UserID = userID
	townHall := entities.Construction{CityID: city.ID, X: 10, Y: 10}
	city.Constructions = append(city.Constructions, townHall)
	db.Create(&city)
	db.NewRecord(city)

}
