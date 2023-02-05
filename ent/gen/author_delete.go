// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmhafiz/go8/ent/gen/author"
	"github.com/gmhafiz/go8/ent/gen/predicate"
)

// AuthorDelete is the builder for deleting a Author entity.
type AuthorDelete struct {
	config
	hooks    []Hook
	mutation *AuthorMutation
}

// Where appends a list predicates to the AuthorDelete builder.
func (ad *AuthorDelete) Where(ps ...predicate.Author) *AuthorDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AuthorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AuthorMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AuthorDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AuthorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: author.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: author.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AuthorDeleteOne is the builder for deleting a single Author entity.
type AuthorDeleteOne struct {
	ad *AuthorDelete
}

// Where appends a list predicates to the AuthorDelete builder.
func (ado *AuthorDeleteOne) Where(ps ...predicate.Author) *AuthorDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AuthorDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{author.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AuthorDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
