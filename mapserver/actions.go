package mapserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

func (ms *mapserver) createCityAction(x int, y int, user *models.User) (city models.City, err error) {
	city = models.City{
		X:         x,
		Y:         y,
		Type:      0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName:  null.NewString(user.Name, true),
		CityName:  "New City",
	}
	err = city.Insert(context.Background(), ms.GetConn(), boil.Infer())
	return
}

func (ms *mapserver) getCitiesFromUserAction(name string) (cities models.CitySlice, err error) {
	cities, err = models.Cities(qm.Where("user_name=?", name)).All(context.Background(), ms.GetConn())
	if err != nil {
		ms.LogContext.WithField("When", "getCitiesFromUserAction").Error(err.Error())
	}
	return
}

func (ms *mapserver) getCitiesAction(x int, y int, rang int) (cities models.CitySlice, err error) {
	cities, err = models.Cities(qm.Where(
		"x>? AND x<? AND y>? AND y<?", x-rang, x+rang, y-rang, y+rang)).All(context.Background(), ms.GetConn())
	if err != nil {
		ms.LogContext.WithField("When", "getCitiesAction").Error(err.Error())
	}
	return

}
