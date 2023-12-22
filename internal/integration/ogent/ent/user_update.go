// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/tiagoposse/ogent/internal/integration/ogent/ent/pet"
	"github.com/tiagoposse/ogent/internal/integration/ogent/ent/predicate"
	"github.com/tiagoposse/ogent/internal/integration/ogent/ent/schema"
	"github.com/tiagoposse/ogent/internal/integration/ogent/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(u uint) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(u)
	return uu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge(u *uint) *UserUpdate {
	if u != nil {
		uu.SetAge(*u)
	}
	return uu
}

// AddAge adds u to the "age" field.
func (uu *UserUpdate) AddAge(u int) *UserUpdate {
	uu.mutation.AddAge(u)
	return uu
}

// SetHeight sets the "height" field.
func (uu *UserUpdate) SetHeight(u uint) *UserUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(u)
	return uu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHeight(u *uint) *UserUpdate {
	if u != nil {
		uu.SetHeight(*u)
	}
	return uu
}

// AddHeight adds u to the "height" field.
func (uu *UserUpdate) AddHeight(u int) *UserUpdate {
	uu.mutation.AddHeight(u)
	return uu
}

// ClearHeight clears the value of the "height" field.
func (uu *UserUpdate) ClearHeight() *UserUpdate {
	uu.mutation.ClearHeight()
	return uu
}

// SetFavoriteCatBreed sets the "favorite_cat_breed" field.
func (uu *UserUpdate) SetFavoriteCatBreed(ucb user.FavoriteCatBreed) *UserUpdate {
	uu.mutation.SetFavoriteCatBreed(ucb)
	return uu
}

// SetNillableFavoriteCatBreed sets the "favorite_cat_breed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFavoriteCatBreed(ucb *user.FavoriteCatBreed) *UserUpdate {
	if ucb != nil {
		uu.SetFavoriteCatBreed(*ucb)
	}
	return uu
}

// SetFavoriteDogBreed sets the "favorite_dog_breed" field.
func (uu *UserUpdate) SetFavoriteDogBreed(udb user.FavoriteDogBreed) *UserUpdate {
	uu.mutation.SetFavoriteDogBreed(udb)
	return uu
}

// SetNillableFavoriteDogBreed sets the "favorite_dog_breed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFavoriteDogBreed(udb *user.FavoriteDogBreed) *UserUpdate {
	if udb != nil {
		uu.SetFavoriteDogBreed(*udb)
	}
	return uu
}

// ClearFavoriteDogBreed clears the value of the "favorite_dog_breed" field.
func (uu *UserUpdate) ClearFavoriteDogBreed() *UserUpdate {
	uu.mutation.ClearFavoriteDogBreed()
	return uu
}

// SetFavoriteFishBreed sets the "favorite_fish_breed" field.
func (uu *UserUpdate) SetFavoriteFishBreed(sb schema.FishBreed) *UserUpdate {
	uu.mutation.SetFavoriteFishBreed(sb)
	return uu
}

// SetNillableFavoriteFishBreed sets the "favorite_fish_breed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFavoriteFishBreed(sb *schema.FishBreed) *UserUpdate {
	if sb != nil {
		uu.SetFavoriteFishBreed(*sb)
	}
	return uu
}

// ClearFavoriteFishBreed clears the value of the "favorite_fish_breed" field.
func (uu *UserUpdate) ClearFavoriteFishBreed() *UserUpdate {
	uu.mutation.ClearFavoriteFishBreed()
	return uu
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddPetIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPetIDs(ids...)
	return uu
}

// AddPets adds the "pets" edges to the Pet entity.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// AddAnimalsSavedIDs adds the "animals_saved" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddAnimalsSavedIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAnimalsSavedIDs(ids...)
	return uu
}

// AddAnimalsSaved adds the "animals_saved" edges to the Pet entity.
func (uu *UserUpdate) AddAnimalsSaved(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddAnimalsSavedIDs(ids...)
}

// SetBestFriendID sets the "best_friend" edge to the User entity by ID.
func (uu *UserUpdate) SetBestFriendID(id int) *UserUpdate {
	uu.mutation.SetBestFriendID(id)
	return uu
}

// SetNillableBestFriendID sets the "best_friend" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableBestFriendID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetBestFriendID(*id)
	}
	return uu
}

// SetBestFriend sets the "best_friend" edge to the User entity.
func (uu *UserUpdate) SetBestFriend(u *User) *UserUpdate {
	return uu.SetBestFriendID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (uu *UserUpdate) ClearPets() *UserUpdate {
	uu.mutation.ClearPets()
	return uu
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uu *UserUpdate) RemovePetIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePetIDs(ids...)
	return uu
}

