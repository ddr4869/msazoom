// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ddr4869/msazoom/user-service/ent/schema"
	"github.com/ddr4869/msazoom/user-service/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[2].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(int)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}
