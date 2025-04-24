package schema

// Schema - a schema describes the structure of a project.
type Schema struct {
	// Version - the version of the schema.
	Version string `json:"version" yaml:"version"`

	// Name - the name of the schema.
	Name string `json:"name" yaml:"name"`

	// Description - a description of the schema.
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}
