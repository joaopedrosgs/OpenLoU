package communication

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/joaopedrosgs/OpenLoU/modules"
)

type Payload map[string]string
type Request struct {
	Key      string `valid:"required"`
	EndPoint int    `valid:"required"`
	Type     int    `valid:"required"`
	Data     Payload
}

func (r *Request) FieldsExist(fields ...string) error {
	for _, field := range fields {
		if _, ok := r.Data[field]; !ok {
			return fmt.Errorf("Empty %s value!", field)
		}

	}
	return nil
}
func (r *Request) ToAnswer() *Answer {
	return &Answer{"", r.Key, false, r.Type, "Bad Request", false}
}
func (r *Request) ToConstruction() (*models.Construction, error) {
	x, _ := strconv.Atoi(r.Data["X"])
	y, _ := strconv.Atoi(r.Data["Y"])
	cityX, _ := strconv.Atoi(r.Data["CityX"])
	cityY, _ := strconv.Atoi(r.Data["CityY"])
	t, _ := strconv.Atoi(r.Data["Type"])

	construction := &models.Construction{
		Tile:     models.Coord{x, y},
		CityTile: models.Coord{cityX, cityY},
		Type:     t}
	var err error
	if construction.Tile.X < 0 || construction.Tile.X > 21 {
		err = errors.New("Bad X value")
	} else if construction.Tile.Y < 0 || construction.Tile.Y > 19 {
		err = errors.New("Bar Y value")
	} else if _, ok := modules.RegisteredConstructions[construction.Type]; !ok {
		err = errors.New("Bad type valye")
	} else if construction.CityTile.Y < 0 || construction.CityTile.Y > 600 {
		err = errors.New("Bad City Y value")
	} else if construction.CityTile.X < 0 || construction.CityTile.X > 600 {
		err = errors.New("Bad City X value")
	}
	return construction, err
}

func (r *Request) ToCityCoord() (*models.Coord, error) {
	cityX, _ := strconv.Atoi(r.Data["CityX"])
	cityY, _ := strconv.Atoi(r.Data["CityY"])
	cityCoord := &models.Coord{
		X: cityX,
		Y: cityY}

	if cityCoord.Y < 0 || cityCoord.Y > 600 {
		return nil, errors.New("Bad City Y value")
	}
	if cityCoord.X < 0 || cityCoord.X > 600 {
		return nil, errors.New("Bad City X value")
	}
	return cityCoord, nil
}

func (r *Request) ToUpgrade() (*models.Upgrade, error) {
	x, err := strconv.Atoi(r.Data["X"])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(r.Data["Y"])
	if err != nil {
		return nil, err
	}
	cityX, err := strconv.Atoi(r.Data["CityX"])
	if err != nil {
		return nil, err
	}
	cityY, err := strconv.Atoi(r.Data["CityY"])
	if err != nil {
		return nil, err
	}
	return &models.Upgrade{
		Tile:     models.Coord{x, y},
		CityTile: models.Coord{cityX, cityY}}, nil
}
