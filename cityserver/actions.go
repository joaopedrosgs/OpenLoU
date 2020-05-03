package cityserver

import (
	"context"
	"github.com/joaopedrosgs/openlou/ent"
	"github.com/joaopedrosgs/openlou/ent/city"
	"github.com/joaopedrosgs/openlou/ent/construction"
	"github.com/joaopedrosgs/openlou/ent/queue"
	"github.com/joaopedrosgs/openlou/ent/user"
	"github.com/joaopedrosgs/openlou/session"
	"time"
)

func (cs *cityServer) upgradeConstructionAction(cityX int, cityY int, x int, y int) (*ent.Queue, error) {
	city, err := cs.GetClient().
		City.
		Query().
		Where(city.And(
			city.XEQ(cityX),
			city.YEQ(cityY))).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	construction, err := city.QueryConstructions().Where(construction.And(construction.XEQ(x), construction.YEQ(y))).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return cs.GetClient().Queue.Create().SetAction(1).SetConstruction(construction).SetCity(city).SetCompletion(city.QueueTime.Add(time.Second * 10)).Save(context.Background())

}

func (cs *cityServer) newConstructionAction(cityX int, cityY int, x int, y int, cType int, userSession *session.Session) (construction *ent.Construction, err error) {

	city, err := cs.GetClient().
		City.
		Query().
		Where(city.And(
			city.XEQ(cityX),
			city.YEQ(cityY),
			city.HasOwnerWith(
				user.NameEQ(userSession.User.Name)))).
		Only(context.Background())

	if err != nil {
		return
	}

	construction, err = cs.GetClient().Construction.Create().
		SetCity(city).
		SetLevel(0).
		SetX(x).SetY(y).
		SetType(cType).
		Save(context.Background())

	if err != nil {
		return
	}

	if city.QueueTime.Before(time.Now()) {
		city.Update().SetQueueTime(time.Now()).Save(context.Background())
	}

	_, err = cs.GetClient().Queue.Create().
		SetCity(city).
		SetConstruction(construction).
		SetAction(1).
		SetCompletion(city.QueueTime.Add(time.Second * 10)).
		Save(context.Background())

	return
}

func (cs *cityServer) getConstructionsAction(cityX int, cityY int, userSession *session.Session) (constructions []*ent.Construction, err error) {

	constructions, err = cs.GetClient().
		Construction.Query().
		Where(
			construction.HasCityWith(city.And(
				city.XEQ(cityX),
				city.YEQ(cityY),
				city.HasOwnerWith(
					user.NameEQ(userSession.User.Name))))).
		All(context.Background())
	return
}

func (cs *cityServer) getUpgradesAction(cityX int, cityY int, userSession *session.Session) (queueItems []*ent.Queue, err error) {
	queueItems, err = cs.GetClient().
		Queue.Query().
		Where(
			queue.HasCityWith(
				city.And(
					city.XEQ(cityX),
					city.YEQ(cityY),
					city.HasOwnerWith(
						user.NameEQ(userSession.User.Name))))).
		All(context.Background())
	return
}
