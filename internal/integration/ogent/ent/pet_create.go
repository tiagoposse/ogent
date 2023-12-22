// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ariga.io/ogent/internal/integration/ogent/ent/category"
	"ariga.io/ogent/internal/integration/ogent/ent/pet"
	"ariga.io/ogent/internal/integration/ogent/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetWeight sets the "weight" field.
func (pc *PetCreate) SetWeight(i int) *PetCreate {
	pc.mutation.SetWeight(i)
	return pc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (pc *PetCreate) SetNillableWeight(i *int) *PetCreate {
	if i != nil {
		pc.SetWeight(*i)
	}
	return pc
}

// SetBirthday sets the "birthday" field.
func (pc *PetCreate) SetBirthday(t time.Time) *PetCreate {
	pc.mutation.SetBirthday(t)
	return pc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (pc *PetCreate) SetNillableBirthday(t *time.Time) *PetCreate {
	if t != nil {
		pc.SetBirthday(*t)
	}
	return pc
}

// SetTagID sets the "tag_id" field.
func (pc *PetCreate) SetTagID(b []byte) *PetCreate {
	pc.mutation.SetTagID(b)
	return pc
}

// SetHeight sets the "height" field.
func (pc *PetCreate) SetHeight(i int) *PetCreate {
	pc.mutation.SetHeight(i)
	return pc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (pc *PetCreate) SetNillableHeight(i *int) *PetCreate {
	if i != nil {
		pc.SetHeight(*i)
	}
	return pc
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (pc *PetCreate) AddCategoryIDs(ids ...int) *PetCreate {
	pc.mutation.AddCategoryIDs(ids...)
	return pc
}

// AddCategories adds the "categories" edges to the Category entity.
func (pc *PetCreate) AddCategories(c ...*Category) *PetCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCategoryIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PetCreate) SetOwnerID(id int) *PetCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// AddRescuerIDs adds the "rescuer" edge to the User entity by IDs.
func (pc *PetCreate) AddRescuerIDs(ids ...int) *PetCreate {
	pc.mutation.AddRescuerIDs(ids...)
	return pc
}

// AddRescuer adds the "rescuer" edges to the User entity.
func (pc *PetCreate) AddRescuer(u ...*User) *PetCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddRescuerIDs(ids...)
}

// AddFriendIDs adds the "friends" edge to the Pet entity by IDs.
func (pc *PetCreate) AddFriendIDs(ids ...int) *PetCreate {
	pc.mutation.AddFriendIDs(ids...)
	return pc
}

// AddFriends adds the "friends" edges to the Pet entity.
func (pc *PetCreate) AddFriends(p ...*Pet) *PetCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddFriendIDs(ids...)
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PetCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PetCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Pet.owner"`)}
	}
	return nil
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pet.Table, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Weight(); ok {
		_spec.SetField(pet.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if value, ok := pc.mutation.Birthday(); ok {
		_spec.SetField(pet.FieldBirthday, field.TypeTime, value)
		_node.Birthday = value
	}
	if value, ok := pc.mutation.TagID(); ok {
		_spec.SetField(pet.FieldTagID, field.TypeBytes, value)
		_node.TagID = value
	}
	if value, ok := pc.mutation.Height(); ok {
		_spec.SetField(pet.FieldHeight, field.TypeInt, value)
		_node.Height = &value
	}
	if nodes := pc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   pet.CategoriesTable,
			Columns: pet.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_pets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RescuerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   pet.RescuerTable,
			Columns: pet.RescuerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   pet.FriendsTable,
			Columns: pet.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	err      error
	builders []*PetCreate
}

// Save creates the Pet entities in the database.
func (pcb *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PetCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
