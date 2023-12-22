// Code generated by ent, DO NOT EDIT.

package category

import (
	"github.com/tiagoposse/ogent/internal/integration/ogent/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// Readonly applies equality check predicate on the "readonly" field. It's identical to ReadonlyEQ.
func Readonly(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldReadonly, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldName, v))
}

// ReadonlyEQ applies the EQ predicate on the "readonly" field.
func ReadonlyEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldReadonly, v))
}

// ReadonlyNEQ applies the NEQ predicate on the "readonly" field.
func ReadonlyNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldReadonly, v))
}

// ReadonlyIn applies the In predicate on the "readonly" field.
func ReadonlyIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldReadonly, vs...))
}

// ReadonlyNotIn applies the NotIn predicate on the "readonly" field.
func ReadonlyNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldReadonly, vs...))
}

// ReadonlyGT applies the GT predicate on the "readonly" field.
func ReadonlyGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldReadonly, v))
}

// ReadonlyGTE applies the GTE predicate on the "readonly" field.
func ReadonlyGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldReadonly, v))
}

// ReadonlyLT applies the LT predicate on the "readonly" field.
func ReadonlyLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldReadonly, v))
}

// ReadonlyLTE applies the LTE predicate on the "readonly" field.
func ReadonlyLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldReadonly, v))
}

// ReadonlyContains applies the Contains predicate on the "readonly" field.
func ReadonlyContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldReadonly, v))
}

// ReadonlyHasPrefix applies the HasPrefix predicate on the "readonly" field.
func ReadonlyHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldReadonly, v))
}

// ReadonlyHasSuffix applies the HasSuffix predicate on the "readonly" field.
func ReadonlyHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldReadonly, v))
}

// ReadonlyIsNil applies the IsNil predicate on the "readonly" field.
func ReadonlyIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldReadonly))
}

// ReadonlyNotNil applies the NotNil predicate on the "readonly" field.
func ReadonlyNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldReadonly))
}

// ReadonlyEqualFold applies the EqualFold predicate on the "readonly" field.
func ReadonlyEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldReadonly, v))
}

// ReadonlyContainsFold applies the ContainsFold predicate on the "readonly" field.
func ReadonlyContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldReadonly, v))
}

// HasPets applies the HasEdge predicate on the "pets" edge.
func HasPets() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PetsTable, PetsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPetsWith applies the HasEdge predicate on the "pets" edge with a given conditions (other predicates).
func HasPetsWith(preds ...predicate.Pet) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := newPetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(sql.NotPredicates(p))
}
