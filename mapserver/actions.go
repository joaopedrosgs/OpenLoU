package mapserver

import (
	"context"
	"openlou/ent"
	"openlou/ent/city"
	"openlou/ent/user"
)

func (ms *mapserver) createCityAction(x int, y int, user *ent.User) (*ent.City, error) {
	return ms.GetClient().City.Create().SetY(y).SetX(x).SetOwnerID(user.ID).Save(context.Background())
}

func (ms *mapserver) getCitiesFromUserAction(u *ent.User) ([]*ent.City, error) {
	return ms.GetClient().City.Query().Where(city.HasOwnerWith(user.IDEQ(u.ID))).All(context.Background())
}

func (ms *mapserver) getCitiesAction(x int, y int, rang int) ([]*ent.City, error) {
	return ms.GetClient().City.Query().Where(city.And(city.XGTE(x-rang), city.XLTE(x+rang), city.YGTE(y-rang), city.YLTE(y+rang))).All(context.Background())

}
