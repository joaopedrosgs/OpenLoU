// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"openlou/ent/city"
	"openlou/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// X holds the value of the "x" field.
	X int `json:"x,omitempty"`
	// Y holds the value of the "y" field.
	Y int `json:"y,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Points holds the value of the "points" field.
	Points int `json:"points,omitempty"`
	// WoodProduction holds the value of the "wood_production" field.
	WoodProduction int `json:"wood_production,omitempty"`
	// StoneProduction holds the value of the "stone_production" field.
	StoneProduction int `json:"stone_production,omitempty"`
	// IronProduction holds the value of the "iron_production" field.
	IronProduction int `json:"iron_production,omitempty"`
	// FoodProduction holds the value of the "food_production" field.
	FoodProduction int `json:"food_production,omitempty"`
	// WoodStored holds the value of the "wood_stored" field.
	WoodStored int `json:"wood_stored,omitempty"`
	// StoneStored holds the value of the "stone_stored" field.
	StoneStored int `json:"stone_stored,omitempty"`
	// IronStored holds the value of the "iron_stored" field.
	IronStored int `json:"iron_stored,omitempty"`
	// FoodStored holds the value of the "food_stored" field.
	FoodStored int `json:"food_stored,omitempty"`
	// WoodLimit holds the value of the "wood_limit" field.
	WoodLimit int `json:"wood_limit,omitempty"`
	// StoneLimit holds the value of the "stone_limit" field.
	StoneLimit int `json:"stone_limit,omitempty"`
	// IronLimit holds the value of the "iron_limit" field.
	IronLimit int `json:"iron_limit,omitempty"`
	// FoodLimit holds the value of the "food_limit" field.
	FoodLimit int `json:"food_limit,omitempty"`
	// QueueTime holds the value of the "queue_time" field.
	QueueTime time.Time `json:"queue_time,omitempty"`
	// ConstructionSpeed holds the value of the "construction_speed" field.
	ConstructionSpeed int `json:"construction_speed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CityQuery when eager-loading is set.
	Edges       CityEdges `json:"edges"`
	user_cities *int
}

