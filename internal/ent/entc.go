//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
	"github.com/ogen-go/ogen"
)

func main() {
	// Ensure the schema directory exists before running entc.
	_ = os.Mkdir("schema", 0755)

	// Add OpenAPI Gen extension
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithRelaySpec(true),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{
		Target:  "./internal/ent/generated",
		Package: "github.com/datumforge/go-template/internal/ent/generated",
	},
		entc.Extensions(
			oas,
			entviz.Extension{},
			gqlExt,
		)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
