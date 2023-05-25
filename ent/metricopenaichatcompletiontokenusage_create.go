// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/metricopenaichatcompletiontokenusage"
)

// MetricOpenAIChatCompletionTokenUsageCreate is the builder for creating a MetricOpenAIChatCompletionTokenUsage entity.
type MetricOpenAIChatCompletionTokenUsageCreate struct {
	config
	mutation *MetricOpenAIChatCompletionTokenUsageMutation
	hooks    []Hook
}

// SetPromptOperation sets the "prompt_operation" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetPromptOperation(s string) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetPromptOperation(s)
	return moacctuc
}

// SetNillablePromptOperation sets the "prompt_operation" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillablePromptOperation(s *string) *MetricOpenAIChatCompletionTokenUsageCreate {
	if s != nil {
		moacctuc.SetPromptOperation(*s)
	}
	return moacctuc
}

// SetPromptCharacterLength sets the "prompt_character_length" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetPromptCharacterLength(i int) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetPromptCharacterLength(i)
	return moacctuc
}

// SetNillablePromptCharacterLength sets the "prompt_character_length" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillablePromptCharacterLength(i *int) *MetricOpenAIChatCompletionTokenUsageCreate {
	if i != nil {
		moacctuc.SetPromptCharacterLength(*i)
	}
	return moacctuc
}

// SetPromptTokenUsage sets the "prompt_token_usage" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetPromptTokenUsage(i int) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetPromptTokenUsage(i)
	return moacctuc
}

// SetNillablePromptTokenUsage sets the "prompt_token_usage" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillablePromptTokenUsage(i *int) *MetricOpenAIChatCompletionTokenUsageCreate {
	if i != nil {
		moacctuc.SetPromptTokenUsage(*i)
	}
	return moacctuc
}

// SetCompletionCharacterLength sets the "completion_character_length" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetCompletionCharacterLength(i int) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetCompletionCharacterLength(i)
	return moacctuc
}

// SetNillableCompletionCharacterLength sets the "completion_character_length" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillableCompletionCharacterLength(i *int) *MetricOpenAIChatCompletionTokenUsageCreate {
	if i != nil {
		moacctuc.SetCompletionCharacterLength(*i)
	}
	return moacctuc
}

// SetCompletionTokenUsage sets the "completion_token_usage" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetCompletionTokenUsage(i int) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetCompletionTokenUsage(i)
	return moacctuc
}

// SetNillableCompletionTokenUsage sets the "completion_token_usage" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillableCompletionTokenUsage(i *int) *MetricOpenAIChatCompletionTokenUsageCreate {
	if i != nil {
		moacctuc.SetCompletionTokenUsage(*i)
	}
	return moacctuc
}

// SetTotalTokenUsage sets the "total_token_usage" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetTotalTokenUsage(i int) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetTotalTokenUsage(i)
	return moacctuc
}

// SetNillableTotalTokenUsage sets the "total_token_usage" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillableTotalTokenUsage(i *int) *MetricOpenAIChatCompletionTokenUsageCreate {
	if i != nil {
		moacctuc.SetTotalTokenUsage(*i)
	}
	return moacctuc
}

// SetCreatedAt sets the "created_at" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetCreatedAt(i int64) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetCreatedAt(i)
	return moacctuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillableCreatedAt(i *int64) *MetricOpenAIChatCompletionTokenUsageCreate {
	if i != nil {
		moacctuc.SetCreatedAt(*i)
	}
	return moacctuc
}

// SetID sets the "id" field.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetID(u uuid.UUID) *MetricOpenAIChatCompletionTokenUsageCreate {
	moacctuc.mutation.SetID(u)
	return moacctuc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SetNillableID(u *uuid.UUID) *MetricOpenAIChatCompletionTokenUsageCreate {
	if u != nil {
		moacctuc.SetID(*u)
	}
	return moacctuc
}

// Mutation returns the MetricOpenAIChatCompletionTokenUsageMutation object of the builder.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) Mutation() *MetricOpenAIChatCompletionTokenUsageMutation {
	return moacctuc.mutation
}

