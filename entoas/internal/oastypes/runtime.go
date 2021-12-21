// Code generated by entc, DO NOT EDIT.

package oastypes

import (
	"entgo.io/contrib/entoas/internal/oastypes/oastypes"
	"entgo.io/contrib/entoas/internal/oastypes/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	oastypesFields := schema.OASTypes{}.Fields()
	_ = oastypesFields
	// oastypesDescUUID is the schema descriptor for uuid field.
	oastypesDescUUID := oastypesFields[14].Descriptor()
	// oastypes.DefaultUUID holds the default value on creation for the uuid field.
	oastypes.DefaultUUID = oastypesDescUUID.Default.(func() uuid.UUID)
	// oastypesDescOther is the schema descriptor for other field.
	oastypesDescOther := oastypesFields[25].Descriptor()
	// oastypes.DefaultOther holds the default value on creation for the other field.
	oastypes.DefaultOther = oastypesDescOther.Default.(*schema.Link)
}