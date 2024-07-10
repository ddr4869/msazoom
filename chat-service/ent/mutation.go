// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ddr4869/msazoom/chat-service/ent/chat"
	"github.com/ddr4869/msazoom/chat-service/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeChat = "Chat"
)

// ChatMutation represents an operation that mutates the Chat nodes in the graph.
type ChatMutation struct {
	config
	op            Op
	typ           string
	id            *int
	chat_name     *string
	chat_user     *string
	chat_password *string
	createdAt     *time.Time
	updatedAt     *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Chat, error)
	predicates    []predicate.Chat
}

var _ ent.Mutation = (*ChatMutation)(nil)

// chatOption allows management of the mutation configuration using functional options.
type chatOption func(*ChatMutation)

// newChatMutation creates new mutation for the Chat entity.
func newChatMutation(c config, op Op, opts ...chatOption) *ChatMutation {
	m := &ChatMutation{
		config:        c,
		op:            op,
		typ:           TypeChat,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatID sets the ID field of the mutation.
func withChatID(id int) chatOption {
	return func(m *ChatMutation) {
		var (
			err   error
			once  sync.Once
			value *Chat
		)
		m.oldValue = func(ctx context.Context) (*Chat, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Chat.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChat sets the old Chat of the mutation.
func withChat(node *Chat) chatOption {
	return func(m *ChatMutation) {
		m.oldValue = func(context.Context) (*Chat, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Chat.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetChatName sets the "chat_name" field.
func (m *ChatMutation) SetChatName(s string) {
	m.chat_name = &s
}

// ChatName returns the value of the "chat_name" field in the mutation.
func (m *ChatMutation) ChatName() (r string, exists bool) {
	v := m.chat_name
	if v == nil {
		return
	}
	return *v, true
}

// OldChatName returns the old "chat_name" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldChatName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldChatName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldChatName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldChatName: %w", err)
	}
	return oldValue.ChatName, nil
}

// ResetChatName resets all changes to the "chat_name" field.
func (m *ChatMutation) ResetChatName() {
	m.chat_name = nil
}

// SetChatUser sets the "chat_user" field.
func (m *ChatMutation) SetChatUser(s string) {
	m.chat_user = &s
}

// ChatUser returns the value of the "chat_user" field in the mutation.
func (m *ChatMutation) ChatUser() (r string, exists bool) {
	v := m.chat_user
	if v == nil {
		return
	}
	return *v, true
}

// OldChatUser returns the old "chat_user" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldChatUser(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldChatUser is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldChatUser requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldChatUser: %w", err)
	}
	return oldValue.ChatUser, nil
}

// ResetChatUser resets all changes to the "chat_user" field.
func (m *ChatMutation) ResetChatUser() {
	m.chat_user = nil
}

// SetChatPassword sets the "chat_password" field.
func (m *ChatMutation) SetChatPassword(s string) {
	m.chat_password = &s
}

// ChatPassword returns the value of the "chat_password" field in the mutation.
func (m *ChatMutation) ChatPassword() (r string, exists bool) {
	v := m.chat_password
	if v == nil {
		return
	}
	return *v, true
}

// OldChatPassword returns the old "chat_password" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldChatPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldChatPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldChatPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldChatPassword: %w", err)
	}
	return oldValue.ChatPassword, nil
}

// ClearChatPassword clears the value of the "chat_password" field.
func (m *ChatMutation) ClearChatPassword() {
	m.chat_password = nil
	m.clearedFields[chat.FieldChatPassword] = struct{}{}
}

// ChatPasswordCleared returns if the "chat_password" field was cleared in this mutation.
func (m *ChatMutation) ChatPasswordCleared() bool {
	_, ok := m.clearedFields[chat.FieldChatPassword]
	return ok
}

// ResetChatPassword resets all changes to the "chat_password" field.
func (m *ChatMutation) ResetChatPassword() {
	m.chat_password = nil
	delete(m.clearedFields, chat.FieldChatPassword)
}

// SetCreatedAt sets the "createdAt" field.
func (m *ChatMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *ChatMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *ChatMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// SetUpdatedAt sets the "updatedAt" field.
func (m *ChatMutation) SetUpdatedAt(t time.Time) {
	m.updatedAt = &t
}

// UpdatedAt returns the value of the "updatedAt" field in the mutation.
func (m *ChatMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updatedAt
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updatedAt" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updatedAt" field.
func (m *ChatMutation) ResetUpdatedAt() {
	m.updatedAt = nil
}

// Where appends a list predicates to the ChatMutation builder.
func (m *ChatMutation) Where(ps ...predicate.Chat) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Chat, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Chat).
func (m *ChatMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.chat_name != nil {
		fields = append(fields, chat.FieldChatName)
	}
	if m.chat_user != nil {
		fields = append(fields, chat.FieldChatUser)
	}
	if m.chat_password != nil {
		fields = append(fields, chat.FieldChatPassword)
	}
	if m.createdAt != nil {
		fields = append(fields, chat.FieldCreatedAt)
	}
	if m.updatedAt != nil {
		fields = append(fields, chat.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chat.FieldChatName:
		return m.ChatName()
	case chat.FieldChatUser:
		return m.ChatUser()
	case chat.FieldChatPassword:
		return m.ChatPassword()
	case chat.FieldCreatedAt:
		return m.CreatedAt()
	case chat.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chat.FieldChatName:
		return m.OldChatName(ctx)
	case chat.FieldChatUser:
		return m.OldChatUser(ctx)
	case chat.FieldChatPassword:
		return m.OldChatPassword(ctx)
	case chat.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case chat.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Chat field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chat.FieldChatName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetChatName(v)
		return nil
	case chat.FieldChatUser:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetChatUser(v)
		return nil
	case chat.FieldChatPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetChatPassword(v)
		return nil
	case chat.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case chat.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Chat numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(chat.FieldChatPassword) {
		fields = append(fields, chat.FieldChatPassword)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatMutation) ClearField(name string) error {
	switch name {
	case chat.FieldChatPassword:
		m.ClearChatPassword()
		return nil
	}
	return fmt.Errorf("unknown Chat nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatMutation) ResetField(name string) error {
	switch name {
	case chat.FieldChatName:
		m.ResetChatName()
		return nil
	case chat.FieldChatUser:
		m.ResetChatUser()
		return nil
	case chat.FieldChatPassword:
		m.ResetChatPassword()
		return nil
	case chat.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case chat.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Chat unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Chat edge %s", name)
}