// RemovePets removes "pets" edges to Pet entities.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
}

// ClearAnimalsSaved clears all "animals_saved" edges to the Pet entity.
func (uu *UserUpdate) ClearAnimalsSaved() *UserUpdate {
	uu.mutation.ClearAnimalsSaved()
	return uu
}

// RemoveAnimalsSavedIDs removes the "animals_saved" edge to Pet entities by IDs.
func (uu *UserUpdate) RemoveAnimalsSavedIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAnimalsSavedIDs(ids...)
	return uu
}

// RemoveAnimalsSaved removes "animals_saved" edges to Pet entities.
func (uu *UserUpdate) RemoveAnimalsSaved(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveAnimalsSavedIDs(ids...)
}

// ClearBestFriend clears the "best_friend" edge to the User entity.
func (uu *UserUpdate) ClearBestFriend() *UserUpdate {
	uu.mutation.ClearBestFriend()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.FavoriteCatBreed(); ok {
		if err := user.FavoriteCatBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_cat_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_cat_breed": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FavoriteDogBreed(); ok {
		if err := user.FavoriteDogBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_dog_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_dog_breed": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FavoriteFishBreed(); ok {
		if err := user.FavoriteFishBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_fish_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_fish_breed": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeUint, value)
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeUint, value)
	}
	if uu.mutation.HeightCleared() {
		_spec.ClearField(user.FieldHeight, field.TypeUint)
	}
	if value, ok := uu.mutation.FavoriteCatBreed(); ok {
		_spec.SetField(user.FieldFavoriteCatBreed, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.FavoriteDogBreed(); ok {
		_spec.SetField(user.FieldFavoriteDogBreed, field.TypeEnum, value)
	}
	if uu.mutation.FavoriteDogBreedCleared() {
		_spec.ClearField(user.FieldFavoriteDogBreed, field.TypeEnum)
	}
	if value, ok := uu.mutation.FavoriteFishBreed(); ok {
		_spec.SetField(user.FieldFavoriteFishBreed, field.TypeEnum, value)
	}
	if uu.mutation.FavoriteFishBreedCleared() {
		_spec.ClearField(user.FieldFavoriteFishBreed, field.TypeEnum)
	}
	if uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AnimalsSavedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AnimalsSavedTable,
			Columns: user.AnimalsSavedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAnimalsSavedIDs(); len(nodes) > 0 && !uu.mutation.AnimalsSavedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AnimalsSavedTable,
			Columns: user.AnimalsSavedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AnimalsSavedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AnimalsSavedTable,
			Columns: user.AnimalsSavedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.BestFriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(u uint) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(u)
	return uuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge(u *uint) *UserUpdateOne {
	if u != nil {
		uuo.SetAge(*u)
	}
	return uuo
}

// AddAge adds u to the "age" field.
func (uuo *UserUpdateOne) AddAge(u int) *UserUpdateOne {
	uuo.mutation.AddAge(u)
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UserUpdateOne) SetHeight(u uint) *UserUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(u)
	return uuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHeight(u *uint) *UserUpdateOne {
	if u != nil {
		uuo.SetHeight(*u)
	}
	return uuo
}

// AddHeight adds u to the "height" field.
func (uuo *UserUpdateOne) AddHeight(u int) *UserUpdateOne {
	uuo.mutation.AddHeight(u)
	return uuo
}

// ClearHeight clears the value of the "height" field.
func (uuo *UserUpdateOne) ClearHeight() *UserUpdateOne {
	uuo.mutation.ClearHeight()
	return uuo
}

// SetFavoriteCatBreed sets the "favorite_cat_breed" field.
func (uuo *UserUpdateOne) SetFavoriteCatBreed(ucb user.FavoriteCatBreed) *UserUpdateOne {
	uuo.mutation.SetFavoriteCatBreed(ucb)
	return uuo
}

// SetNillableFavoriteCatBreed sets the "favorite_cat_breed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFavoriteCatBreed(ucb *user.FavoriteCatBreed) *UserUpdateOne {
	if ucb != nil {
		uuo.SetFavoriteCatBreed(*ucb)
	}
	return uuo
}

// SetFavoriteDogBreed sets the "favorite_dog_breed" field.
func (uuo *UserUpdateOne) SetFavoriteDogBreed(udb user.FavoriteDogBreed) *UserUpdateOne {
	uuo.mutation.SetFavoriteDogBreed(udb)
	return uuo
}

// SetNillableFavoriteDogBreed sets the "favorite_dog_breed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFavoriteDogBreed(udb *user.FavoriteDogBreed) *UserUpdateOne {
	if udb != nil {
		uuo.SetFavoriteDogBreed(*udb)
	}
	return uuo
}

