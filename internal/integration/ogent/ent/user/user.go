// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/tiagoposse/ogent/internal/integration/ogent/ent/schema"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldFavoriteCatBreed holds the string denoting the favorite_cat_breed field in the database.
	FieldFavoriteCatBreed = "favorite_cat_breed"
	// FieldFavoriteColor holds the string denoting the favorite_color field in the database.
	FieldFavoriteColor = "favorite_color"
	// FieldFavoriteDogBreed holds the string denoting the favorite_dog_breed field in the database.
	FieldFavoriteDogBreed = "favorite_dog_breed"
	// FieldFavoriteFishBreed holds the string denoting the favorite_fish_breed field in the database.
	FieldFavoriteFishBreed = "favorite_fish_breed"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeAnimalsSaved holds the string denoting the animals_saved edge name in mutations.
	EdgeAnimalsSaved = "animals_saved"
	// EdgeBestFriend holds the string denoting the best_friend edge name in mutations.
	EdgeBestFriend = "best_friend"
	// EdgeFavoriteHat holds the string denoting the favorite_hat edge name in mutations.
	EdgeFavoriteHat = "favorite_hat"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "user_pets"
	// AnimalsSavedTable is the table that holds the animals_saved relation/edge. The primary key declared below.
	AnimalsSavedTable = "user_animals_saved"
	// AnimalsSavedInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	AnimalsSavedInverseTable = "pets"
	// BestFriendTable is the table that holds the best_friend relation/edge.
	BestFriendTable = "users"
	// BestFriendColumn is the table column denoting the best_friend relation/edge.
	BestFriendColumn = "user_best_friend"
	// FavoriteHatTable is the table that holds the favorite_hat relation/edge.
	FavoriteHatTable = "hats"
	// FavoriteHatInverseTable is the table name for the Hat entity.
	// It exists in this package in order to avoid circular dependency with the "hat" package.
	FavoriteHatInverseTable = "hats"
	// FavoriteHatColumn is the table column denoting the favorite_hat relation/edge.
	FavoriteHatColumn = "user_favorite_hat"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
	FieldHeight,
	FieldFavoriteCatBreed,
	FieldFavoriteColor,
	FieldFavoriteDogBreed,
	FieldFavoriteFishBreed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_best_friend",
}

var (
	// AnimalsSavedPrimaryKey and AnimalsSavedColumn2 are the table columns denoting the
	// primary key for the animals_saved relation (M2M).
	AnimalsSavedPrimaryKey = []string{"user_id", "pet_id"}
)

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

// FavoriteCatBreed defines the type for the "favorite_cat_breed" enum field.
type FavoriteCatBreed string

// FavoriteCatBreed values.
const (
	FavoriteCatBreedSiamese FavoriteCatBreed = "siamese"
	FavoriteCatBreedBengal  FavoriteCatBreed = "bengal"
	FavoriteCatBreedLion    FavoriteCatBreed = "lion"
	FavoriteCatBreedTiger   FavoriteCatBreed = "tiger"
	FavoriteCatBreedLeopard FavoriteCatBreed = "leopard"
	FavoriteCatBreedOther   FavoriteCatBreed = "other"
)

func (fcb FavoriteCatBreed) String() string {
	return string(fcb)
}

// FavoriteCatBreedValidator is a validator for the "favorite_cat_breed" field enum values. It is called by the builders before save.
func FavoriteCatBreedValidator(fcb FavoriteCatBreed) error {
	switch fcb {
	case FavoriteCatBreedSiamese, FavoriteCatBreedBengal, FavoriteCatBreedLion, FavoriteCatBreedTiger, FavoriteCatBreedLeopard, FavoriteCatBreedOther:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for favorite_cat_breed field: %q", fcb)
	}
}

// FavoriteColor defines the type for the "favorite_color" enum field.
type FavoriteColor string

// FavoriteColorRed is the default value of the FavoriteColor enum.
const DefaultFavoriteColor = FavoriteColorRed

// FavoriteColor values.
const (
	FavoriteColorRed   FavoriteColor = "red"
	FavoriteColorGreen FavoriteColor = "green"
	FavoriteColorBlue  FavoriteColor = "blue"
)

func (fc FavoriteColor) String() string {
	return string(fc)
}

// FavoriteColorValidator is a validator for the "favorite_color" field enum values. It is called by the builders before save.
func FavoriteColorValidator(fc FavoriteColor) error {
	switch fc {
	case FavoriteColorRed, FavoriteColorGreen, FavoriteColorBlue:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for favorite_color field: %q", fc)
	}
}

// FavoriteDogBreed defines the type for the "favorite_dog_breed" enum field.
type FavoriteDogBreed string

// FavoriteDogBreed values.
const (
	FavoriteDogBreedKuro FavoriteDogBreed = "Kuro"
)

func (fdb FavoriteDogBreed) String() string {
	return string(fdb)
}

// FavoriteDogBreedValidator is a validator for the "favorite_dog_breed" field enum values. It is called by the builders before save.
func FavoriteDogBreedValidator(fdb FavoriteDogBreed) error {
	switch fdb {
	case FavoriteDogBreedKuro:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for favorite_dog_breed field: %q", fdb)
	}
}

// FavoriteFishBreedValidator is a validator for the "favorite_fish_breed" field enum values. It is called by the builders before save.
func FavoriteFishBreedValidator(ffb schema.FishBreed) error {
	switch ffb {
	case "gold", "koi", "shark":
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for favorite_fish_breed field: %q", ffb)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByFavoriteCatBreed orders the results by the favorite_cat_breed field.
func ByFavoriteCatBreed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavoriteCatBreed, opts...).ToFunc()
}

// ByFavoriteColor orders the results by the favorite_color field.
func ByFavoriteColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavoriteColor, opts...).ToFunc()
}

// ByFavoriteDogBreed orders the results by the favorite_dog_breed field.
func ByFavoriteDogBreed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavoriteDogBreed, opts...).ToFunc()
}

// ByFavoriteFishBreed orders the results by the favorite_fish_breed field.
func ByFavoriteFishBreed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavoriteFishBreed, opts...).ToFunc()
}

// ByPetsCount orders the results by pets count.
func ByPetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPetsStep(), opts...)
	}
}

// ByPets orders the results by pets terms.
func ByPets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAnimalsSavedCount orders the results by animals_saved count.
func ByAnimalsSavedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnimalsSavedStep(), opts...)
	}
}

// ByAnimalsSaved orders the results by animals_saved terms.
func ByAnimalsSaved(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnimalsSavedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBestFriendField orders the results by best_friend field.
func ByBestFriendField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBestFriendStep(), sql.OrderByField(field, opts...))
	}
}

// ByFavoriteHatField orders the results by favorite_hat field.
func ByFavoriteHatField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavoriteHatStep(), sql.OrderByField(field, opts...))
	}
}
func newPetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PetsTable, PetsColumn),
	)
}
func newAnimalsSavedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnimalsSavedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AnimalsSavedTable, AnimalsSavedPrimaryKey...),
	)
}
func newBestFriendStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, BestFriendTable, BestFriendColumn),
	)
}
func newFavoriteHatStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FavoriteHatInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FavoriteHatTable, FavoriteHatColumn),
	)
}
