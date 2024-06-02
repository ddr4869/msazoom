// Code generated by ent, DO NOT EDIT.

package chat

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the chat type in the database.
	Label = "chat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChatName holds the string denoting the chat_name field in the database.
	FieldChatName = "chat_name"
	// FieldChatUser holds the string denoting the chat_user field in the database.
	FieldChatUser = "chat_user"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the chat in the database.
	Table = "chats"
)

// Columns holds all SQL columns for chat fields.
var Columns = []string{
	FieldID,
	FieldChatName,
	FieldChatUser,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultChatName holds the default value on creation for the "chat_name" field.
	DefaultChatName string
	// DefaultChatUser holds the default value on creation for the "chat_user" field.
	DefaultChatUser string
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Chat queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChatName orders the results by the chat_name field.
func ByChatName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatName, opts...).ToFunc()
}

// ByChatUser orders the results by the chat_user field.
func ByChatUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatUser, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