// ClearFavoriteDogBreed clears the value of the "favorite_dog_breed" field.
func (uuo *UserUpdateOne) ClearFavoriteDogBreed() *UserUpdateOne {
	uuo.mutation.ClearFavoriteDogBreed()
	return uuo
}

// SetFavoriteFishBreed sets the "favorite_fish_breed" field.
func (uuo *UserUpdateOne) SetFavoriteFishBreed(sb schema.FishBreed) *UserUpdateOne {
	uuo.mutation.SetFavoriteFishBreed(sb)
	return uuo
}

// SetNillableFavoriteFishBreed sets the "favorite_fish_breed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFavoriteFishBreed(sb *schema.FishBreed) *UserUpdateOne {
	if sb != nil {
		uuo.SetFavoriteFishBreed(*sb)
	}
	return uuo
}

// ClearFavoriteFishBreed clears the value of the "favorite_fish_breed" field.
func (uuo *UserUpdateOne) ClearFavoriteFishBreed() *UserUpdateOne {
	uuo.mutation.ClearFavoriteFishBreed()
	return uuo
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddPetIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPetIDs(ids...)
	return uuo
}

// AddPets adds the "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// AddAnimalsSavedIDs adds the "animals_saved" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddAnimalsSavedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAnimalsSavedIDs(ids...)
	return uuo
}

// AddAnimalsSaved adds the "animals_saved" edges to the Pet entity.
func (uuo *UserUpdateOne) AddAnimalsSaved(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddAnimalsSavedIDs(ids...)
}

// SetBestFriendID sets the "best_friend" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetBestFriendID(id int) *UserUpdateOne {
	uuo.mutation.SetBestFriendID(id)
	return uuo
}

// SetNillableBestFriendID sets the "best_friend" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBestFriendID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetBestFriendID(*id)
	}
	return uuo
}

// SetBestFriend sets the "best_friend" edge to the User entity.
func (uuo *UserUpdateOne) SetBestFriend(u *User) *UserUpdateOne {
	return uuo.SetBestFriendID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearPets() *UserUpdateOne {
	uuo.mutation.ClearPets()
	return uuo
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePetIDs(ids...)
	return uuo
}

// RemovePets removes "pets" edges to Pet entities.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
}

// ClearAnimalsSaved clears all "animals_saved" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearAnimalsSaved() *UserUpdateOne {
	uuo.mutation.ClearAnimalsSaved()
	return uuo
}

// RemoveAnimalsSavedIDs removes the "animals_saved" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemoveAnimalsSavedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAnimalsSavedIDs(ids...)
	return uuo
}

// RemoveAnimalsSaved removes "animals_saved" edges to Pet entities.
func (uuo *UserUpdateOne) RemoveAnimalsSaved(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveAnimalsSavedIDs(ids...)
}

// ClearBestFriend clears the "best_friend" edge to the User entity.
func (uuo *UserUpdateOne) ClearBestFriend() *UserUpdateOne {
	uuo.mutation.ClearBestFriend()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.FavoriteCatBreed(); ok {
		if err := user.FavoriteCatBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_cat_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_cat_breed": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FavoriteDogBreed(); ok {
		if err := user.FavoriteDogBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_dog_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_dog_breed": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FavoriteFishBreed(); ok {
		if err := user.FavoriteFishBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_fish_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_fish_breed": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeUint, value)
	}
	if uuo.mutation.HeightCleared() {
		_spec.ClearField(user.FieldHeight, field.TypeUint)
	}
	if value, ok := uuo.mutation.FavoriteCatBreed(); ok {
		_spec.SetField(user.FieldFavoriteCatBreed, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.FavoriteDogBreed(); ok {
		_spec.SetField(user.FieldFavoriteDogBreed, field.TypeEnum, value)
	}
	if uuo.mutation.FavoriteDogBreedCleared() {
		_spec.ClearField(user.FieldFavoriteDogBreed, field.TypeEnum)
	}
	if value, ok := uuo.mutation.FavoriteFishBreed(); ok {
		_spec.SetField(user.FieldFavoriteFishBreed, field.TypeEnum, value)
	}
	if uuo.mutation.FavoriteFishBreedCleared() {
		_spec.ClearField(user.FieldFavoriteFishBreed, field.TypeEnum)
	}
	if uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AnimalsSavedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AnimalsSavedTable,
			Columns: user.AnimalsSavedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAnimalsSavedIDs(); len(nodes) > 0 && !uuo.mutation.AnimalsSavedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AnimalsSavedTable,
			Columns: user.AnimalsSavedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AnimalsSavedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AnimalsSavedTable,
			Columns: user.AnimalsSavedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.BestFriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
