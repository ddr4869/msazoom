// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BoardsColumns holds the columns for the "boards" table.
	BoardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "board_name", Type: field.TypeString, Default: "unknown"},
		{Name: "board_admin", Type: field.TypeString, Default: "unknown"},
		{Name: "board_password", Type: field.TypeString, Default: "unknown"},
		{Name: "board_star", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// BoardsTable holds the schema information for the "boards" table.
	BoardsTable = &schema.Table{
		Name:       "boards",
		Columns:    BoardsColumns,
		PrimaryKey: []*schema.Column{BoardsColumns[0]},
	}
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "chat_name", Type: field.TypeString, Default: "unknown"},
		{Name: "chat_user", Type: field.TypeString, Default: "unknown"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
	}
	// FriendsColumns holds the columns for the "friends" table.
	FriendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "friend", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// FriendsTable holds the schema information for the "friends" table.
	FriendsTable = &schema.Table{
		Name:       "friends",
		Columns:    FriendsColumns,
		PrimaryKey: []*schema.Column{FriendsColumns[0]},
	}
	// HotBoardsColumns holds the columns for the "hot_boards" table.
	HotBoardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// HotBoardsTable holds the schema information for the "hot_boards" table.
	HotBoardsTable = &schema.Table{
		Name:       "hot_boards",
		Columns:    HotBoardsColumns,
		PrimaryKey: []*schema.Column{HotBoardsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message", Type: field.TypeString},
		{Name: "writer", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "board_messages", Type: field.TypeInt, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_boards_messages",
				Columns:    []*schema.Column{MessagesColumns[5]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BoardsTable,
		ChatsTable,
		FriendsTable,
		HotBoardsTable,
		MessagesTable,
		UsersTable,
	}
)

func init() {
	MessagesTable.ForeignKeys[0].RefTable = BoardsTable
}
