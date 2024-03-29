// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldGold holds the string denoting the gold field in the database.
	FieldGold = "gold"
	// FieldDiamonds holds the string denoting the diamonds field in the database.
	FieldDiamonds = "diamonds"
	// FieldDarkwood holds the string denoting the darkwood field in the database.
	FieldDarkwood = "darkwood"
	// FieldRunestone holds the string denoting the runestone field in the database.
	FieldRunestone = "runestone"
	// FieldVeritium holds the string denoting the veritium field in the database.
	FieldVeritium = "veritium"
	// FieldTrueseed holds the string denoting the trueseed field in the database.
	FieldTrueseed = "trueseed"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "rank"
	// FieldAllianceRank holds the string denoting the alliance_rank field in the database.
	FieldAllianceRank = "alliance_rank"
	// EdgeCities holds the string denoting the cities edge name in mutations.
	EdgeCities = "cities"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CitiesTable is the table that holds the cities relation/edge.
	CitiesTable = "cities"
	// CitiesInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CitiesInverseTable = "cities"
	// CitiesColumn is the table column denoting the cities relation/edge.
	CitiesColumn = "user_cities"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPasswordHash,
	FieldGold,
	FieldDiamonds,
	FieldDarkwood,
	FieldRunestone,
	FieldVeritium,
	FieldTrueseed,
	FieldRank,
	FieldAllianceRank,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// DefaultGold holds the default value on creation for the "gold" field.
	DefaultGold int
	// DefaultDiamonds holds the default value on creation for the "diamonds" field.
	DefaultDiamonds int
	// DefaultDarkwood holds the default value on creation for the "darkwood" field.
	DefaultDarkwood int
	// DefaultRunestone holds the default value on creation for the "runestone" field.
	DefaultRunestone int
	// DefaultVeritium holds the default value on creation for the "veritium" field.
	DefaultVeritium int
	// DefaultTrueseed holds the default value on creation for the "trueseed" field.
	DefaultTrueseed int
	// DefaultRank holds the default value on creation for the "rank" field.
	DefaultRank int
	// DefaultAllianceRank holds the default value on creation for the "alliance_rank" field.
	DefaultAllianceRank int
)