// Save creates the MetricOpenAIChatCompletionTokenUsage in the database.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) Save(ctx context.Context) (*MetricOpenAIChatCompletionTokenUsage, error) {
	moacctuc.defaults()
	return withHooks(ctx, moacctuc.sqlSave, moacctuc.mutation, moacctuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) SaveX(ctx context.Context) *MetricOpenAIChatCompletionTokenUsage {
	v, err := moacctuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) Exec(ctx context.Context) error {
	_, err := moacctuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) ExecX(ctx context.Context) {
	if err := moacctuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) defaults() {
	if _, ok := moacctuc.mutation.PromptOperation(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultPromptOperation
		moacctuc.mutation.SetPromptOperation(v)
	}
	if _, ok := moacctuc.mutation.PromptCharacterLength(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultPromptCharacterLength
		moacctuc.mutation.SetPromptCharacterLength(v)
	}
	if _, ok := moacctuc.mutation.PromptTokenUsage(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultPromptTokenUsage
		moacctuc.mutation.SetPromptTokenUsage(v)
	}
	if _, ok := moacctuc.mutation.CompletionCharacterLength(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultCompletionCharacterLength
		moacctuc.mutation.SetCompletionCharacterLength(v)
	}
	if _, ok := moacctuc.mutation.CompletionTokenUsage(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultCompletionTokenUsage
		moacctuc.mutation.SetCompletionTokenUsage(v)
	}
	if _, ok := moacctuc.mutation.TotalTokenUsage(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultTotalTokenUsage
		moacctuc.mutation.SetTotalTokenUsage(v)
	}
	if _, ok := moacctuc.mutation.CreatedAt(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultCreatedAt()
		moacctuc.mutation.SetCreatedAt(v)
	}
	if _, ok := moacctuc.mutation.ID(); !ok {
		v := metricopenaichatcompletiontokenusage.DefaultID()
		moacctuc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) check() error {
	if _, ok := moacctuc.mutation.PromptOperation(); !ok {
		return &ValidationError{Name: "prompt_operation", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.prompt_operation"`)}
	}
	if _, ok := moacctuc.mutation.PromptCharacterLength(); !ok {
		return &ValidationError{Name: "prompt_character_length", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.prompt_character_length"`)}
	}
	if _, ok := moacctuc.mutation.PromptTokenUsage(); !ok {
		return &ValidationError{Name: "prompt_token_usage", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.prompt_token_usage"`)}
	}
	if _, ok := moacctuc.mutation.CompletionCharacterLength(); !ok {
		return &ValidationError{Name: "completion_character_length", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.completion_character_length"`)}
	}
	if _, ok := moacctuc.mutation.CompletionTokenUsage(); !ok {
		return &ValidationError{Name: "completion_token_usage", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.completion_token_usage"`)}
	}
	if _, ok := moacctuc.mutation.TotalTokenUsage(); !ok {
		return &ValidationError{Name: "total_token_usage", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.total_token_usage"`)}
	}
	if _, ok := moacctuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MetricOpenAIChatCompletionTokenUsage.created_at"`)}
	}
	return nil
}

func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) sqlSave(ctx context.Context) (*MetricOpenAIChatCompletionTokenUsage, error) {
	if err := moacctuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := moacctuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, moacctuc.driver, _spec); err != nil {
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
	moacctuc.mutation.id = &_node.ID
	moacctuc.mutation.done = true
	return _node, nil
}

func (moacctuc *MetricOpenAIChatCompletionTokenUsageCreate) createSpec() (*MetricOpenAIChatCompletionTokenUsage, *sqlgraph.CreateSpec) {
	var (
		_node = &MetricOpenAIChatCompletionTokenUsage{config: moacctuc.config}
		_spec = sqlgraph.NewCreateSpec(metricopenaichatcompletiontokenusage.Table, sqlgraph.NewFieldSpec(metricopenaichatcompletiontokenusage.FieldID, field.TypeUUID))
	)
	_spec.Schema = moacctuc.schemaConfig.MetricOpenAIChatCompletionTokenUsage
	if id, ok := moacctuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := moacctuc.mutation.PromptOperation(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldPromptOperation, field.TypeString, value)
		_node.PromptOperation = value
	}
	if value, ok := moacctuc.mutation.PromptCharacterLength(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldPromptCharacterLength, field.TypeInt, value)
		_node.PromptCharacterLength = value
	}
	if value, ok := moacctuc.mutation.PromptTokenUsage(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldPromptTokenUsage, field.TypeInt, value)
		_node.PromptTokenUsage = value
	}
	if value, ok := moacctuc.mutation.CompletionCharacterLength(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldCompletionCharacterLength, field.TypeInt, value)
		_node.CompletionCharacterLength = value
	}
	if value, ok := moacctuc.mutation.CompletionTokenUsage(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldCompletionTokenUsage, field.TypeInt, value)
		_node.CompletionTokenUsage = value
	}
	if value, ok := moacctuc.mutation.TotalTokenUsage(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldTotalTokenUsage, field.TypeInt, value)
		_node.TotalTokenUsage = value
	}
	if value, ok := moacctuc.mutation.CreatedAt(); ok {
		_spec.SetField(metricopenaichatcompletiontokenusage.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// MetricOpenAIChatCompletionTokenUsageCreateBulk is the builder for creating many MetricOpenAIChatCompletionTokenUsage entities in bulk.
type MetricOpenAIChatCompletionTokenUsageCreateBulk struct {
	config
	builders []*MetricOpenAIChatCompletionTokenUsageCreate
}

// Save creates the MetricOpenAIChatCompletionTokenUsage entities in the database.
func (moacctucb *MetricOpenAIChatCompletionTokenUsageCreateBulk) Save(ctx context.Context) ([]*MetricOpenAIChatCompletionTokenUsage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(moacctucb.builders))
	nodes := make([]*MetricOpenAIChatCompletionTokenUsage, len(moacctucb.builders))
	mutators := make([]Mutator, len(moacctucb.builders))
	for i := range moacctucb.builders {
		func(i int, root context.Context) {
			builder := moacctucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricOpenAIChatCompletionTokenUsageMutation)
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
					_, err = mutators[i+1].Mutate(root, moacctucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, moacctucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, moacctucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (moacctucb *MetricOpenAIChatCompletionTokenUsageCreateBulk) SaveX(ctx context.Context) []*MetricOpenAIChatCompletionTokenUsage {
	v, err := moacctucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (moacctucb *MetricOpenAIChatCompletionTokenUsageCreateBulk) Exec(ctx context.Context) error {
	_, err := moacctucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (moacctucb *MetricOpenAIChatCompletionTokenUsageCreateBulk) ExecX(ctx context.Context) {
	if err := moacctucb.Exec(ctx); err != nil {
		panic(err)
	}
}
