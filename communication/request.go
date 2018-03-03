package communication

import (
	"errors"
	"fmt"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

type Request struct {
	Key      string
	Type     int
	Data     map[string]uint
	isSystem bool
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
func (r *Request) ToConstruction() (*entities.Construction, error) {
	construction := &entities.Construction{X: r.Data["X"], Y: r.Data["Y"], CityX: r.Data["CityX"], CityY: r.Data["CityY"], Type: r.Data["Type"]}
	var err error
	if construction.X < 0 || construction.X > 21 {
		err = errors.New("Bad X value")
	} else if construction.Y < 0 || construction.Y > 19 {
		err = errors.New("Bar Y value")
	} else if _, ok := entities.RegisteredConstructions[construction.Type]; !ok {
		err = errors.New("Bad type valye")
	} else if construction.CityY < 0 || construction.CityY > 600 {
		err = errors.New("Bad City Y value")
	} else if construction.CityX < 0 || construction.CityX > 600 {
		err = errors.New("Bad City X value")
	}
	return construction, err
}
