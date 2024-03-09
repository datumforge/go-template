// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/datumforge/go-template/internal/ent/schema\",\"Package\":\"github.com/datumforge/go-template/internal/ent/generated\",\"Schemas\":[{\"name\":\"Todo\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"the name of the organization\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"An optional description of the organization\"}],\"indexes\":[{\"unique\":true,\"fields\":[\"name\"]}],\"annotations\":{\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}]}}}],\"Features\":[\"sql/versioned-migration\",\"privacy\",\"schema/snapshot\",\"entql\",\"namedges\",\"sql/schemaconfig\",\"intercept\",\"namedges\"]}"