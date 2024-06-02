// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/msazoom/ent/chat"
)

// ChatCreate is the builder for creating a Chat entity.
type ChatCreate struct {
	config
	mutation *ChatMutation
	hooks    []Hook
}

// SetChatName sets the "chat_name" field.
func (cc *ChatCreate) SetChatName(s string) *ChatCreate {
	cc.mutation.SetChatName(s)
	return cc
}

// SetNillableChatName sets the "chat_name" field if the given value is not nil.
func (cc *ChatCreate) SetNillableChatName(s *string) *ChatCreate {
	if s != nil {
		cc.SetChatName(*s)
	}
	return cc
}

// SetChatUser sets the "chat_user" field.
func (cc *ChatCreate) SetChatUser(s string) *ChatCreate {
	cc.mutation.SetChatUser(s)
	return cc
}

// SetNillableChatUser sets the "chat_user" field if the given value is not nil.
func (cc *ChatCreate) SetNillableChatUser(s *string) *ChatCreate {
	if s != nil {
		cc.SetChatUser(*s)
	}
	return cc
}

// SetCreatedAt sets the "createdAt" field.
func (cc *ChatCreate) SetCreatedAt(t time.Time) *ChatCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (cc *ChatCreate) SetNillableCreatedAt(t *time.Time) *ChatCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updatedAt" field.
func (cc *ChatCreate) SetUpdatedAt(t time.Time) *ChatCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (cc *ChatCreate) SetNillableUpdatedAt(t *time.Time) *ChatCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// Mutation returns the ChatMutation object of the builder.
func (cc *ChatCreate) Mutation() *ChatMutation {
	return cc.mutation
}

// Save creates the Chat in the database.
func (cc *ChatCreate) Save(ctx context.Context) (*Chat, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChatCreate) SaveX(ctx context.Context) *Chat {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChatCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChatCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChatCreate) defaults() {
	if _, ok := cc.mutation.ChatName(); !ok {
		v := chat.DefaultChatName
		cc.mutation.SetChatName(v)
	}
	if _, ok := cc.mutation.ChatUser(); !ok {
		v := chat.DefaultChatUser
		cc.mutation.SetChatUser(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := chat.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := chat.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChatCreate) check() error {
	if _, ok := cc.mutation.ChatName(); !ok {
		return &ValidationError{Name: "chat_name", err: errors.New(`ent: missing required field "Chat.chat_name"`)}
	}
	if _, ok := cc.mutation.ChatUser(); !ok {
		return &ValidationError{Name: "chat_user", err: errors.New(`ent: missing required field "Chat.chat_user"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Chat.createdAt"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Chat.updatedAt"`)}
	}
	return nil
}

func (cc *ChatCreate) sqlSave(ctx context.Context) (*Chat, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChatCreate) createSpec() (*Chat, *sqlgraph.CreateSpec) {
	var (
		_node = &Chat{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chat.Table, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.ChatName(); ok {
		_spec.SetField(chat.FieldChatName, field.TypeString, value)
		_node.ChatName = value
	}
	if value, ok := cc.mutation.ChatUser(); ok {
		_spec.SetField(chat.FieldChatUser, field.TypeString, value)
		_node.ChatUser = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(chat.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(chat.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// ChatCreateBulk is the builder for creating many Chat entities in bulk.
type ChatCreateBulk struct {
	config
	err      error
	builders []*ChatCreate
}

// Save creates the Chat entities in the database.
func (ccb *ChatCreateBulk) Save(ctx context.Context) ([]*Chat, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chat, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChatCreateBulk) SaveX(ctx context.Context) []*Chat {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChatCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChatCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
