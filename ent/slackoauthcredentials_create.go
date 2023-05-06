// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/slackoauthcredentials"
)

// SlackOAuthCredentialsCreate is the builder for creating a SlackOAuthCredentials entity.
type SlackOAuthCredentialsCreate struct {
	config
	mutation *SlackOAuthCredentialsMutation
	hooks    []Hook
}

// SetTeamID sets the "team_id" field.
func (socc *SlackOAuthCredentialsCreate) SetTeamID(s string) *SlackOAuthCredentialsCreate {
	socc.mutation.SetTeamID(s)
	return socc
}

// SetAccessToken sets the "access_token" field.
func (socc *SlackOAuthCredentialsCreate) SetAccessToken(s string) *SlackOAuthCredentialsCreate {
	socc.mutation.SetAccessToken(s)
	return socc
}

// SetCreatedAt sets the "created_at" field.
func (socc *SlackOAuthCredentialsCreate) SetCreatedAt(i int64) *SlackOAuthCredentialsCreate {
	socc.mutation.SetCreatedAt(i)
	return socc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (socc *SlackOAuthCredentialsCreate) SetNillableCreatedAt(i *int64) *SlackOAuthCredentialsCreate {
	if i != nil {
		socc.SetCreatedAt(*i)
	}
	return socc
}

// SetUpdatedAt sets the "updated_at" field.
func (socc *SlackOAuthCredentialsCreate) SetUpdatedAt(i int64) *SlackOAuthCredentialsCreate {
	socc.mutation.SetUpdatedAt(i)
	return socc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (socc *SlackOAuthCredentialsCreate) SetNillableUpdatedAt(i *int64) *SlackOAuthCredentialsCreate {
	if i != nil {
		socc.SetUpdatedAt(*i)
	}
	return socc
}

// SetID sets the "id" field.
func (socc *SlackOAuthCredentialsCreate) SetID(u uuid.UUID) *SlackOAuthCredentialsCreate {
	socc.mutation.SetID(u)
	return socc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (socc *SlackOAuthCredentialsCreate) SetNillableID(u *uuid.UUID) *SlackOAuthCredentialsCreate {
	if u != nil {
		socc.SetID(*u)
	}
	return socc
}

// Mutation returns the SlackOAuthCredentialsMutation object of the builder.
func (socc *SlackOAuthCredentialsCreate) Mutation() *SlackOAuthCredentialsMutation {
	return socc.mutation
}

// Save creates the SlackOAuthCredentials in the database.
func (socc *SlackOAuthCredentialsCreate) Save(ctx context.Context) (*SlackOAuthCredentials, error) {
	socc.defaults()
	return withHooks[*SlackOAuthCredentials, SlackOAuthCredentialsMutation](ctx, socc.sqlSave, socc.mutation, socc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (socc *SlackOAuthCredentialsCreate) SaveX(ctx context.Context) *SlackOAuthCredentials {
	v, err := socc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (socc *SlackOAuthCredentialsCreate) Exec(ctx context.Context) error {
	_, err := socc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (socc *SlackOAuthCredentialsCreate) ExecX(ctx context.Context) {
	if err := socc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (socc *SlackOAuthCredentialsCreate) defaults() {
	if _, ok := socc.mutation.CreatedAt(); !ok {
		v := slackoauthcredentials.DefaultCreatedAt()
		socc.mutation.SetCreatedAt(v)
	}
	if _, ok := socc.mutation.UpdatedAt(); !ok {
		v := slackoauthcredentials.DefaultUpdatedAt()
		socc.mutation.SetUpdatedAt(v)
	}
	if _, ok := socc.mutation.ID(); !ok {
		v := slackoauthcredentials.DefaultID()
		socc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (socc *SlackOAuthCredentialsCreate) check() error {
	if _, ok := socc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`ent: missing required field "SlackOAuthCredentials.team_id"`)}
	}
	if v, ok := socc.mutation.TeamID(); ok {
		if err := slackoauthcredentials.TeamIDValidator(v); err != nil {
			return &ValidationError{Name: "team_id", err: fmt.Errorf(`ent: validator failed for field "SlackOAuthCredentials.team_id": %w`, err)}
		}
	}
	if _, ok := socc.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New(`ent: missing required field "SlackOAuthCredentials.access_token"`)}
	}
	if v, ok := socc.mutation.AccessToken(); ok {
		if err := slackoauthcredentials.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "SlackOAuthCredentials.access_token": %w`, err)}
		}
	}
	if _, ok := socc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SlackOAuthCredentials.created_at"`)}
	}
	if _, ok := socc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SlackOAuthCredentials.updated_at"`)}
	}
	return nil
}

func (socc *SlackOAuthCredentialsCreate) sqlSave(ctx context.Context) (*SlackOAuthCredentials, error) {
	if err := socc.check(); err != nil {
		return nil, err
	}
	_node, _spec := socc.createSpec()
	if err := sqlgraph.CreateNode(ctx, socc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	socc.mutation.id = &_node.ID
	socc.mutation.done = true
	return _node, nil
}

func (socc *SlackOAuthCredentialsCreate) createSpec() (*SlackOAuthCredentials, *sqlgraph.CreateSpec) {
	var (
		_node = &SlackOAuthCredentials{config: socc.config}
		_spec = sqlgraph.NewCreateSpec(slackoauthcredentials.Table, sqlgraph.NewFieldSpec(slackoauthcredentials.FieldID, field.TypeUUID))
	)
	_spec.Schema = socc.schemaConfig.SlackOAuthCredentials
	if id, ok := socc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := socc.mutation.TeamID(); ok {
		_spec.SetField(slackoauthcredentials.FieldTeamID, field.TypeString, value)
		_node.TeamID = value
	}
	if value, ok := socc.mutation.AccessToken(); ok {
		_spec.SetField(slackoauthcredentials.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := socc.mutation.CreatedAt(); ok {
		_spec.SetField(slackoauthcredentials.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := socc.mutation.UpdatedAt(); ok {
		_spec.SetField(slackoauthcredentials.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// SlackOAuthCredentialsCreateBulk is the builder for creating many SlackOAuthCredentials entities in bulk.
type SlackOAuthCredentialsCreateBulk struct {
	config
	builders []*SlackOAuthCredentialsCreate
}

// Save creates the SlackOAuthCredentials entities in the database.
func (soccb *SlackOAuthCredentialsCreateBulk) Save(ctx context.Context) ([]*SlackOAuthCredentials, error) {
	specs := make([]*sqlgraph.CreateSpec, len(soccb.builders))
	nodes := make([]*SlackOAuthCredentials, len(soccb.builders))
	mutators := make([]Mutator, len(soccb.builders))
	for i := range soccb.builders {
		func(i int, root context.Context) {
			builder := soccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SlackOAuthCredentialsMutation)
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
					_, err = mutators[i+1].Mutate(root, soccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, soccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, soccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (soccb *SlackOAuthCredentialsCreateBulk) SaveX(ctx context.Context) []*SlackOAuthCredentials {
	v, err := soccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (soccb *SlackOAuthCredentialsCreateBulk) Exec(ctx context.Context) error {
	_, err := soccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (soccb *SlackOAuthCredentialsCreateBulk) ExecX(ctx context.Context) {
	if err := soccb.Exec(ctx); err != nil {
		panic(err)
	}
}