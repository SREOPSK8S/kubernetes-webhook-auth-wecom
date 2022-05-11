// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/ent/audit"
)

// AuditCreate is the builder for creating a Audit entity.
type AuditCreate struct {
	config
	mutation *AuditMutation
	hooks    []Hook
}

// SetUID sets the "u_id" field.
func (ac *AuditCreate) SetUID(s string) *AuditCreate {
	ac.mutation.SetUID(s)
	return ac
}

// SetCertificationTime sets the "certification_time" field.
func (ac *AuditCreate) SetCertificationTime(t time.Time) *AuditCreate {
	ac.mutation.SetCertificationTime(t)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AuditCreate) SetCreatedAt(t time.Time) *AuditCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AuditCreate) SetNillableCreatedAt(t *time.Time) *AuditCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AuditCreate) SetUpdatedAt(t time.Time) *AuditCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AuditCreate) SetNillableUpdatedAt(t *time.Time) *AuditCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// Mutation returns the AuditMutation object of the builder.
func (ac *AuditCreate) Mutation() *AuditMutation {
	return ac.mutation
}

// Save creates the Audit in the database.
func (ac *AuditCreate) Save(ctx context.Context) (*Audit, error) {
	var (
		err  error
		node *Audit
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuditMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuditCreate) SaveX(ctx context.Context) *Audit {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuditCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuditCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuditCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := audit.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := audit.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuditCreate) check() error {
	if _, ok := ac.mutation.UID(); !ok {
		return &ValidationError{Name: "u_id", err: errors.New(`ent: missing required field "Audit.u_id"`)}
	}
	if _, ok := ac.mutation.CertificationTime(); !ok {
		return &ValidationError{Name: "certification_time", err: errors.New(`ent: missing required field "Audit.certification_time"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Audit.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Audit.updated_at"`)}
	}
	return nil
}

func (ac *AuditCreate) sqlSave(ctx context.Context) (*Audit, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AuditCreate) createSpec() (*Audit, *sqlgraph.CreateSpec) {
	var (
		_node = &Audit{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: audit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: audit.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: audit.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := ac.mutation.CertificationTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: audit.FieldCertificationTime,
		})
		_node.CertificationTime = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: audit.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: audit.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AuditCreateBulk is the builder for creating many Audit entities in bulk.
type AuditCreateBulk struct {
	config
	builders []*AuditCreate
}

// Save creates the Audit entities in the database.
func (acb *AuditCreateBulk) Save(ctx context.Context) ([]*Audit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Audit, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuditMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuditCreateBulk) SaveX(ctx context.Context) []*Audit {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuditCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuditCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
