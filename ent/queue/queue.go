// Code generated by ent, DO NOT EDIT.

package queue

const (
	// Label holds the string label denoting the queue type in the database.
	Label = "queue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCompletion holds the string denoting the completion field in the database.
	FieldCompletion = "completion"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// EdgeCity holds the string denoting the city edge name in mutations.
	EdgeCity = "city"
	// EdgeConstruction holds the string denoting the construction edge name in mutations.
	EdgeConstruction = "construction"
	// Table holds the table name of the queue in the database.
	Table = "queues"
	// CityTable is the table that holds the city relation/edge.
	CityTable = "queues"
	// CityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CityInverseTable = "cities"
	// CityColumn is the table column denoting the city relation/edge.
	CityColumn = "city_queue"
	// ConstructionTable is the table that holds the construction relation/edge.
	ConstructionTable = "queues"
	// ConstructionInverseTable is the table name for the Construction entity.
	// It exists in this package in order to avoid circular dependency with the "construction" package.
	ConstructionInverseTable = "constructions"
	// ConstructionColumn is the table column denoting the construction relation/edge.
	ConstructionColumn = "construction_queue"
)

// Columns holds all SQL columns for queue fields.
var Columns = []string{
	FieldID,
	FieldCompletion,
	FieldAction,
	FieldOrder,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "queues"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"city_queue",
	"construction_queue",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