// CityEdges holds the relations/edges for other nodes in the graph.
type CityEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Constructions holds the value of the constructions edge.
	Constructions []*Construction `json:"constructions,omitempty"`
	// Queue holds the value of the queue edge.
	Queue []*Queue `json:"queue,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CityEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ConstructionsOrErr returns the Constructions value or an error if the edge
// was not loaded in eager-loading.
func (e CityEdges) ConstructionsOrErr() ([]*Construction, error) {
	if e.loadedTypes[1] {
		return e.Constructions, nil
	}
	return nil, &NotLoadedError{edge: "constructions"}
}

// QueueOrErr returns the Queue value or an error if the edge
// was not loaded in eager-loading.
func (e CityEdges) QueueOrErr() ([]*Queue, error) {
	if e.loadedTypes[2] {
		return e.Queue, nil
	}
	return nil, &NotLoadedError{edge: "queue"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*City) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case city.FieldID, city.FieldX, city.FieldY, city.FieldPoints, city.FieldWoodProduction, city.FieldStoneProduction, city.FieldIronProduction, city.FieldFoodProduction, city.FieldWoodStored, city.FieldStoneStored, city.FieldIronStored, city.FieldFoodStored, city.FieldWoodLimit, city.FieldStoneLimit, city.FieldIronLimit, city.FieldFoodLimit, city.FieldConstructionSpeed:
			values[i] = new(sql.NullInt64)
		case city.FieldName:
			values[i] = new(sql.NullString)
		case city.FieldQueueTime:
			values[i] = new(sql.NullTime)
		case city.ForeignKeys[0]: // user_cities
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type City", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the City fields.
func (c *City) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case city.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case city.FieldX:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field x", values[i])
			} else if value.Valid {
				c.X = int(value.Int64)
			}
		case city.FieldY:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field y", values[i])
			} else if value.Valid {
				c.Y = int(value.Int64)
			}
		case city.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case city.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				c.Points = int(value.Int64)
			}
		case city.FieldWoodProduction:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wood_production", values[i])
			} else if value.Valid {
				c.WoodProduction = int(value.Int64)
			}
		case city.FieldStoneProduction:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stone_production", values[i])
			} else if value.Valid {
				c.StoneProduction = int(value.Int64)
			}
		case city.FieldIronProduction:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field iron_production", values[i])
			} else if value.Valid {
				c.IronProduction = int(value.Int64)
			}
		case city.FieldFoodProduction:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field food_production", values[i])
			} else if value.Valid {
				c.FoodProduction = int(value.Int64)
			}
		case city.FieldWoodStored:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wood_stored", values[i])
			} else if value.Valid {
				c.WoodStored = int(value.Int64)
			}
		case city.FieldStoneStored:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stone_stored", values[i])
			} else if value.Valid {
				c.StoneStored = int(value.Int64)
			}
		case city.FieldIronStored:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field iron_stored", values[i])
			} else if value.Valid {
				c.IronStored = int(value.Int64)
			}
		case city.FieldFoodStored:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field food_stored", values[i])
			} else if value.Valid {
				c.FoodStored = int(value.Int64)
			}
		case city.FieldWoodLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wood_limit", values[i])
			} else if value.Valid {
				c.WoodLimit = int(value.Int64)
			}
		case city.FieldStoneLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stone_limit", values[i])
			} else if value.Valid {
				c.StoneLimit = int(value.Int64)
			}
		case city.FieldIronLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field iron_limit", values[i])
			} else if value.Valid {
				c.IronLimit = int(value.Int64)
			}
		case city.FieldFoodLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field food_limit", values[i])
			} else if value.Valid {
				c.FoodLimit = int(value.Int64)
			}
		case city.FieldQueueTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field queue_time", values[i])
			} else if value.Valid {
				c.QueueTime = value.Time
			}
		case city.FieldConstructionSpeed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field construction_speed", values[i])
			} else if value.Valid {
				c.ConstructionSpeed = int(value.Int64)
			}
		case city.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_cities", value)
			} else if value.Valid {
				c.user_cities = new(int)
				*c.user_cities = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the City entity.
func (c *City) QueryOwner() *UserQuery {
	return (&CityClient{config: c.config}).QueryOwner(c)
}

// QueryConstructions queries the "constructions" edge of the City entity.
func (c *City) QueryConstructions() *ConstructionQuery {
	return (&CityClient{config: c.config}).QueryConstructions(c)
}

// QueryQueue queries the "queue" edge of the City entity.
func (c *City) QueryQueue() *QueueQuery {
	return (&CityClient{config: c.config}).QueryQueue(c)
}

// Update returns a builder for updating this City.
// Note that you need to call City.Unwrap() before calling this method if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *City) Update() *CityUpdateOne {
	return (&CityClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the City entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *City) Unwrap() *City {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *City) String() string {
	var builder strings.Builder
	builder.WriteString("City(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("x=")
	builder.WriteString(fmt.Sprintf("%v", c.X))
	builder.WriteString(", ")
	builder.WriteString("y=")
	builder.WriteString(fmt.Sprintf("%v", c.Y))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("points=")
	builder.WriteString(fmt.Sprintf("%v", c.Points))
	builder.WriteString(", ")
	builder.WriteString("wood_production=")
	builder.WriteString(fmt.Sprintf("%v", c.WoodProduction))
	builder.WriteString(", ")
	builder.WriteString("stone_production=")
	builder.WriteString(fmt.Sprintf("%v", c.StoneProduction))
	builder.WriteString(", ")
	builder.WriteString("iron_production=")
	builder.WriteString(fmt.Sprintf("%v", c.IronProduction))
	builder.WriteString(", ")
	builder.WriteString("food_production=")
	builder.WriteString(fmt.Sprintf("%v", c.FoodProduction))
	builder.WriteString(", ")
	builder.WriteString("wood_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.WoodStored))
	builder.WriteString(", ")
	builder.WriteString("stone_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.StoneStored))
	builder.WriteString(", ")
	builder.WriteString("iron_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.IronStored))
	builder.WriteString(", ")
	builder.WriteString("food_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.FoodStored))
	builder.WriteString(", ")
	builder.WriteString("wood_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.WoodLimit))
	builder.WriteString(", ")
	builder.WriteString("stone_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.StoneLimit))
	builder.WriteString(", ")
	builder.WriteString("iron_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.IronLimit))
	builder.WriteString(", ")
	builder.WriteString("food_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.FoodLimit))
	builder.WriteString(", ")
	builder.WriteString("queue_time=")
	builder.WriteString(c.QueueTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("construction_speed=")
	builder.WriteString(fmt.Sprintf("%v", c.ConstructionSpeed))
	builder.WriteByte(')')
	return builder.String()
}

// Cities is a parsable slice of City.
type Cities []*City

func (c Cities) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
